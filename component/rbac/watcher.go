package rbac

import (
	"github.com/kenSevLeb/go-framework/util/strings"
	"sync"

	"github.com/casbin/casbin/v2/persist"
	"github.com/go-redis/redis"
)

var _ persist.Watcher = &Watcher{}

// New creates a Watcher
func NewWatcher(redisConn redis.UniversalClient) (*Watcher, error) {
	w := &Watcher{
		publisher:  redisConn,
		subscriber: redisConn,
		callback:   nil,
		once:       sync.Once{},
		closeChan:  nil,
		channel:    "/casbin",
		localID:    strings.UUID(),
	}

	if err := w.init(); err != nil {
		return nil, err
	}

	return w, nil
}

// Watcher implements Casbin's persist.Watcher using Redis as a backend
type Watcher struct {
	publisher  redis.UniversalClient
	subscriber redis.UniversalClient
	callback   func(string)
	once       sync.Once
	closeChan  chan struct{}
	channel    string
	localID    string
}

// SetUpdateCallback sets the callback function that the watcher will call
// when the policy in DB has been changed by other instances.
// A classic callback is Enforcer.LoadPolicy().
func (w *Watcher) SetUpdateCallback(callback func(string)) error {
	w.callback = callback
	return nil
}

// Update calls the update callback of other instances to synchronize their policy.
// It is usually called after changing the policy in DB, like Enforcer.SavePolicy(),
// Enforcer.AddPolicy(), Enforcer.RemovePolicy(), etc.
func (w *Watcher) Update() error {
	if err := w.publisher.Publish(w.channel, w.localID).Err(); err != nil {
		return err
	}

	return nil
}

// Close stops and releases the watcher, the callback function will not be called any more.
func (w *Watcher) Close() {
	w.once.Do(func() {
		close(w.closeChan)
	})
}

func (w *Watcher) init() error {

	pubsub := w.subscriber.Subscribe(w.channel)

	if _, err := pubsub.Receive(); err != nil {
		return err
	}

	w.closeChan = make(chan struct{})
	go func() {
		for {
			select {
			case msg := <-pubsub.Channel():
				w.messageReceived(msg.Payload)

			case <-w.closeChan:
				pubsub.Close()
				w.subscriber.Close()
				w.publisher.Close()
				return
			}
		}
	}()

	return nil
}

func (w *Watcher) messageReceived(publisherID string) {
	if publisherID == w.localID {
		// ignore messages from itself
		return
	}

	if w.callback != nil {
		w.callback(publisherID)
	}
}

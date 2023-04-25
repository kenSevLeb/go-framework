package event

import "fmt"

// default dispatcher
var defaultInstance = New()

type Dispatcher interface {
	// 注册事件
	Register(name string, listener Listener)
	// 判断是否存在事件
	Has(name string) bool
	// 移除事件
	Remove(name string)
	// 触发事件
	Trigger(name string, params interface{})
}

// eventDispatcher implement of Dispatcher
type dispatcher struct {
	events map[string][]Listener
}

// New 实例化
func New() *dispatcher {
	return &dispatcher{events: map[string][]Listener{}}
}

// Register quickly register event with handler
func Register(name string, listener Listener) {
	defaultInstance.Register(name, listener)
}

// Register
func (instance *dispatcher) Register(name string, listener Listener) {
	if _, ok := instance.events[name]; !ok {
		instance.events[name] = make([]Listener, 0)
	}
	instance.events[name] = append(instance.events[name], listener)
}

func Listen(name string, handle Handle) {
	defaultInstance.Listen(name, handle)
}

// Listen
func (instance *dispatcher) Listen(name string, handle Handle) {
	instance.Register(name, Listener{Handle: handle})
}

// Remove remove listener from dispatcher
func Remove(name string) {
	defaultInstance.Remove(name)
}
func (instance *dispatcher) Remove(name string) {
	delete(instance.events, name)
}

// Has check dispatcher is or not include the name
func Has(name string) bool {
	return defaultInstance.Has(name)
}

// Has check dispatcher is or not include the name
func (instance *dispatcher) Has(name string) bool {
	_, ok := instance.events[name]
	return ok
}

// Trigger 触发事件
func Trigger(name string, params interface{}) error {
	return defaultInstance.Trigger(name, params)
}

// Trigger 触发事件
func (instance *dispatcher) Trigger(name string, params interface{}) error {
	if !Has(name) {
		return fmt.Errorf("event: %s not exist", name)
	}

	e := Event{
		Name:   name,
		Params: params,
	}

	for _, listener := range instance.events[name] {
		if listener.Mode == Async {
			go listener.Handle(e)
		} else {
			listener.Handle(e)
		}

	}
	return nil
}

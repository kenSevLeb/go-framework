// event 事件模块
package event

// Event
type Event struct {
	// event name
	Name string
	// event params
	Params interface{}
}

// Listener process event
type Listener struct {
	//
	Mode Mode

	// handle event
	Handle Handle
}

type Handle func(ev Event)

const (
	BeforeHttpStart       = "http.start.before"
	AfterHttpStart        = "http.start.after"
	BeforeHttpShutdown    = "http.shutdown.before"
	AfterHttpShutdown     = "http.shutdown.after"
	BeforeDatabaseConnect = "database.connect.before"
	AfterDatabaseConnect  = "database.connect.after"
)

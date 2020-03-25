package iotmaker_db_mongodb_config

type ServiceExecutor int

var serviceExecutors = [...]string{
	"",
	"synchronous",
	"adaptive",
}

func (el ServiceExecutor) String() string {
	return `"` + serviceExecutors[el] + `"`
}

const (
	//The mongos or mongod uses synchronous networking and manages its networking thread
	//pool on a per connection basis. Previous versions of MongoDB managed threads in this
	//way.
	KServiceExecutorSynchronous ServiceExecutor = iota + 1

	//The mongos or mongod uses the new experimental asynchronous networking mode with an
	//adaptive thread pool which manages threads on a per request basis. This mode should
	//have more consistent performance and use less resources when there are more inactive
	//connections than database requests.
	KServiceExecutorAdaptive
)

package iotmaker_db_mongodb_config

type SystemLogDestination int

var systemLogDestinations = [...]string{
	"",
	"file",
	"syslog",
}

func (el SystemLogDestination) String() string {
	return `"` + systemLogDestinations[el] + `"`
}

const (
	//you must also specify systemLog.path
	KSystemLogDestinationFile SystemLogDestination = iota + 1

	//WARNING:
	//    The syslog daemon generates timestamps when it logs a message, not when MongoDB
	//    issues the message.
	//    This can lead to misleading timestamps for log entries, especially when the
	//    system is under heavy load.
	//    We recommend using the file option for production systems to ensure accurate
	//    timestamps.
	KSystemLogDestinationSyslog
)

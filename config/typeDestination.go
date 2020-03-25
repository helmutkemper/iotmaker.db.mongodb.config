package iotmaker_db_mongodb_config

type Destination int

var destinations = [...]string{
	"",
	"syslog",
	"console",
	"file",
}

func (el Destination) String() string {
	return `"` + destinations[el] + `"`
}

const (
	//Output the audit events to syslog in JSON format. Not available on Windows. Audit
	//messages have a syslog severity level of info and a facility level of user.
	//
	//The syslog message limit can result in the truncation of audit messages. The auditing
	//system will neither detect the truncation nor error upon its occurrence.
	KDestinationSyslog Destination = iota + 1

	//Output the audit events to stdout in JSON format.
	KDestinationConsole

	//Output the audit events to the file specified in auditLog.path in the format
	//specified in auditLog.format.
	KDestinationFile
)

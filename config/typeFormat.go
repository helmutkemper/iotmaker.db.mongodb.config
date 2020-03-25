package iotmaker_db_mongodb_config

type Format int

var formats = [...]string{
	"",
	"JSON",
	"BSON",
}

func (el Format) String() string {
	return `"` + formats[el] + `"`
}

const (
	//Output the audit events in JSON format to the file specified in auditLog.path.
	KFormatJSon Format = iota + 1

	//Output the audit events in BSON binary format to the file specified in auditLog.path.
	KFormatBSon
)

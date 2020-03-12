package iotmaker_db_mongodb_config

type LogRotate int

var logRotates = [...]string{
	"",
	"rename.",
	"reopen",
}

func (el LogRotate) String() string {
	return logRotates[el]
}

const (
	//renames the log file.
	KLogRotateRename LogRotate = iota + 1

	//reopen closes and reopens the log file following the typical Linux/Unix log rotate behavior. Use reopen when using the Linux/Unix logrotate utility to avoid log loss.
	//
	//If you specify reopen, you must also set systemLog.logAppend to true.
	KLogRotateReopen
)

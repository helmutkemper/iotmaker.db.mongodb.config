package iotmaker_db_mongodb_config

type TimeStampFormat int

var timeStampFormats = [...]string{
	"",
	"ctime",
	"iso8601-utc",
	"iso8601-local",
}

func (el TimeStampFormat) String() string {
	return timeStampFormats[el]
}

const (
	//Displays timestamps as Wed Dec 31 18:17:54.811.
	KTimeStampFormatCTime TimeStampFormat = iota + 1

	//Displays timestamps in Coordinated Universal Time (UTC) in the ISO-8601 format. For example, for New York at the start of the Epoch: 1970-01-01T00:00:00.000Z
	KTimeStampFormatISO8601UTC

	//Displays timestamps in local time in the ISO-8601 format. For example, for New York at the start of the Epoch: 1969-12-31T19:00:00.000-0500
	KTimeStampFormatISO8601Local
)

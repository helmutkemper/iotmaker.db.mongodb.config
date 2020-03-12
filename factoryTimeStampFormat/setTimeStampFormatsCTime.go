package factoryTimeStampFormat

import iotmaker_db_mongodb_config "github.com/helmut.kemper/iotmaker.db.mongodb.config/config"

//Displays timestamps as Wed Dec 31 18:17:54.811.
func SetTimeStampFormatsCTime(configuration *iotmaker_db_mongodb_config.Configuration) {
	configuration.TimeStampFormat = iotmaker_db_mongodb_config.KTimeStampFormatCTime
}

//Displays timestamps in Coordinated Universal Time (UTC) in the ISO-8601 format. For example, for New York at the start of the Epoch: 1970-01-01T00:00:00.000Z
func SetTimeStampFormatsISO8601UTC(configuration *iotmaker_db_mongodb_config.Configuration) {
	configuration.TimeStampFormat = iotmaker_db_mongodb_config.KTimeStampFormatISO8601UTC
}

//Displays timestamps in local time in the ISO-8601 format. For example, for New York at the start of the Epoch: 1969-12-31T19:00:00.000-0500
func SetTimeStampFormatsISO8601Local(configuration *iotmaker_db_mongodb_config.Configuration) {
	configuration.TimeStampFormat = iotmaker_db_mongodb_config.KTimeStampFormatISO8601Local
}

package factoryMongoDBConfigTimeStampFormat

import iotmaker_db_mongodb_config "github.com/helmutkemper/iotmaker.db.mongodb.config/config"

//Displays timestamps in local time in the ISO-8601 format. For example, for New York at the start of the Epoch: 1969-12-31T19:00:00.000-0500
func SetTimeStampFormatsISO8601Local(configuration *iotmaker_db_mongodb_config.Configuration) {
	configuration.SystemLog.TimeStampFormat = iotmaker_db_mongodb_config.KTimeStampFormatISO8601Local
}

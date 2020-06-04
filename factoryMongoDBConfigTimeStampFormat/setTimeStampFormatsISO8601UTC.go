package factoryMongoDBConfigTimeStampFormat

import iotmaker_db_mongodb_config "github.com/helmutkemper/iotmaker.db.mongodb.config/config"

//Displays timestamps in Coordinated Universal Time (UTC) in the ISO-8601 format. For example, for New York at the start of the Epoch: 1970-01-01T00:00:00.000Z
func SetTimeStampFormatsISO8601UTC(configuration *iotmaker_db_mongodb_config.Configuration) {
	configuration.SystemLog.TimeStampFormat = iotmaker_db_mongodb_config.KTimeStampFormatISO8601UTC
}

package factoryMongoDBConfigTimeStampFormat

import iotmaker_db_mongodb_config "github.com/helmutkemper/iotmaker.db.mongodb.config/config"

//Displays timestamps as Wed Dec 31 18:17:54.811.
func SetTimeStampFormatsCTime(configuration *iotmaker_db_mongodb_config.Configuration) {
	configuration.SystemLog.TimeStampFormat = iotmaker_db_mongodb_config.KTimeStampFormatCTime
}

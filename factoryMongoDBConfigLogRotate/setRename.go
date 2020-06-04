package factoryMongoDBConfigLogRotate

import iotmaker_db_mongodb_config "github.com/helmutkemper/iotmaker.db.mongodb.config/config"

//renames the log file.
func SetRename(configuration *iotmaker_db_mongodb_config.Configuration) {
	configuration.SystemLog.LogRotate = iotmaker_db_mongodb_config.KLogRotateRename
}

//reopen closes and reopens the log file following the typical Linux/Unix log rotate behavior. Use reopen when using the Linux/Unix logrotate utility to avoid log loss.
//
//If you specify reopen, you must also set systemLog.logAppend to true.
func SetReopen(configuration *iotmaker_db_mongodb_config.Configuration) {
	configuration.SystemLog.LogRotate = iotmaker_db_mongodb_config.KLogRotateReopen
	configuration.SystemLog.LogAppend = iotmaker_db_mongodb_config.TRUE
}

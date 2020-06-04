package factoryMongoDBConfigFreeState

import iotmaker_db_mongodb_config "github.com/helmutkemper/iotmaker.db.mongodb.config/config"

//Default. You can enable or disable free monitoring during runtime.
//
//To enable or disable free monitoring during runtime, see db.enableFreeMonitoring() and db.disableFreeMonitoring().
//
//To enable or disable free monitoring during runtime when running with access control, users must have required privileges. See db.enableFreeMonitoring() and db.disableFreeMonitoring() for details.
func SetRuntime(configuration *iotmaker_db_mongodb_config.Configuration) {
	configuration.Cloud.Monitoring.Free.State = iotmaker_db_mongodb_config.KFreeStateRuntime
}

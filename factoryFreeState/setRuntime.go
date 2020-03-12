package factoryFreeState

import iotmaker_db_mongodb_config "github.com/helmut.kemper/iotmaker.db.mongodb.config/config"

//Default. You can enable or disable free monitoring during runtime.
//
//To enable or disable free monitoring during runtime, see db.enableFreeMonitoring() and db.disableFreeMonitoring().
//
//To enable or disable free monitoring during runtime when running with access control, users must have required privileges. See db.enableFreeMonitoring() and db.disableFreeMonitoring() for details.
func SetRuntime(configuration *iotmaker_db_mongodb_config.Configuration) {
	configuration.Cloud.Monitoring.Free.State = iotmaker_db_mongodb_config.KFreeStateRuntime
}

//Enables free monitoring at startup; i.e. registers for free monitoring. When enabled at startup, you cannot disable free monitoring during runtime.
func SetOn(configuration *iotmaker_db_mongodb_config.Configuration) {
	configuration.Cloud.Monitoring.Free.State = iotmaker_db_mongodb_config.KFreeStateOn
}

//Disables free monitoring at startup, regardless of whether you have previously registered for free monitoring. When disabled at startup, you cannot enable free monitoring during runtime.
func SetOff(configuration *iotmaker_db_mongodb_config.Configuration) {
	configuration.Cloud.Monitoring.Free.State = iotmaker_db_mongodb_config.KFreeStateOff
}

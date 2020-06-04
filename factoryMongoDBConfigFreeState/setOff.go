package factoryMongoDBConfigFreeState

import iotmaker_db_mongodb_config "github.com/helmutkemper/iotmaker.db.mongodb.config/config"

//Disables free monitoring at startup, regardless of whether you have previously registered for free monitoring. When disabled at startup, you cannot enable free monitoring during runtime.
func SetOff(configuration *iotmaker_db_mongodb_config.Configuration) {
	configuration.Cloud.Monitoring.Free.State = iotmaker_db_mongodb_config.KFreeStateOff
}

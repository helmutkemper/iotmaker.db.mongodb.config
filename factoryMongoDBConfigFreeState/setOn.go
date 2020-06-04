package factoryMongoDBConfigFreeState

import iotmaker_db_mongodb_config "github.com/helmutkemper/iotmaker.db.mongodb.config/config"

//Enables free monitoring at startup; i.e. registers for free monitoring. When enabled at startup, you cannot disable free monitoring during runtime.
func SetOn(configuration *iotmaker_db_mongodb_config.Configuration) {
	configuration.Cloud.Monitoring.Free.State = iotmaker_db_mongodb_config.KFreeStateOn
}

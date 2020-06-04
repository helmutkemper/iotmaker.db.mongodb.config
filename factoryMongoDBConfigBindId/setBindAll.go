package factoryMongoDBConfigBindId

import iotmaker_db_mongodb_config "github.com/helmutkemper/iotmaker.db.mongodb.config/config"

func SetBindAll(configuration *iotmaker_db_mongodb_config.Configuration, value bool) {
	configuration.Net.BindIpAll = iotmaker_db_mongodb_config.FALSE
	if value == true {
		configuration.Net.BindIpAll = iotmaker_db_mongodb_config.TRUE
	}
}

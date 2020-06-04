package factoryMongoDBConfigNetPort

import iotmaker_db_mongodb_config "github.com/helmutkemper/iotmaker.db.mongodb.config/config"

func SetMongodIsAConfigServerMember(configuration *iotmaker_db_mongodb_config.Configuration) {
	configuration.Net.Port = iotmaker_db_mongodb_config.KNetPortMongodIsAConfigServerMember
}

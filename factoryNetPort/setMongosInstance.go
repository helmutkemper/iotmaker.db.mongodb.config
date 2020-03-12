package factoryNetPort

import iotmaker_db_mongodb_config "github.com/helmut.kemper/iotmaker.db.mongodb.config/config"

func SetMongodInstance(configuration *iotmaker_db_mongodb_config.Configuration) {
	configuration.Net.Port = iotmaker_db_mongodb_config.KNetPortMongodOrMongosInstance
}

func SetMongodIsAShardMember(configuration *iotmaker_db_mongodb_config.Configuration) {
	configuration.Net.Port = iotmaker_db_mongodb_config.KNetPortMongodIsAShardMember
}

func SetMongodIsAConfigServerMember(configuration *iotmaker_db_mongodb_config.Configuration) {
	configuration.Net.Port = iotmaker_db_mongodb_config.KNetPortMongodIsAConfigServerMember
}

package factoryMongoDBConfigReplicaSet

import iotmaker_db_mongodb_config "github.com/helmutkemper/iotmaker.db.mongodb.config/config"

func SetServerAsReplicaSet(configuration *iotmaker_db_mongodb_config.Configuration, name, dbPath string, bindIpList []string) {
	configuration.Storage.DbPath = dbPath
	configuration.Storage.DirectoryPerDB = iotmaker_db_mongodb_config.TRUE
	configuration.Storage.Journal.Enabled = iotmaker_db_mongodb_config.TRUE
	configuration.Storage.Journal.CommitIntervalMs = 100

	configuration.Replication.ReplSetName = name
	configuration.Replication.EnableMajorityReadConcern = iotmaker_db_mongodb_config.TRUE

	configuration.Net.Port = iotmaker_db_mongodb_config.KNetPortMongodOrMongosInstance
	configuration.Net.BindIp = bindIpList
}

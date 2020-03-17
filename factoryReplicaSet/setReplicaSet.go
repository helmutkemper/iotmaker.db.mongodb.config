package factoryReplicaSet

import iotmaker_db_mongodb_config "github.com/helmut.kemper/iotmaker.db.mongodb.config/config"

func SetServerAsReplicaSet(configuration *iotmaker_db_mongodb_config.Configuration, name, dbPath string, bindIpList []string) {
	configuration.Storage.DbPath = dbPath
	configuration.Storage.DirectoryPerDB = true
	configuration.Storage.Journal.Enabled = true
	configuration.Storage.Journal.CommitIntervalMs = 100

	configuration.Replication.ReplSetName = name
	configuration.Replication.EnableMajorityReadConcern = true

	configuration.Net.Port = iotmaker_db_mongodb_config.KNetPortMongodOrMongosInstance
	configuration.Net.BindIp = bindIpList
}

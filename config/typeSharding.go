package iotmaker_db_mongodb_config

type Sharding struct {
	//The role that the mongod instance has in the sharded cluster. Set this setting to one of the following:
	//
	//Value			Description
	//configsvr		Start this instance as a config server. The instance starts on port 27019 by default.
	//shardsvr		Start this instance as a shard. The instance starts on port 27018 by default.
	//NOTE
	//
	//Setting sharding.clusterRole requires the mongod instance to be running with replication. To deploy the instance as a replica set member, use the replSetName setting and specify the name of the replica set.
	//
	//The sharding.clusterRole setting is available only for mongod.
	ClusterRole string `yaml:"clusterRole"`

	//Changed in version 3.2: Starting in 3.2, MongoDB uses false as the default.
	//
	//During chunk migration, a shard does not save documents migrated from the shard.
	archiveMovedChunks LogicBoolean `yaml:"archiveMovedChunks"`
}

package iotmaker_db_mongodb_config

type ReplicationLog struct {
	//Default: 0
	//
	//The log message verbosity level for components related to replication. See REPL components.
	//
	//The verbosity level can range from 0 to 5:
	//    0 is the MongoDB’s default log verbosity level, to include Informational messages.
	//    1 to 5 increases the verbosity level to include Debug messages.
	Verbosity Verbosity `yaml:"verbosity"`

	Election    ElectionLog    `yaml:"Election"`
	Heartbeats  HeartbeatsLog  `yaml:"heartbeats"`
	InitialSync InitialSyncLog `yaml:"initialSync"`
	Rollback    RollbackLog    `yaml:"rollback"`
	Sharding    ShardingLog    `yaml:"sharding"`
	Storage     StorageLog     `yaml:"storage"`
}

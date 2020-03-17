package iotmaker_db_mongodb_config

type ReplicationLog struct {
	//Default: 0
	//
	//The log message verbosity level for components related to replication. See REPL components.
	//
	//The verbosity level can range from 0 to 5:
	//    0 is the MongoDB’s default log verbosity level, to include Informational messages.
	//    1 to 5 increases the verbosity level to include Debug messages.
	Verbosity Verbosity `yaml:"verbosity"` //ok

	Election    ElectionLog    `yaml:"Election"`    //ok
	Heartbeats  HeartbeatsLog  `yaml:"heartbeats"`  //ok
	InitialSync InitialSyncLog `yaml:"initialSync"` //ok
	Rollback    RollbackLog    `yaml:"rollback"`    //ok
	Sharding    ShardingLog    `yaml:"sharding"`    //ok
	Storage     StorageLog     `yaml:"storage"`     //ok
}

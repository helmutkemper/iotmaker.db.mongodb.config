package iotmaker_db_mongodb_config

type InitialSyncLog struct {
	//Default: 0
	//
	// > New in version 4.2.
	//
	//The log message verbosity level for components related to initialSync. See INITSYNC components.
	//
	//If systemLog.component.replication.initialSync.verbosity is unset, systemLog.component.replication.verbosity level also applies to initialSync components.
	//
	//The verbosity level can range from 0 to 5:
	//    0 is the MongoDB’s default log verbosity level, to include Informational messages.
	//    1 to 5 increases the verbosity level to include Debug messages.
	Verbosity Verbosity `yaml:"verbosity"`
}

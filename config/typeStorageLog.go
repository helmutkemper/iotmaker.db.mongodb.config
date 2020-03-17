package iotmaker_db_mongodb_config

type StorageLog struct {
	//Default: 0
	//
	//The log message verbosity level for components related to storage. See STORAGE components.
	//
	//If systemLog.component.storage.journal.verbosity is unset, systemLog.component.storage.verbosity level also applies to journaling components.
	//
	//The verbosity level can range from 0 to 5:
	//    0 is the MongoDB’s default log verbosity level, to include Informational messages.
	//    1 to 5 increases the verbosity level to include Debug messages.
	Verbosity Verbosity `yaml:"verbosity"`

	Journal  JournalLog  `yaml:"-"`
	Recovery RecoveryLog `yaml:"-"`
}

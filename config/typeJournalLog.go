package iotmaker_db_mongodb_config

type JournalLog struct {
	//Default: 0
	//
	//The log message verbosity level for components related to journaling. See JOURNAL components.
	//
	//If systemLog.component.storage.journal.verbosity is unset, the journaling components have the same verbosity level as the parent storage components: i.e. either the systemLog.component.storage.verbosity level if set or the default verbosity level.
	//
	//The verbosity level can range from 0 to 5:
	//    0 is the MongoDB’s default log verbosity level, to include Informational messages.
	//    1 to 5 increases the verbosity level to include Debug messages.
	Verbosity Verbosity `yaml:"verbosity"`
}

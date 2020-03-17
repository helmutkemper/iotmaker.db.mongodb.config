package iotmaker_db_mongodb_config

type QueryLog struct {
	// Default: 0
	//
	//The log message verbosity level for components related to query operations. See QUERY components.
	//
	//The verbosity level can range from 0 to 5:
	//    0 is the MongoDB’s default log verbosity level, to include Informational messages.
	//    1 to 5 increases the verbosity level to include Debug messages.
	Verbosity Verbosity `yaml:"verbosity"`
}

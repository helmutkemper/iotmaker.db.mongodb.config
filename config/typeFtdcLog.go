package iotmaker_db_mongodb_config

type FtdcLog struct {
	//Default: 0
	//
	// > New in version 3.2.
	//
	//The log message verbosity level for components related to diagnostic data collection
	//operations. See FTDC components.
	//
	//The verbosity level can range from 0 to 5:
	//    0 is the MongoDB’s default log verbosity level, to include Informational
	//      messages.
	//    1 to 5 increases the verbosity level to include Debug messages.
	Verbosity Verbosity `yaml:"verbosity"`
}

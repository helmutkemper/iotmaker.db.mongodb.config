package iotmaker_db_mongodb_config

type IndexConfig struct {
	//Default: true
	//
	//Enables or disables prefix compression for index data.
	//
	//Specify true for storage.wiredTiger.indexConfig.prefixCompression to enable prefix
	//compression for index data, or false to disable prefix compression for index data.
	//
	//The storage.wiredTiger.indexConfig.prefixCompression setting affects all indexes
	//created. If you change the value of storage.wiredTiger.indexConfig.prefixCompression
	//on an existing MongoDB deployment, all new indexes will use prefix compression.
	//Existing indexes are not affected.
	PrefixCompression LogicBoolean `yaml:"prefixCompression"`
}

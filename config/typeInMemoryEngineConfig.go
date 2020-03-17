package iotmaker_db_mongodb_config

type InMemoryEngineConfig struct {
	//Default: 50% of physical RAM less 1 GB
	//
	//Changed in version 3.4: Values can range from 256MB to 10TB and can be a float.
	//
	//Maximum amount of memory to allocate for in-memory storage engine data, including indexes, oplog if the mongod is part of replica set, replica set or sharded cluster metadata, etc.
	//
	//By default, the in-memory storage engine uses 50% of physical RAM minus 1 GB.
	//
	//ENTERPRISE FEATURE: Available in MongoDB Enterprise only.
	InMemorySizeGB float64 `yaml:"inMemorySizeGB"`
}

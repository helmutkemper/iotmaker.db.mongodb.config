package iotmaker_db_mongodb_config

type NetPort int

const (
	KNetPortMongodOrMongosInstance      NetPort = 27017
	KNetPortMongodIsAShardMember        NetPort = 27018
	KNetPortMongodIsAConfigServerMember NetPort = 27019
)

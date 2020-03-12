package iotmaker_db_mongodb_config

type NetPort int

const (
	KNetPortMongodOrMongosInstance      = 27017
	KNetPortMongodIsAShardMember        = 27018
	KNetPortMongodIsAConfigServerMember = 27019
)

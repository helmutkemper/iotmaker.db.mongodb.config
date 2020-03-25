package iotmaker_db_mongodb_config

import "strconv"

type NetPort int

func (el *NetPort) String() string {
	return strconv.FormatInt(int64(*el), 10)
}

const (
	KNetPortMongodOrMongosInstance      NetPort = 27017
	KNetPortMongodIsAShardMember        NetPort = 27018
	KNetPortMongodIsAConfigServerMember NetPort = 27019
)

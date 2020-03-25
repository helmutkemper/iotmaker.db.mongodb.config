package iotmaker_db_mongodb_config

type ShardingClusterRole int

var shardingClusterRoles = [...]string{
	"",
	"configsvr",
	"shardsvr",
}

func (el ShardingClusterRole) String() string {
	return `"` + shardingClusterRoles[el] + `"`
}

const (
	//Start this instance as a config server. The instance starts on port 27019 by default.
	KShardingClusterRoleASConfigServer ShardingClusterRole = iota + 1

	//Start this instance as a shard. The instance starts on port 27018 by default.
	KShardingClusterRoleAsShardServer
)

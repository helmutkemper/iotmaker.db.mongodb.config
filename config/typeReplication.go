package iotmaker_db_mongodb_config

type Replication struct {
	//The maximum size in megabytes for the replication operation log (i.e., the oplog).
	//
	//NOTE: Starting in MongoDB 4.0, the oplog can grow past its configured size limit to
	//      avoid deleting the majority commit point.
	//
	//By default, the mongod process creates an oplog based on the maximum amount of space
	//available. For 64-bit systems, the oplog is typically 5% of available disk space.
	//
	//Once the mongod has created the oplog for the first time, changing the
	//replication.oplogSizeMB option will not affect the size of the oplog.
	//
	//To change the oplog size of a running replica set member, use the replSetResizeOplog
	//administrative command. replSetResizeOplog enables you to resize the oplog
	//dynamically without restarting the mongod process.
	//
	//See Oplog Size for more information.
	//
	//The replication.oplogSizeMB setting is available only for mongod.
	OpLogSizeMB int64 `yaml:"-"`

	//The name of the replica set that the mongod is part of. All hosts in the replica set
	//must have the same set name.
	//
	//If your application connects to more than one replica set, each set should have a
	//distinct name. Some drivers group replica set connections by replica set name.
	//
	//The replication.replSetName setting is available only for mongod.
	//
	//Starting in MongoDB 4.0:
	//
	//The setting replication.replSetName cannot be used in conjunction with
	//storage.indexBuildRetry.
	//For the WiredTiger storage engine, storage.journal.enabled: false cannot be used in
	//conjunction with replication.replSetName.
	ReplSetName string `yaml:"-"`

	//Starting in MongoDB 3.6, MongoDB enables support for "majority" read concern by
	//default.
	//
	//You can disable read concern "majority" to prevent the storage cache pressure from
	//immobilizing a deployment with a three-member primary-secondary-arbiter (PSA)
	//architecture.
	//For more information about disabling read concern "majority", see Disable Read
	//Concern Majority.
	//
	//To disable, set replication.enableMajorityReadConcern to false.
	//replication.enableMajorityReadConcern has no effect for MongoDB versions: 4.0.0,
	//4.0.1, 4.0.2, 3.6.0.
	//
	//IMPORTANT: In general, avoid disabling "majority" read concern unless necessary.
	//However, if you have a three-member replica set with a primary-secondary-arbiter
	//(PSA) architecture or a sharded cluster with a three-member PSA shards, disable to
	//prevent the storage cache pressure from immobilizing the deployment.
	//
	//    Disabling "majority" read concern affects support for transactions on sharded
	//    clusters. Specifically:
	//
	//    A transaction cannot use read concern "snapshot" if the transaction involves a
	//    shard that has disabled read concern “majority”.
	//
	//    A transaction that writes to multiple shards errors if any of the transaction’s
	//    read or write operations involves a shard that has disabled read concern
	//    "majority".
	//
	//    However, it does not affect transactions on replica sets. For transactions on
	//    replica sets, you can specify read concern "majority" (or "snapshot" or "local" )
	//    for multi-document transactions even if read concern "majority" is disabled.
	//
	//    Disabling "majority" read concern disables support for Change Streams for MongoDB
	//    4.0 and earlier. For MongoDB 4.2+, disabling read concern "majority" has no
	//    effect on change streams availability.
	EnableMajorityReadConcern LogicBoolean
}

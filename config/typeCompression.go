package iotmaker_db_mongodb_config

type Compression struct {
	//Default: snappy,zstd,zlib
	//
	//New in version 3.4.
	//
	//Specifies the default compressor(s) to use for communication between this mongod or
	//mongos instance and:
	//
	//* other members of the deployment if the instance is part of a replica set or a
	//sharded cluster
	//
	//* a mongo shell
	//
	//* drivers that support the OP_COMPRESSED message format.
	//
	//MongoDB supports the following compressors:
	//
	//    Value  |  Description
	//    ---------------------------------------------------------------------------------
	//    snappy |
	//    ---------------------------------------------------------------------------------
	//    zlib   | Available starting in MongoDB 3.6
	//    ---------------------------------------------------------------------------------
	//    zstd   | Available starting in MongoDB 4.2
	//
	//In versions 3.6 and 4.0, mongod and mongos enable network compression by default with
	//snappy as the compressor.
	//
	//Starting in version 4.2, mongod and mongos instances default to both snappy,zstd,zlib
	//compressors, in that order.
	//
	//To disable network compression, set the value to disabled.
	//
	//IMPORTANT: Messages are compressed when both parties enable network compression.
	//Otherwise, messages between the parties are uncompressed.
	//
	//If you specify multiple compressors, then the order in which you list the compressors
	//matter as well as the communication initiator. For example, if a mongo shell
	//specifies the following network compressors zlib,snappy and the mongod specifies
	//snappy, zlib, messages between mongo shell and mongod uses zlib.
	//
	//If the parties do not share at least one common compressor, messages between the
	//parties are uncompressed. For example, if a mongo shell specifies the network
	//compressor zlib and mongod specifies snappy, messages between mongo shell and mongod
	//are not compressed.
	Compressors Compressor `yaml:"compressors"`
}

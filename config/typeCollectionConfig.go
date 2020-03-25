package iotmaker_db_mongodb_config

type CollectionConfig struct {
	//Default: snappy
	//
	//Specifies the default compression for collection data. You can override this on a
	//per-collection basis when creating collections.
	//
	//Available compressors are:
	//
	//    Value  |  Description
	//    ---------------------------------------------------------------------------------
	//    none   |
	//    ---------------------------------------------------------------------------------
	//    snappy |
	//    ---------------------------------------------------------------------------------
	//    zlib   |
	//    ---------------------------------------------------------------------------------
	//    zstd   | Available starting MongoDB 4.2
	//
	//storage.wiredTiger.collectionConfig.blockCompressor affects all collections created.
	//
	//If you change the value of storage.wiredTiger.collectionConfig.blockCompressor on an
	//existing MongoDB deployment, all new collections will use the specified compressor.
	//
	//Existing collections will continue to use the compressor specified when they were
	//created, or the default compressor at that time.
	BlockCompressor string `yaml:"blockCompressor"`
}

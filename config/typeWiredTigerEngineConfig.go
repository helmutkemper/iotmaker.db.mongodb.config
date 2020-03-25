package iotmaker_db_mongodb_config

type WiredTigerEngineConfig struct {
	//Defines the maximum size of the internal cache that WiredTiger will use for all data. The memory consumed by an index build (see maxIndexBuildMemoryUsageMegabytes) is separate from the WiredTiger cache memory.
	//
	//Values can range from 0.25 GB to 10000 GB.
	//
	//Starting in MongoDB 3.4, the default WiredTiger internal cache size is the larger of either:
	//
	//50% of (RAM - 1 GB), or
	//256 MB.
	//For example, on a system with a total of 4GB of RAM the WiredTiger cache will use 1.5GB of RAM (0.5 * (4 GB - 1 GB) = 1.5 GB). Conversely, a system with a total of 1.25 GB of RAM will allocate 256 MB to the WiredTiger cache because that is more than half of the total RAM minus one gigabyte (0.5 * (1.25 GB - 1 GB) = 128 MB < 256 MB).
	//
	//NOTE
	//
	//In some instances, such as when running in a container, the database can have memory constraints that are lower than the total system memory. In such instances, this memory limit, rather than the total system memory, is used as the maximum RAM available.
	//
	//To see the memory limit, see hostInfo.system.memLimitMB.
	//
	//Avoid increasing the WiredTiger internal cache size above its default value.
	//
	//With WiredTiger, MongoDB utilizes both the WiredTiger internal cache and the filesystem cache.
	//
	//Via the filesystem cache, MongoDB automatically uses all free memory that is not used by the WiredTiger cache or by other processes.
	//
	//NOTE
	//
	//The storage.wiredTiger.engineConfig.cacheSizeGB limits the size of the WiredTiger internal cache. The operating system will use the available free memory for filesystem cache, which allows the compressed MongoDB data files to stay in memory. In addition, the operating system will use any free RAM to buffer file system blocks and file system cache.
	//
	//To accommodate the additional consumers of RAM, you may have to decrease WiredTiger internal cache size.
	//
	//The default WiredTiger internal cache size value assumes that there is a single mongod instance per machine. If a single machine contains multiple MongoDB instances, then you should decrease the setting to accommodate the other mongod instances.
	//
	//If you run mongod in a container (e.g. lxc, cgroups, Docker, etc.) that does not have access to all of the RAM available in a system, you must set storage.wiredTiger.engineConfig.cacheSizeGB to a value less than the amount of RAM available in the container. The exact amount depends on the other processes running in the container. See memLimitMB.
	CacheSizeGB float64 `yaml:"cacheSizeGB"`

	//Default: snappy
	//
	//Specifies the type of compression to use to compress WiredTiger journal data.
	//
	//Available compressors are:
	//
	//none
	//snappy	A compression/decompression library designed to balance efficient computation requirements with reasonable compression rates. Snappy is the default compression library for MongoDB’s use of WiredTiger. See Snappy and the WiredTiger compression documentation for more information.
	//			https://google.github.io/snappy/
	//zlib		A data compression library that provides higher compression rates at the cost of more CPU, compared to MongoDB’s use of snappy. You can configure WiredTiger to use zlib as its compression library. See http://www.zlib.net and the WiredTiger compression documentation for more information.
	//			http://www.zlib.net/
	//zstd 		(Available starting in MongoDB 4.2) A data compression library that provides higher compression rates and lower CPU usage when compared to zlib.
	//			http://www.zlib.net/
	JournalCompressor string `yaml:"journalCompressor"`

	//Default: false
	//
	//When storage.wiredTiger.engineConfig.directoryForIndexes is true, mongod stores indexes and collections in separate subdirectories under the data (i.e. storage.dbPath) directory. Specifically, mongod stores the indexes in a subdirectory named index and the collection data in a subdirectory named collection.
	//
	//By using a symbolic link, you can specify a different location for the indexes. Specifically, when mongod instance is not running, move the index subdirectory to the destination and create a symbolic link named index under the data directory to the new destination.
	DirectoryForIndexes LogicBoolean `yaml:"directoryForIndexes"`

	//Specifies the maximum size (in GB) for the “lookaside (or cache overflow) table” file WiredTigerLAS.wt.
	//
	//The setting can accept the following values:
	//
	//Value	Description
	//0	The default value. If set to 0, the file size is unbounded.
	//number >= 0.1	The maximum size (in GB). If the WiredTigerLAS.wt file exceeds this size, mongod exits with a fatal assertion. You can clear the WiredTigerLAS.wt file and restart mongod.
	//To change the maximum size during runtime, use the wiredTigerMaxCacheOverflowSizeGB parameter.
	//
	//Available starting in MongoDB 4.2.1 (and 4.0.12)
	MaxCacheOverflowFileSizeGB float64 `yaml:"maxCacheOverflowFileSizeGB"`

	CollectionConfig CollectionConfig `yaml:"collectionConfig"`
}

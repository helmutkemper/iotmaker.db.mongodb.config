package iotmaker_db_mongodb_config

type Storage struct {
	//Default:
	//
	///data/db on Linux and macOS
	//\data\db on Windows
	//The directory where the mongod instance stores its data.
	//
	//The storage.dbPath setting is available only for mongod.
	//
	//CONFIGURATION FILES:
	//The default mongod.conf configuration file included with package manager installations uses the following platform-specific default values for storage.dbPath:
	//
	//Platform						Package Manager			Default storage.dbPath
	//RHEL / CentOS and Amazon		yum						/var/lib/mongo
	//SUSE							zypper					/var/lib/mongo
	//Ubuntu and Debian				apt						/var/lib/mongodb
	//macOS							brew					/usr/local/var/mongodb
	//The Linux package init scripts do not expect storage.dbPath to change from the defaults. If you use the Linux packages and change storage.dbPath, you will have to use your own init scripts and disable the built-in scripts.
	DbPath string `yaml:"dbPath"`

	//Default: true
	//
	//Specifies whether mongod rebuilds incomplete indexes on the next start up. This applies in cases where mongod restarts after it has shut down or stopped in the middle of an index build. In such cases, mongod always removes any incomplete indexes, and then, by default, attempts to rebuild them. To stop mongod from rebuilding indexes, set this option to false.
	//
	//Changed in version 4.0: The setting storage.indexBuildRetry cannot be used in conjunction with replication.replSetName.
	//
	//The storage.indexBuildRetry setting is available only for mongod.
	//
	//Not available for mongod instances that use the in-memory storage engine.
	IndexBuildRetry LogicBoolean `yaml:"indexBuildRetry"`

	Journal Journal `yaml:"journal"`

	//Default: false
	//
	//When true, MongoDB uses a separate directory to store data for each database. The directories are under the storage.dbPath directory, and each subdirectory name corresponds to the database name.
	//
	//The storage.directoryPerDB setting is available only for mongod.
	//
	//Not available for mongod instances that use the in-memory storage engine.
	//
	//To change the storage.directoryPerDB option for existing deployments:
	//
	//For standalone instances:
	//1.Use mongodump on the existing mongod instance to generate a backup.
	//2.Stop the mongod instance.
	//3.Add the storage.directoryPerDB value and configure a new data directory
	//4.Restart the mongod instance.
	//5.Use mongorestore to populate the new data directory.
	//For replica sets:
	//1.Stop a secondary member.
	//2.Add the storage.directoryPerDB value and configure a new data directory to that secondary member.
	//3.Restart that secondary.
	//4.Use initial sync to populate the new data directory.
	//5.Update remaining secondaries in the same fashion.
	//6.Step down the primary, and update the stepped-down member in the same fashion.
	//todo: make true
	DirectoryPerDB LogicBoolean `yaml:"directoryPerDB"`

	//Default: 60
	//
	//The amount of time that can pass before MongoDB flushes data to the data files via an fsync operation.
	//
	//Do not set this value on production systems. In almost every situation, you should use the default setting.
	//
	//WARNING: If you set storage.syncPeriodSecs to 0, MongoDB will not sync the memory mapped files to disk.
	//
	//The mongod process writes data very quickly to the journal and lazily to the data files. storage.syncPeriodSecs has no effect on the journal files or journaling, but if storage.syncPeriodSecs is set to 0 the journal will eventually consume all available disk space. If you set storage.syncPeriodSecs to 0 for testing purposes, you should also set --nojournal to true.
	//
	//The serverStatus command reports the background flush thread’s status via the backgroundFlushing field.
	//
	//The storage.syncPeriodSecs setting is available only for mongod.
	//
	//Not available for mongod instances that use the in-memory storage engine.
	SyncPeriodSecs int `yaml:"syncPeriodSecs"`

	//Default: wiredTiger
	//
	//NOTE: Starting in version 4.2, MongoDB removes the deprecated MMAPv1 storage engine.
	//
	//The storage engine for the mongod database. Available values include:
	//
	//Value	Description
	//wiredTiger	To specify the WiredTiger Storage Engine.
	//inMemory
	//To specify the In-Memory Storage Engine.
	//
	//New in version 3.2: Available in MongoDB Enterprise only.
	//
	//If you attempt to start a mongod with a storage.dbPath that contains data files produced by a storage engine other than the one specified by storage.engine, mongod will refuse to start.
	Engine string `yaml:"engine"`

	WiredTiger WiredTiger `yaml:"wiredTiger"`
	InMemory   InMemory   `yaml:"inMemory"`
}

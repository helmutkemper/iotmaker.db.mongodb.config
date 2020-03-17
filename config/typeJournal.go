package iotmaker_db_mongodb_config

type Journal struct {
	//Default: true on 64-bit systems, false on 32-bit systems
	//
	//Enable or disable the durability journal to ensure data files remain valid and recoverable. This option applies only when you specify the storage.dbPath setting. mongod enables journaling by default.
	//
	//The storage.journal.enabled setting is available only for mongod.
	//
	//Not available for mongod instances that use the in-memory storage engine.
	//
	//Starting in MongoDB 4.0, you cannot specify --nojournal option or storage.journal.enabled: false for replica set members that use the WiredTiger storage engine.
	Enabled LogicBoolean `yaml:"-"`

	//Default: 100
	//
	//The maximum amount of time in milliseconds that the mongod process allows between journal operations. Values can range from 1 to 500 milliseconds. Lower values increase the durability of the journal, at the expense of disk performance.
	//
	//On WiredTiger, the default journal commit interval is 100 milliseconds. Additionally, a write that includes or implies j:true will cause an immediate sync of the journal. For details or additional conditions that affect the frequency of the sync, see Journaling Process.
	//
	//The storage.journal.commitIntervalMs setting is available only for mongod.
	//
	//Not available for mongod instances that use the in-memory storage engine.
	//
	//NOTE: Known Issue in 4.2.0: The storage.journal.commitIntervalMs is missing in 4.2.0.
	CommitIntervalMs int `yaml:"-"`
}

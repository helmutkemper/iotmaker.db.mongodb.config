package iotmaker_db_mongodb_config

type OperationProfiling struct {
	//Default: off
	//
	//Specifies which operations should be profiled. The following profiler levels are available:
	//
	//Level	Description
	//off		The profiler is off and does not collect any data. This is the default profiler level.
	//slowOp	The profiler collects data for operations that take longer than the value of slowms.
	//all		The profiler collects data for all operations.
	//IMPORTANT: Profiling can impact performance and shares settings with the system log. Carefully consider any performance and security implications before configuring and enabling the profiler on a production deployment.
	//
	//See Profiler Overhead for more information on potential performance degradation.
	Mode string `yaml:"mode"`

	//Default: 100
	//
	//The slow operation time threshold, in milliseconds. Operations that run for longer than this threshold are considered slow.
	//
	//When logLevel is set to 0, MongoDB records slow operations to the diagnostic log at a rate determined by slowOpSampleRate. Starting in MongoDB 4.2, the secondaries of replica sets log all oplog entry messages that take longer than the slow operation threshold to apply regardless of the sample rate.
	//
	//At higher logLevel settings, all operations appear in the diagnostic log regardless of their latency with the following exception: the logging of slow oplog entry messages by the secondaries. The secondaries log only the slow oplog entries; increasing the logLevel does not log all oplog entries.
	//
	//Changed in version 4.0: The slowOpThresholdMs setting is available for mongod and mongos. In earlier versions, slowOpThresholdMs is available for mongod only.
	//
	//For mongod instances, the setting affects both the diagnostic log and, if enabled, the profiler.
	//For mongos instances, the setting affects the diagnostic log only and not the profiler since profiling is not available on mongos.
	SlowOpThresholdMs int `yaml:"slowOpThresholdMs"`

	//Default: 1.0
	//
	//The fraction of slow operations that should be profiled or logged. operationProfiling.slowOpSampleRate accepts values between 0 and 1, inclusive.
	//
	//operationProfiling.slowOpSampleRate does not affect the slow oplog entry logging by the secondary members of a replica set. Secondary members log all oplog entries that take longer than the slow operation threshold regardless of the operationProfiling.slowOpSampleRate.
	//
	//Changed in version 4.0: The slowOpSampleRate setting is available for mongod and mongos. In earlier versions, slowOpSampleRate is available for mongod only.
	//
	//For mongod instances, the setting affects both the diagnostic log and, if enabled, the profiler.
	//For mongos instances, the setting affects the diagnostic log only and not the profiler since profiling is not available on mongos.
	SlowOpSampleRate float64 `yaml:"slowOpSampleRate"`
}

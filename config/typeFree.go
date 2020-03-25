package iotmaker_db_mongodb_config

type Free struct {
	//New in version 4.0: Available for MongoDB Community Edition.
	//
	//Enables or disables free MongoDB Cloud monitoring. cloud.monitoring.free.state
	//accepts the following values:
	//
	//runtime
	//Default. You can enable or disable free monitoring during runtime.
	//
	//To enable or disable free monitoring during runtime, see db.enableFreeMonitoring()
	//and db.disableFreeMonitoring().
	//
	//To enable or disable free monitoring during runtime when running with access control,
	//users must have required privileges. See db.enableFreeMonitoring() and
	//db.disableFreeMonitoring() for details.
	//
	//    Value  |  Description
	//    ---------------------------------------------------------------------------------
	//    on     |  Enables free monitoring at startup;
	//           |  i.e. registers for free monitoring.
	//           |  When enabled at startup, you cannot disable free monitoring during
	//           |  runtime.
	//    ---------------------------------------------------------------------------------
	//    off    |  Disables free monitoring at startup, regardless of whether you have
	//           |  previously registered for free monitoring.
	//           |  When disabled at startup, you cannot enable free monitoring during
	//           |  runtime.
	//
	//Once enabled, the free monitoring state remains enabled until explicitly disabled.
	//That is, you do not need to re-enable each time you start the server.
	//
	//For the corresponding command-line option, see --enableFreeMonitoring.
	State FreeState `yaml:"state"`

	//New in version 4.0: Available for MongoDB Community Edition.
	//
	//Optional tag to describe environment context. The tag can be sent as part of the free
	//MongoDB Cloud monitoring registration at start up.
	//
	//For the corresponding command-line option, see --freeMonitoringTag.
	Tags string `yaml:"tags"`
}

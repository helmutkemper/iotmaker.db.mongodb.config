package iotmaker_db_mongodb_config

type Snmp struct {
	//Default: false
	//
	//Disables SNMP access to mongod. The option is incompatible with snmp.subagent and snmp.master.
	//
	//Set to true to disable SNMP access.
	//
	//The snmp.disabled setting is available only for mongod.
	//
	//New in version 4.0.6.
	Disabled LogicBoolean `yaml:"disabled"`

	//When snmp.subagent is true, SNMP runs as a subagent. The option is incompatible with snmp.disabled set to true.
	//
	//The snmp.subagent setting is available only for mongod.
	SubAgent LogicBoolean `yaml:"subagent"`

	//When snmp.master is true, SNMP runs as a master. The option is incompatible with snmp.disabled set to true.
	//
	//The snmp.master setting is available only for mongod.
	master LogicBoolean `yaml:"master"`
}

package iotmaker_db_mongodb_config

import (
	"bytes"
	"errors"
	"fmt"
	"reflect"
	"strconv"
)

//https://docs.mongodb.com/manual/reference/configuration-options/#cloud.monitoring.free.state
//todo:https://docs.mongodb.com/manual/reference/configuration-options/#mongos-only-options

type AccessControlLog struct {
	//Default: 0
	//
	//The log message verbosity level for components related to access control. See ACCESS components.
	//
	//The verbosity level can range from 0 to 5:
	//    0 is the MongoDB’s default log verbosity level, to include Informational messages.
	//    1 to 5 increases the verbosity level to include Debug messages.
	Verbosity Verbosity `yaml:"verbosity"`
}

type CommandLog struct {
	//Default: 0
	//
	//The log message verbosity level for components related to commands. See COMMAND components.
	//
	//The verbosity level can range from 0 to 5:
	//    0 is the MongoDB’s default log verbosity level, to include Informational messages.
	//    1 to 5 increases the verbosity level to include Debug messages.
	Verbosity Verbosity `yaml:"verbosity"`
}

type ControlLog struct {
	//Default: 0
	//
	//The log message verbosity level for components related to control operations. See CONTROL components.
	//
	//The verbosity level can range from 0 to 5:
	//    0 is the MongoDB’s default log verbosity level, to include Informational messages.
	//    1 to 5 increases the verbosity level to include Debug messages.
	Verbosity Verbosity `yaml:"verbosity"`
}

type FtdcLog struct {
	//Default: 0
	//
	// > New in version 3.2.
	//
	//The log message verbosity level for components related to diagnostic data collection operations. See FTDC components.
	//
	//The verbosity level can range from 0 to 5:
	//    0 is the MongoDB’s default log verbosity level, to include Informational messages.
	//    1 to 5 increases the verbosity level to include Debug messages.
	Verbosity Verbosity `yaml:"verbosity"`
}

type GeoLog struct {
	//Default: 0
	//
	//The log message verbosity level for components related to geospatial parsing operations. See GEO components.
	//
	//The verbosity level can range from 0 to 5:
	//    0 is the MongoDB’s default log verbosity level, to include Informational messages.
	//    1 to 5 increases the verbosity level to include Debug messages.
	Verbosity Verbosity `yaml:"verbosity"`
}

type IndexLog struct {
	//Default: 0
	//
	//The log message verbosity level for components related to indexing operations. See INDEX components.
	//
	//The verbosity level can range from 0 to 5:
	//    0 is the MongoDB’s default log verbosity level, to include Informational messages.
	//    1 to 5 increases the verbosity level to include Debug messages.
	Verbosity Verbosity `yaml:"verbosity"`
}

type NetworkLog struct {
	//Default: 0
	//
	//The log message verbosity level for components related to networking operations. See NETWORK components.
	//
	//The verbosity level can range from 0 to 5:
	//    0 is the MongoDB’s default log verbosity level, to include Informational messages.
	//    1 to 5 increases the verbosity level to include Debug messages.
	Verbosity Verbosity `yaml:"verbosity"`
}

type QueryLog struct {
	// Default: 0
	//
	//The log message verbosity level for components related to query operations. See QUERY components.
	//
	//The verbosity level can range from 0 to 5:
	//    0 is the MongoDB’s default log verbosity level, to include Informational messages.
	//    1 to 5 increases the verbosity level to include Debug messages.
	Verbosity Verbosity `yaml:"verbosity"`
}

type ElectionLog struct {
	//Default: 0
	//
	// > New in version 4.2.
	//
	//The log message verbosity level for components related to election. See ELECTION components.
	//
	//If systemLog.component.replication.election.verbosity is unset, systemLog.component.replication.verbosity level also applies to election components.
	//
	//The verbosity level can range from 0 to 5:
	//    0 is the MongoDB’s default log verbosity level, to include Informational messages.
	//    1 to 5 increases the verbosity level to include Debug messages.
	Verbosity Verbosity `yaml:"verbosity"`
}

type ReplicationLog struct {
	//Default: 0
	//
	//The log message verbosity level for components related to replication. See REPL components.
	//
	//The verbosity level can range from 0 to 5:
	//    0 is the MongoDB’s default log verbosity level, to include Informational messages.
	//    1 to 5 increases the verbosity level to include Debug messages.
	Verbosity Verbosity `yaml:"verbosity"` //ok

	Election    ElectionLog    `yaml:"Election"`    //ok
	Heartbeats  HeartbeatsLog  `yaml:"heartbeats"`  //ok
	InitialSync InitialSyncLog `yaml:"initialSync"` //ok
	Rollback    RollbackLog    `yaml:"rollback"`    //ok
	Sharding    ShardingLog    `yaml:"sharding"`    //ok
	Storage     StorageLog     `yaml:"storage"`     //ok
}

type StorageLog struct {
	//Default: 0
	//
	//The log message verbosity level for components related to storage. See STORAGE components.
	//
	//If systemLog.component.storage.journal.verbosity is unset, systemLog.component.storage.verbosity level also applies to journaling components.
	//
	//The verbosity level can range from 0 to 5:
	//    0 is the MongoDB’s default log verbosity level, to include Informational messages.
	//    1 to 5 increases the verbosity level to include Debug messages.
	Verbosity Verbosity `yaml:"verbosity"`

	Journal  JournalLog  `yaml:"journal"`
	Recovery RecoveryLog `yaml:"recovery"`
}

type RecoveryLog struct {
	//Default: 0
	//
	// > New in version 4.0.
	//
	//The log message verbosity level for components related to recovery. See RECOVERY components.
	//
	//If systemLog.component.storage.recovery.verbosity is unset, systemLog.component.storage.verbosity level also applies to recovery components.
	//
	//The verbosity level can range from 0 to 5:
	//    0 is the MongoDB’s default log verbosity level, to include Informational messages.
	//    1 to 5 increases the verbosity level to include Debug messages.
	Verbosity Verbosity `yaml:"verbosity"`
}

type JournalLog struct {
	//Default: 0
	//
	//The log message verbosity level for components related to journaling. See JOURNAL components.
	//
	//If systemLog.component.storage.journal.verbosity is unset, the journaling components have the same verbosity level as the parent storage components: i.e. either the systemLog.component.storage.verbosity level if set or the default verbosity level.
	//
	//The verbosity level can range from 0 to 5:
	//    0 is the MongoDB’s default log verbosity level, to include Informational messages.
	//    1 to 5 increases the verbosity level to include Debug messages.
	Verbosity Verbosity `yaml:"verbosity"`
}

type ShardingLog struct {
	//Default: 0
	//
	//The log message verbosity level for components related to sharding. See SHARDING components.
	//
	//The verbosity level can range from 0 to 5:
	//    0 is the MongoDB’s default log verbosity level, to include Informational messages.
	//    1 to 5 increases the verbosity level to include Debug messages.
	Verbosity Verbosity `yaml:"verbosity"`
}

type RollbackLog struct {
	//Default: 0
	//
	// > New in version 3.6.
	//
	//The log message verbosity level for components related to rollback. See ROLLBACK components.
	//
	//If systemLog.component.replication.rollback.verbosity is unset, systemLog.component.replication.verbosity level also applies to rollback components.
	//
	//The verbosity level can range from 0 to 5:
	//    0 is the MongoDB’s default log verbosity level, to include Informational messages.
	//    1 to 5 increases the verbosity level to include Debug messages.
	Verbosity Verbosity `yaml:"verbosity"`
}

type HeartbeatsLog struct {
	//Default: 0
	//
	// > New in version 3.6.
	//
	//The log message verbosity level for components related to heartbeats. See REPL_HB components.
	//
	//If systemLog.component.replication.heartbeats.verbosity is unset, systemLog.component.replication.verbosity level also applies to heartbeats components.
	//
	//The verbosity level can range from 0 to 5:
	//    0 is the MongoDB’s default log verbosity level, to include Informational messages.
	//    1 to 5 increases the verbosity level to include Debug messages.
	Verbosity Verbosity `yaml:"verbosity"`
}

type InitialSyncLog struct {
	//Default: 0
	//
	// > New in version 4.2.
	//
	//The log message verbosity level for components related to initialSync. See INITSYNC components.
	//
	//If systemLog.component.replication.initialSync.verbosity is unset, systemLog.component.replication.verbosity level also applies to initialSync components.
	//
	//The verbosity level can range from 0 to 5:
	//    0 is the MongoDB’s default log verbosity level, to include Informational messages.
	//    1 to 5 increases the verbosity level to include Debug messages.
	Verbosity Verbosity `yaml:"verbosity"`
}

type WriteLog struct {
	//Default: 0
	//
	//The log message verbosity level for components related to write operations. See WRITE components.
	//
	//The verbosity level can range from 0 to 5:
	//    0 is the MongoDB’s default log verbosity level, to include Informational messages.
	//    1 to 5 increases the verbosity level to include Debug messages.
	Verbosity Verbosity `yaml:"verbosity"`
}

type TransactionLog struct {
	//Default: 0
	//
	// > New in version 4.0.2.
	//
	//The log message verbosity level for components related to transaction. See TXN components.
	//
	//The verbosity level can range from 0 to 5:
	//    0 is the MongoDB’s default log verbosity level, to include Informational messages.
	//    1 to 5 increases the verbosity level to include Debug messages.
	Verbosity Verbosity `yaml:"verbosity"`
}

type Component struct {
	AccessControl AccessControlLog `yaml:"accessControl"` //ok
	Command       CommandLog       `yaml:"command"`       //ok
	Control       ControlLog       `yaml:"control"`       //ok
	Ftdc          FtdcLog          `yaml:"ftdc"`          //ok
	Geo           GeoLog           `yaml:"geo"`           //ok
	Index         IndexLog         `yaml:"index"`         //ok
	Network       NetworkLog       `yaml:"Network"`       //ok
	Query         QueryLog         `yaml:"query"`         //ok
	Replication   ReplicationLog   `yaml:"replication"`   //ok
	Transaction   TransactionLog   `yaml:"transaction"`   //ok
	Write         WriteLog         `yaml:"write"`         //ok
}

type ProcessManagement struct {
	//Default: false
	//
	//Enable a daemon mode that runs the mongos or mongod process in the background. By default mongos or mongod does not run as a daemon: typically you will run mongos or mongod as a daemon, either by using processManagement.fork or by using a controlling process that handles the daemonization process (e.g. as with upstart and systemd).
	//
	//The processManagement.fork option is not supported on Windows.
	//
	//The Linux package init scripts do not expect processManagement.fork to change from the defaults. If you use the Linux packages and change processManagement.fork, you will have to use your own init scripts and disable the built-in scripts.
	Fork LogicBoolean `yaml:"fork"`

	//Specifies a file location to store the process ID (PID) of the mongos or mongod process . The user running the the mongod or mongos process must be able to write to this path. If the processManagement.pidFilePath option is not specified, the process does not create a PID file. This option is generally only useful in combination with the processManagement.fork setting.
	//    LINUX: On Linux, PID file management is generally the responsibility of your distro’s init system:
	//      usually a service file in the /etc/init.d directory, or a systemd unit file registered with systemctl. Only use the processManagement.pidFilePath option if you are not using one of these init systems. For more information, please see the respective Installation Guide for your operating system.
	//    MACOS: On macOS, PID file management is generally handled by brew. Only use the processManagement.pidFilePath option if you are not using brew on your macOS system. For more information, please see the respective Installation Guide for your operating system.
	PidFilePath string `yaml:"pidFilePath"`

	//The full path from which to load the time zone database. If this option is not provided, then MongoDB will use its built-in time zone database.
	//
	//The configuration file included with Linux and macOS packages sets the time zone database path to /usr/share/zoneinfo by default.
	//
	//The built-in time zone database is a copy of the Olson/IANA time zone database. It is updated along with MongoDB releases, but the release cycle of the time zone database differs from the release cycle of MongoDB. A copy of the most recent release of the time zone database can be downloaded from https://downloads.mongodb.org/olson_tz_db/timezonedb-latest.zip.
	TimeZoneInfo string `yaml:"timeZoneInfo"`
}

type Configuration struct {
	//Default: 0
	//
	//The default log message verbosity level for components. The verbosity level determines the amount of Informational and Debug messages MongoDB outputs. [2]
	//
	//The verbosity level can range from 0 to 5:
	//    0 is the MongoDB’s default log verbosity level, to include Informational messages.
	//    1 to 5 increases the verbosity level to include Debug messages.
	//To use a different verbosity level for a named component, use the component’s verbosity setting. For example, use the systemLog.component.accessControl.verbosity to set the verbosity level specifically for ACCESS components.
	//
	//See the systemLog.component.<name>.verbosity settings for specific component verbosity settings.
	//
	//For various ways to set the log verbosity level, see Configure Log Verbosity Levels.
	//
	//[2]	Starting in version 4.2, MongoDB includes the Debug verbosity level (1-5) in the log messages. For example, if the verbosity level is 2, MongoDB logs D2. In previous versions, MongoDB log messages only specified D for Debug level.
	Verbosity Verbosity `yaml:"verbosity"` //ok

	//Run mongos or mongod in a quiet mode that attempts to limit the amount of output.
	//
	//systemLog.quiet is not recommended for production systems as it may make tracking problems during particular connections much more difficult.
	Quiet LogicBoolean `yaml:"quiet"` //ok

	//Print verbose information for debugging. Use for additional logging for support-related troubleshooting.
	TraceAllExceptions LogicBoolean `yaml:"traceAllExceptions"` //ok

	//Default: user
	//
	//The facility level used when logging messages to syslog. The value you specify must be supported by your operating system’s implementation of syslog. To use this option, you must set systemLog.destination to syslog.
	SyslogFacility int `yaml:"syslogFacility"` //ok

	//The path of the log file to which mongod or mongos should send all diagnostic logging information, rather than the standard output or the host’s syslog. MongoDB creates the log file at the specified path.
	//
	//The Linux package init scripts do not expect systemLog.path to change from the defaults. If you use the Linux packages and change systemLog.path, you will have to use your own init scripts and disable the built-in scripts.
	Path string `yaml:"path"` //ok

	//Default: false
	//
	//When true, mongos or mongod appends new entries to the end of the existing log file when the mongos or mongod instance restarts. Without this option, mongod will back up the existing log and create a new file.
	LogAppend LogicBoolean `yaml:"logAppend"` //ok

	//Default: rename
	//
	//The behavior for the logRotate command. Specify either rename or reopen:
	//
	//rename renames the log file.
	//
	//reopen closes and reopens the log file following the typical Linux/Unix log rotate behavior. Use reopen when using the Linux/Unix logrotate utility to avoid log loss.
	//
	//If you specify reopen, you must also set systemLog.logAppend to true.
	LogRotate LogRotate `yaml:"logRotate"` //ok

	//The destination to which MongoDB sends all log output. Specify either file or syslog. If you specify file, you must also specify systemLog.path.
	//
	//If you do not specify systemLog.destination, MongoDB sends all log output to standard output.
	//
	//WARNING: The syslog daemon generates timestamps when it logs a message, not when MongoDB issues the message. This can lead to misleading timestamps for log entries, especially when the system is under heavy load. We recommend using the file option for production systems to ensure accurate timestamps.
	Destination string `yaml:"destination"` //ok

	//Default: iso8601-local
	//
	//The time format for timestamps in log messages. Specify one of the following values:
	//
	//Value	Description
	//ctime		Displays timestamps as Wed Dec 31 18:17:54.811.
	//iso8601-utc	Displays timestamps in Coordinated Universal Time (UTC) in the ISO-8601 format. For example, for New York at the start of the Epoch: 1970-01-01T00:00:00.000Z
	//iso8601-local	Displays timestamps in local time in the ISO-8601 format. For example, for New York at the start of the Epoch: 1969-12-31T19:00:00.000-0500
	TimeStampFormat TimeStampFormat `yaml:"timeStampFormat"` //ok

	//NOTE: Starting in version 4.2, MongoDB includes the Debug verbosity level (1-5) in the log messages. For example, if the verbosity level is 2, MongoDB logs D2. In previous versions, MongoDB log messages only specified D for Debug level.
	Component         Component         `yaml:"component"`         //ok
	ProcessManagement ProcessManagement `yaml:"processManagement"` //ok
	Cloud             Cloud             `yaml:"cloud"`             //ok

	//Changed in version 4.2: MongoDB 4.2 deprecates ssl options in favor of tls options with identical functionality.
	Net Net `yaml:"net"` //ok

	Security Security `yaml:"security"`

	//Set MongoDB parameter or parameters described in MongoDB Server Parameters
	//
	//To set parameters in the YAML configuration file, use the following format:
	//
	//setParameter:
	//   <parameter1>: <value1>
	//   <parameter2>: <value2>
	//For example, to specify the enableLocalhostAuthBypass in the configuration file:
	//
	//setParameter:
	//   enableLocalhostAuthBypass: false
	SetParameter interface{} `yaml:"setParameter"` //todo: fazer

	//Default: 30
	//
	//For use with mongod servers using LDAP Authorization.
	//
	//The interval (in seconds) mongod waits between external user cache flushes. After mongod flushes the external user cache, MongoDB reacquires authorization data from the LDAP server the next time an LDAP-authorized user issues an operation.
	//
	//Increasing the value specified increases the amount of time mongod and the LDAP server can be out of sync, but reduces the load on the LDAP server. Conversely, decreasing the value specified decreases the time mongod and the LDAP server can be out of sync while increasing the load on the LDAP server.
	//
	//setParameter:
	//   ldapUserCacheInvalidationInterval: <int>
	LdapUserCacheInvalidationInterval int `yaml:"ldapUserCacheInvalidationInterval"`

	Storage            Storage            `yaml:"storage"`
	OperationProfiling OperationProfiling `yaml:"operationProfiling"`
	Replication        Replication        `yaml:"replication"`
	Sharding           Sharding           `yaml:"sharding"`
	AuditLog           AuditLog           `yaml:"auditLog"`
	Snmp               Snmp               `yaml:"snmp"`
}

func (el *Configuration) getTagData(tag reflect.StructTag) (string, string) {
	var tagName, tagValue string

	tagName = "htmlAttr"
	tagValue = tag.Get(tagName)
	if tagValue != "" {
		return tagName, tagValue
	}

	tagName = "htmlAttrSet"
	tagValue = tag.Get(tagName)
	if tagValue != "" {
		return tagName, tagValue
	}

	tagName = "htmlAttrOnOff"
	tagValue = tag.Get(tagName)
	if tagValue != "" {
		return tagName, tagValue
	}

	tagName = "htmlAttrTrueFalse"
	tagValue = tag.Get(tagName)
	if tagValue != "" {
		return tagName, tagValue
	}

	return "", ""
}

func (el *Configuration) getTagDataByName(tagName string, tag reflect.StructTag) (string, string) {
	var tagValue string

	tagValue = tag.Get(tagName)
	if tagValue != "" {
		return tagName, tagValue
	}

	return "", ""
}

func (el *Configuration) writeSpaces(buffer *bytes.Buffer, spaces int) {
	for i := 0; i != spaces; i += 1 {
		buffer.WriteString(" ")
	}
}

func (el *Configuration) writeKeyValue(buffer *bytes.Buffer, key, value string) {
	buffer.WriteString(key)
	buffer.WriteString(": ")
	buffer.WriteString(value)
	buffer.WriteString("\r\n")
}

func (el *Configuration) ToYaml(level int, element reflect.Value) (error, []byte) {
	var buffer bytes.Buffer
	var err = el.ToText(level, element, &buffer)
	return err, buffer.Bytes()
}

func (el *Configuration) ToText(level int, element reflect.Value, buffer *bytes.Buffer) error {
	var tagName, tagValue string

	for i := 0; i < element.NumField(); i += 1 {
		field := element.Field(i)
		typeField := element.Type().Field(i)
		tag := typeField.Tag

		tagName, tagValue = el.getTagData(tag)

		if tagValue == "-" {
			continue
		}

		switch field.Type().String() {
		case "Verbosity":

			if field.Interface().(Verbosity) == 0 {
				continue
			}

			if field.Interface().(Verbosity) < 0 {
				return errors.New("verbosity must be between 0 and 5")
			}

			if field.Interface().(Verbosity) > 5 {
				return errors.New("verbosity must be between 0 and 5")
			}

			el.writeSpaces(buffer, level)
			str := strconv.FormatInt(int64(field.Interface().(Verbosity)), 10)
			el.writeKeyValue(buffer, tagValue, str)

		case "LogicBoolean":

			if field.Interface().(LogicBoolean) == -1 {
				el.writeSpaces(buffer, level)
				el.writeKeyValue(buffer, tagValue, "false")
				continue
			}

			if field.Interface().(LogicBoolean) == 0 {
				continue
			}

			if field.Interface().(LogicBoolean) == 1 {
				el.writeSpaces(buffer, level)
				el.writeKeyValue(buffer, tagValue, "true")
				continue
			}

			return errors.New("logical must be -1 for false, 0 for not set and 1 for true")

		case "string":
			if field.Interface().(string) == "" {
				continue
			}

			el.writeSpaces(buffer, level)
			el.writeKeyValue(buffer, tagValue, field.Interface().(string))

		case "int":
			if field.Interface().(int) == 0 {
				continue
			}

			el.writeSpaces(buffer, level)
			str := strconv.FormatInt(int64(field.Interface().(int)), 10)
			el.writeKeyValue(buffer, tagValue, str)

		case "int64":
			if field.Interface().(int) == 0 {
				continue
			}

			el.writeSpaces(buffer, level)
			str := strconv.FormatInt(field.Interface().(int64), 10)
			el.writeKeyValue(buffer, tagValue, str)

		case "TimeStampFormat":
			if field.Interface().(TimeStampFormat).String() == "" {
				continue
			}

			el.writeSpaces(buffer, level)
			str := field.Interface().(TimeStampFormat).String()
			el.writeKeyValue(buffer, tagValue, str)

		case "LogRotate":
			if field.Interface().(LogRotate).String() == "" {
				continue
			}

			el.writeSpaces(buffer, level)
			str := field.Interface().(LogRotate).String()
			el.writeKeyValue(buffer, tagValue, str)

		case "Component":
			fallthrough
		case "Cloud":
			fallthrough
		case "ProcessManagement":
			fallthrough
		case "AccessControlLog":
			fallthrough
		case "CommandLog":
			fallthrough
		case "ControlLog":
			fallthrough
		case "FtdcLog":
			fallthrough
		case "GeoLog":
			fallthrough
		case "IndexLog":
			fallthrough
		case "NetworkLog":
			fallthrough
		case "QueryLog":
			fallthrough
		case "ReplicationLog":
			fallthrough
		case "TransactionLog":
			fallthrough
		case "WriteLog":
			fallthrough
		case "ElectionLog":
			fallthrough
		case "HeartbeatsLog":
			fallthrough
		case "InitialSyncLog":
			fallthrough
		case "RollbackLog":
			fallthrough
		case "ShardingLog":
			fallthrough
		case "StorageLog":
			fallthrough
		case "Net":
			fallthrough
		case "Security":

			err := el.ToText(level+2, field, buffer)
			if err != nil {
				return err
			}

		default:
			fmt.Printf("\nhtmlTag(): %v[ %d ]: %s - %s = %v\n", tagValue, i, field.Type(), field.Interface(), tagName)
		}
	}

	return nil
}

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
	Master LogicBoolean `yaml:"master"`
}

type AuditLog struct {
	//When set, auditLog.destination enables auditing and specifies where mongos or mongod sends all audit events.
	//
	//auditLog.destination can have one of the following values:
	//
	//Value		Description
	//syslog 	Output the audit events to syslog in JSON format. Not available on Windows. Audit messages have a syslog severity level of info and a facility level of user.
	//			The syslog message limit can result in the truncation of audit messages. The auditing system will neither detect the truncation nor error upon its occurrence.
	//
	//console	Output the audit events to stdout in JSON format.
	//file		Output the audit events to the file specified in auditLog.path in the format specified in auditLog.format.
	//
	//NOTE: Available only in MongoDB Enterprise and MongoDB Atlas.
	Destination string `yaml:"destination"`

	//The format of the output file for auditing if destination is file. The auditLog.format option can have one of the following values:
	//
	//Value	Description
	//JSON	Output the audit events in JSON format to the file specified in auditLog.path.
	//BSON	Output the audit events in BSON binary format to the file specified in auditLog.path.
	//Printing audit events to a file in JSON format degrades server performance more than printing to a file in BSON format.
	//
	//NOTE: Available only in MongoDB Enterprise and MongoDB Atlas.
	Format string `yaml:"format"`

	//The output file for auditing if destination has value of file. The auditLog.path option can take either a full path name or a relative path name.
	//
	//NOTE: Available only in MongoDB Enterprise and MongoDB Atlas.
	Path string `yaml:"path"`

	//Type: string representation of a document
	//
	//The filter to limit the types of operations the audit system records. The option takes a string representation of a query document of the form:
	//
	//{ <field1>: <expression1>, ... }
	//The <field> can be any field in the audit message, including fields returned in the param document. The <expression> is a query condition expression.
	//
	//To specify an audit filter, enclose the filter document in single quotes to pass the document as a string.
	//
	//To specify the audit filter in a configuration file, you must use the YAML format of the configuration file.
	//
	//NOTE: Available only in MongoDB Enterprise and MongoDB Atlas.
	Filter string `yaml:"filter"`
}

type Sharding struct {
	//The role that the mongod instance has in the sharded cluster. Set this setting to one of the following:
	//
	//Value			Description
	//configsvr		Start this instance as a config server. The instance starts on port 27019 by default.
	//shardsvr		Start this instance as a shard. The instance starts on port 27018 by default.
	//NOTE
	//
	//Setting sharding.clusterRole requires the mongod instance to be running with replication. To deploy the instance as a replica set member, use the replSetName setting and specify the name of the replica set.
	//
	//The sharding.clusterRole setting is available only for mongod.
	ClusterRole string `yaml:"clusterRole"`

	//Changed in version 3.2: Starting in 3.2, MongoDB uses false as the default.
	//
	//During chunk migration, a shard does not save documents migrated from the shard.
	archiveMovedChunks LogicBoolean `yaml:"archiveMovedChunks"`
}

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
	OpLogSizeMB int64 `yaml:"oplogSizeMB"`

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
	ReplSetName string `yaml:"replSetName"`

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
	EnableMajorityReadConcern LogicBoolean `yaml:"enableMajorityReadConcern"`
}

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
	Inmemory   Inmemory   `yaml:"inMemory"`
}

type Inmemory struct {
	EngineConfig InmemoryEngineConfig `yaml:"engineConfig"`
}

type InmemoryEngineConfig struct {
	//Default: 50% of physical RAM less 1 GB
	//
	//Changed in version 3.4: Values can range from 256MB to 10TB and can be a float.
	//
	//Maximum amount of memory to allocate for in-memory storage engine data, including indexes, oplog if the mongod is part of replica set, replica set or sharded cluster metadata, etc.
	//
	//By default, the in-memory storage engine uses 50% of physical RAM minus 1 GB.
	//
	//ENTERPRISE FEATURE: Available in MongoDB Enterprise only.
	InMemorySizeGB float64 `yaml:"inMemorySizeGB"`
}

type WiredTiger struct {
	EngineConfig WiredTigerEngineConfig `yaml:"engineConfig"`
	IndexConfig  IndexConfig            `yaml:"indexConfig"`
}

type IndexConfig struct {
	//Default: true
	//
	//Enables or disables prefix compression for index data.
	//
	//Specify true for storage.wiredTiger.indexConfig.prefixCompression to enable prefix compression for index data, or false to disable prefix compression for index data.
	//
	//The storage.wiredTiger.indexConfig.prefixCompression setting affects all indexes created. If you change the value of storage.wiredTiger.indexConfig.prefixCompression on an existing MongoDB deployment, all new indexes will use prefix compression. Existing indexes are not affected.
	PrefixCompression LogicBoolean `yaml:"prefixCompression"`
}

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

type CollectionConfig struct {
	//Default: snappy
	//
	//Specifies the default compression for collection data. You can override this on a per-collection basis when creating collections.
	//
	//Available compressors are:
	//
	//none
	//snappy
	//zlib
	//zstd (Available starting MongoDB 4.2)
	//storage.wiredTiger.collectionConfig.blockCompressor affects all collections created. If you change the value of storage.wiredTiger.collectionConfig.blockCompressor on an existing MongoDB deployment, all new collections will use the specified compressor. Existing collections will continue to use the compressor specified when they were created, or the default compressor at that time.
	BlockCompressor string `yaml:"blockCompressor"`
}

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
	Enabled LogicBoolean `yaml:"enabled"`

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
	CommitIntervalMs int `yaml:"commitIntervalMs"`
}

type Security struct {
	//The path to a key file that stores the shared secret that MongoDB instances use to authenticate to each other in a sharded cluster or replica set. keyFile implies security.authorization. See Internal/Membership Authentication for more information.
	//
	//Starting in MongoDB 4.2, keyfiles for internal membership authentication use YAML format to allow for multiple keys in a keyfile. The YAML format accepts content of:
	//
	//a single key string (same as in earlier versions),
	//multiple key strings (each string must be enclosed in quotes), or
	//sequence of key strings.
	//The YAML format is compatible with the existing single-key keyfiles that use the text file format.
	KeyFile string `yaml:"keyFile"`

	//Default: keyFile
	//
	//The authentication mode used for cluster authentication. If you use internal x.509 authentication, specify so here. This option can have one of the following values:
	//
	//Value	Description
	//keyFile	Use a keyfile for authentication. Accept only keyfiles.
	//sendKeyFile	For rolling upgrade purposes. Send a keyfile for authentication but can accept both keyfiles and x.509 certificates.
	//sendX509	For rolling upgrade purposes. Send the x.509 certificate for authentication but can accept both keyfiles and x.509 certificates.
	//x509	Recommended. Send the x.509 certificate for authentication and accept only x.509 certificates.
	//If --tlsCAFile or tls.CAFile is not specified and you are not using x.509 authentication, the system-wide CA certificate store will be used when connecting to an TLS-enabled server.
	//
	//If using x.509 authentication, --tlsCAFile or tls.CAFile must be specified unless using --tlsCertificateSelector.
	//
	//For more information about TLS and MongoDB, see Configure mongod and mongos for TLS/SSL and TLS/SSL Configuration for Clients .
	ClusterAuthMode string `yaml:"clusterAuthMode"`

	//Default: disabled
	//
	//Enable or disable Role-Based Access Control (RBAC) to govern each user’s access to database resources and operations.
	//
	//Set this option to one of the following:
	//
	//Value	Description
	//enabled	A user can access only the database resources and actions for which they have been granted privileges.
	//disabled	A user can access any database and perform any action.
	//See Role-Based Access Control for more information.
	//
	//The security.authorization setting is available only for mongod.
	Authorization string `yaml:"authorization"`

	//Default: false
	//
	//New in version 3.4: Allows the mongod or mongos to accept and create authenticated and non-authenticated connections to and from other mongod and mongos instances in the deployment. Used for performing rolling transition of replica sets or sharded clusters from a no-auth configuration to internal authentication. Requires specifying a internal authentication mechanism such as security.keyFile.
	//
	//For example, if using keyfiles for internal authentication, the mongod or mongos creates an authenticated connection with any mongod or mongos in the deployment using a matching keyfile. If the security mechanisms do not match, the mongod or mongos utilizes a non-authenticated connection instead.
	//
	//A mongod or mongos running with security.transitionToAuth does not enforce user access controls. Users may connect to your deployment without any access control checks and perform read, write, and administrative operations.
	//
	//NOTE
	//
	//A mongod or mongos running with internal authentication and without security.transitionToAuth requires clients to connect using user access controls. Update clients to connect to the mongod or mongos using the appropriate user prior to restarting mongod or mongos without security.transitionToAuth.
	TransitionToAuth LogicBoolean `yaml:"transitionToAuth"`

	//Default: true
	//
	//Enables or disables the server-side JavaScript execution. When disabled, you cannot use operations that perform server-side execution of JavaScript code, such as the $where query operator, mapReduce command and the db.collection.mapReduce() method.
	JavascriptEnabled LogicBoolean `yaml:"javascriptEnabled"`

	//New in version 3.4: Available in MongoDB Enterprise only.
	//
	//A mongod or mongos running with security.redactClientLogData redacts any message accompanying a given log event before logging. This prevents the mongod or mongos from writing potentially sensitive data stored on the database to the diagnostic log. Metadata such as error or operation codes, line numbers, and source file names are still visible in the logs.
	//
	//Use security.redactClientLogData in conjunction with Encryption at Rest and TLS/SSL (Transport Encryption) to assist compliance with regulatory requirements.
	//
	//For example, a MongoDB deployment might store Personally Identifiable Information (PII) in one or more collections. The mongod or mongos logs events such as those related to CRUD operations, sharding metadata, etc. It is possible that the mongod or mongos may expose PII as a part of these logging operations. A mongod or mongos running with security.redactClientLogData removes any message accompanying these events before being output to the log, effectively removing the PII.
	//
	//Diagnostics on a mongod or mongos running with security.redactClientLogData may be more difficult due to the lack of data related to a log event. See the process logging manual page for an example of the effect of security.redactClientLogData on log output.
	//
	//On a running mongod or mongos, use setParameter with the redactClientLogData parameter to configure this setting.
	RedactClientLogData LogicBoolean `yaml:"redactClientLogData"`

	//New in version 3.6.
	//
	//A list of IP addresses/CIDR (Classless Inter-Domain Routing) ranges against which the mongod validates authentication requests from other members of the replica set and, if part of a sharded cluster, the mongos instances. The mongod verifies that the originating IP is either explicitly in the list or belongs to a CIDR range in the list. If the IP address is not present, the server does not authenticate the mongod or mongos.
	//
	//security.clusterIpSourceWhitelist has no effect on a mongod started without authentication.
	//
	//security.clusterIpSourceWhitelist requires specifying each IPv4/6 address or Classless Inter-Domain Routing (CIDR) range as a YAML list:
	//
	//security:
	//  clusterIpSourceWhitelist:
	//    - 192.0.2.0/24
	//    - 127.0.0.1
	//    - ::1
	//
	//IMPORTANT: Ensure security.clusterIpSourceWhitelist includes the IP address or CIDR ranges that include the IP address of each replica set member or mongos in the deployment to ensure healthy communication between cluster components.
	//
	//Key Management Configuration Options
	//security:
	//   enableEncryption: <boolean>
	//   encryptionCipherMode: <string>
	//   encryptionKeyFile: <string>
	//   kmip:
	//      keyIdentifier: <string>
	//      rotateMasterKey: <boolean>
	//      serverName: <string>
	//      port: <string>
	//      clientCertificateFile: <string>
	//      clientCertificatePassword: <string>
	//      clientCertificateSelector: <string>
	//      serverCAFile: <string>
	ClusterIpSourceWhitelist []string `yaml:"clusterIpSourceWhitelist"`

	//Default: false
	//
	//New in version 3.2: Enables encryption for the WiredTiger storage engine. You must set to true to pass in encryption keys and configurations.
	//
	//ENTERPRISE FEATURE
	//
	//Available in MongoDB Enterprise only.
	EnableEncryption LogicBoolean `yaml:"enableEncryption"`

	//Default: AES256-CBC
	//
	//New in version 3.2.
	//
	//The cipher mode to use for encryption at rest:
	//
	//Mode	Description
	//AES256-CBC	256-bit Advanced Encryption Standard in Cipher Block Chaining Mode
	//AES256-GCM
	//256-bit Advanced Encryption Standard in Galois/Counter Mode
	//
	//Available only on Linux.
	//
	//Changed in version 4.0: MongoDB Enterprise on Windows no longer supports AES256-GCM. This cipher is now available only on Linux.
	//
	//ENTERPRISE FEATURE
	//
	//Available in MongoDB Enterprise only.
	EncryptionCipherMode string `yaml:"encryptionCipherMode"`

	//New in version 3.2.
	//
	//The path to the local keyfile when managing keys via process other than KMIP. Only set when managing keys via process other than KMIP. If data is already encrypted using KMIP, MongoDB will throw an error.
	//
	//Requires security.enableEncryption to be true.
	//
	//ENTERPRISE FEATURE
	//
	//Available in MongoDB Enterprise only.
	EncryptionKeyFile string `yaml:"encryptionKeyFile"`
	Kmip              Kmip   `yaml:"kmip"`
	Sasl              Sasl   `yaml:"sasl"`
	Ldap              Ldap   `yaml:"ldap"`
}

type Bind struct {
	//New in version 3.4: Available in MongoDB Enterprise only.
	//
	//The identity with which mongod or mongos binds as, when connecting to or performing queries on an LDAP server.
	//
	//Only required if any of the following are true:
	//
	//Using LDAP authorization.
	//Using an LDAP query for security.ldap.userToDNMapping.
	//The LDAP server disallows anonymous binds
	//You must use queryUser with queryPassword.
	//
	//If unset, mongod or mongos will not attempt to bind to the LDAP server.
	//
	//This setting can be configured on a running mongod or mongos using setParameter.
	//
	//NOTE: Windows MongoDB deployments can use bindWithOSDefaults instead of queryUser and queryPassword. You cannot specify both queryUser and bindWithOSDefaults at the same time.
	QueryUser string `yaml:"queryUser"`

	//New in version 3.4: Available in MongoDB Enterprise only.
	//
	//The password used to bind to an LDAP server when using queryUser. You must use queryPassword with queryUser.
	//
	//If unset, mongod or mongos will not attempt to bind to the LDAP server.
	//
	//This setting can be configured on a running mongod or mongos using setParameter.
	//
	//NOTE: Windows MongoDB deployments can use bindWithOSDefaults instead of queryPassword and queryPassword. You cannot specify both queryPassword and bindWithOSDefaults at the same time.
	QueryPassword string `yaml:"queryPassword"`

	//Default: false
	//
	//New in version 3.4: Available in MongoDB Enterprise for the Windows platform only.
	//
	//Allows mongod or mongos to authenticate, or bind, using your Windows login credentials when connecting to the LDAP server.
	//
	//Only required if:
	//
	//Using LDAP authorization.
	//Using an LDAP query for username transformation.
	//The LDAP server disallows anonymous binds
	//Use useOSDefaults to replace queryUser and queryPassword.
	UseOSDefaults LogicBoolean `yaml:"useOSDefaults"`

	//Default: simple
	//
	//New in version 3.4: Available in MongoDB Enterprise only.
	//
	//The method mongod or mongos uses to authenticate to an LDAP server. Use with queryUser and queryPassword to connect to the LDAP server.
	//
	//method supports the following values:
	//
	//simple - mongod or mongos uses simple authentication.
	//sasl - mongod or mongos uses SASL protocol for authentication
	//If you specify sasl, you can configure the available SASL mechanisms using security.ldap.bind.saslMechanisms. mongod or mongos defaults to using DIGEST-MD5 mechanism.
	Method string `yaml:"method"`

	//Default: DIGEST-MD5
	//
	//New in version 3.4: Available in MongoDB Enterprise only.
	//
	//A comma-separated list of SASL mechanisms mongod or mongos can use when authenticating to the LDAP server. The mongod or mongos and the LDAP server must agree on at least one mechanism. The mongod or mongos dynamically loads any SASL mechanism libraries installed on the host machine at runtime.
	//
	//Install and configure the appropriate libraries for the selected SASL mechanism(s) on both the mongod or mongos host and the remote LDAP server host. Your operating system may include certain SASL libraries by default. Defer to the documentation associated with each SASL mechanism for guidance on installation and configuration.
	//
	//If using the GSSAPI SASL mechanism for use with Kerberos Authentication, verify the following for the mongod or mongos host machine:
	//
	//Linux
	//The KRB5_CLIENT_KTNAME environment variable resolves to the name of the client Linux Keytab Files for the host machine. For more on Kerberos environment variables, please defer to the Kerberos documentation.
	//The client keytab includes a User Principal for the mongod or mongos to use when connecting to the LDAP server and execute LDAP queries.
	//Windows
	//If connecting to an Active Directory server, the Windows Kerberos configuration automatically generates a Ticket-Granting-Ticket when the user logs onto the system. Set useOSDefaults to true to allow mongod or mongos to use the generated credentials when connecting to the Active Directory server and execute queries.
	//Set method to sasl to use this option.
	//
	//NOTE
	//
	//For a complete list of SASL mechanisms see the IANA listing. Defer to the documentation for your LDAP or Active Directory service for identifying the SASL mechanisms compatible with the service.
	//
	//MongoDB is not a source of SASL mechanism libraries, nor is the MongoDB documentation a definitive source for installing or configuring any given SASL mechanism. For documentation and support, defer to the SASL mechanism library vendor or owner.
	//
	//For more information on SASL, defer to the following resources:
	//
	//For Linux, please see the Cyrus SASL documentation.
	//For Windows, please see the Windows SASL documentation.
	SaslMechanisms string `yaml:"saslMechanisms"`
}

type Ldap struct {
	//New in version 3.4: Available in MongoDB Enterprise only.
	//
	//The LDAP server against which the mongod or mongos authenticates users or determines what actions a user is authorized to perform on a given database. If the LDAP server specified has any replicated instances, you may specify the host and port of each replicated server in a comma-delimited list.
	//
	//If your LDAP infrastructure partitions the LDAP directory over multiple LDAP servers, specify one LDAP server or any of its replicated instances to security.ldap.servers. MongoDB supports following LDAP referrals as defined in RFC 4511 4.1.10. Do not use security.ldap.servers for listing every LDAP server in your infrastructure.
	//
	//This setting can be configured on a running mongod or mongos using setParameter.
	//
	//If unset, mongod or mongos cannot use LDAP authentication or authorization.
	Servers string `yaml:"servers"`
	Bind    Bind   `yaml:"bind"`

	//Default: tls
	//
	//New in version 3.4: Available in MongoDB Enterprise only.
	//
	//By default, mongod or mongos creates a TLS/SSL secured connection to the LDAP server.
	//
	//For Linux deployments, you must configure the appropriate TLS Options in /etc/openldap/ldap.conf file. Your operating system’s package manager creates this file as part of the MongoDB Enterprise installation, via the libldap dependency. See the documentation for TLS Options in the ldap.conf OpenLDAP documentation for more complete instructions.
	//
	//For Windows deployment, you must add the LDAP server CA certificates to the Windows certificate management tool. The exact name and functionality of the tool may vary depending on operating system version. Please see the documentation for your version of Windows for more information on certificate management.
	//
	//Set transportSecurity to none to disable TLS/SSL between mongod or mongos and the LDAP server.
	//
	//WARNING: Setting transportSecurity to none transmits plaintext information and possibly credentials between mongod or mongos and the LDAP server.
	TransportSecurity string `yaml:"transportSecurity"`

	//Default: 10000
	//
	//New in version 3.4: Available in MongoDB Enterprise only.
	//
	//The amount of time in milliseconds mongod or mongos should wait for an LDAP server to respond to a request.
	//
	//Increasing the value of timeoutMS may prevent connection failure between the MongoDB server and the LDAP server, if the source of the failure is a connection timeout. Decreasing the value of timeoutMS reduces the time MongoDB waits for a response from the LDAP server.
	//
	//This setting can be configured on a running mongod or mongos using setParameter.
	TimeoutMS int `yaml:"timeoutMS"`

	//New in version 3.4: Available in MongoDB Enterprise only.
	//
	//Maps the username provided to mongod or mongos for authentication to a LDAP Distinguished Name (DN). You may need to use userToDNMapping to transform a username into an LDAP DN in the following scenarios:
	//
	//Performing LDAP authentication with simple LDAP binding, where users authenticate to MongoDB with usernames that are not full LDAP DNs.
	//Using an LDAP authorization query template that requires a DN.
	//Transforming the usernames of clients authenticating to Mongo DB using different authentication mechanisms (e.g. x.509, kerberos) to a full LDAP DN for authorization.
	//userToDNMapping expects a quote-enclosed JSON-string representing an ordered array of documents. Each document contains a regular expression match and either a substitution or ldapQuery template used for transforming the incoming username.
	//
	//Each document in the array has the following form:
	//
	//{
	//  match: "<regex>"
	//  substitution: "<LDAP DN>" | ldapQuery: "<LDAP Query>"
	//}
	//Field	Description	Example
	//match	An ECMAScript-formatted regular expression (regex) to match against a provided username. Each parenthesis-enclosed section represents a regex capture group used by substitution or ldapQuery.	"(.+)ENGINEERING" "(.+)DBA"
	//substitution
	//An LDAP distinguished name (DN) formatting template that converts the authentication name matched by the match regex into a LDAP DN. Each curly bracket-enclosed numeric value is replaced by the corresponding regex capture group extracted from the authentication username via the match regex.
	//
	//The result of the substitution must be an RFC4514 escaped string.
	//
	//"cn={0},ou=engineering, dc=example,dc=com"
	//ldapQuery	A LDAP query formatting template that inserts the authentication name matched by the match regex into an LDAP query URI encoded respecting RFC4515 and RFC4516. Each curly bracket-enclosed numeric value is replaced by the corresponding regex capture group extracted from the authentication username via the match expression. mongod or mongos executes the query against the LDAP server to retrieve the LDAP DN for the authenticated user. mongod or mongos requires exactly one returned result for the transformation to be successful, or mongod or mongos skips this transformation.	"ou=engineering,dc=example, dc=com??one?(user={0})"
	//NOTE
	//
	//An explanation of RFC4514, RFC4515, RFC4516, or LDAP queries is out of scope for the MongoDB Documentation. Please review the RFC directly or use your preferred LDAP resource.
	//
	//For each document in the array, you must use either substitution or ldapQuery. You cannot specify both in the same document.
	//
	//When performing authentication or authorization, mongod or mongos steps through each document in the array in the given order, checking the authentication username against the match filter. If a match is found, mongod or mongos applies the transformation and uses the output for authenticating the user. mongod or mongos does not check the remaining documents in the array.
	//
	//If the given document does not match the provided authentication name, or the transformation described by the document fails, mongod or mongos continues through the list of documents to find additional matches. If no matches are found in any document, mongod or mongos returns an error.
	//
	//EXAMPLE
	//
	//The following shows two transformation documents. The first document matches against any string ending in @ENGINEERING, placing anything preceeding the suffix into a regex capture group. The second document matches against any string ending in @DBA, placing anything preceeding the suffix into a regex capture group.
	//
	//IMPORTANT
	//
	//You must pass the array to userToDNMapping as a string.
	//
	//"[
	//   {
	//      match: "(.+)@ENGINEERING.EXAMPLE.COM",
	//      substitution: "cn={0},ou=engineering,dc=example,dc=com"
	//   },
	//   {
	//      match: "(.+)@DBA.EXAMPLE.COM",
	//      ldapQuery: "ou=dba,dc=example,dc=com??one?(user={0})"
	//
	//   }
	//
	//]"
	//A user with username alice@ENGINEERING.EXAMPLE.COM matches the first document. The regex capture group {0} corresponds to the string alice. The resulting output is the DN "cn=alice,ou=engineering,dc=example,dc=com".
	//
	//A user with username bob@DBA.EXAMPLE.COM matches the second document. The regex capture group {0} corresponds to the string bob. The resulting output is the LDAP query "ou=dba,dc=example,dc=com??one?(user=bob)". mongod or mongos executes this query against the LDAP server, returning the result "cn=bob,ou=dba,dc=example,dc=com".
	//
	//If userToDNMapping is unset, mongod or mongos applies no transformations to the username when attempting to authenticate or authorize a user against the LDAP server.
	//
	//This setting can be configured on a running mongod or mongos using the setParameter database command.
	UserToDNMapping string `yaml:"userToDNMapping"`
	Authz           Authz  `yaml:"authz"`

	//Default: true
	//
	//Available in MongoDB Enterprise
	//
	//A flag that determines if the mongod or mongos instance checks the availability of the LDAP server(s) as part of its startup:
	//
	//If true, the mongod or mongos instance performs the availability check and only continues to start up if the LDAP server is available.
	//If false, the mongod or mongos instance skips the availability check; i.e. the instance starts up even if the LDAP server is unavailable.
	ValidateLDAPServerConfig LogicBoolean `yaml:"validateLDAPServerConfig"`
}

type Authz struct {
	//New in version 3.4: Available in MongoDB Enterprise only.
	//
	//A relative LDAP query URL formatted conforming to RFC4515 and RFC4516 that mongod executes to obtain the LDAP groups to which the authenticated user belongs to. The query is relative to the host or hosts specified in security.ldap.servers.
	//
	//In the URL, you can use the following substituion tokens:
	//
	//Substitution Token	Description
	//{USER}	Substitutes the authenticated username, or the transformed username if a userToDNMapping is specified.
	//{PROVIDED_USER}
	//Substitutes the supplied username, i.e. before either authentication or LDAP transformation.
	//
	//New in version 4.2.
	//
	//When constructing the query URL, ensure that the order of LDAP parameters respects RFC4516:
	//
	//[ dn  [ ? [attributes] [ ? [scope] [ ? [filter] [ ? [Extensions] ] ] ] ] ]
	//If your query includes an attribute, mongod assumes that the query retrieves a list of the DNs which this entity is a member of.
	//
	//If your query does not include an attribute, mongod assumes the query retrieves all entities which the user is member of.
	//
	//For each LDAP DN returned by the query, mongod assigns the authorized user a corresponding role on the admin database. If a role on the on the admin database exactly matches the DN, mongod grants the user the roles and privileges assigned to that role. See the db.createRole() method for more information on creating roles.
	//
	//EXAMPLE
	//
	//This LDAP query returns any groups listed in the LDAP user object’s memberOf attribute.
	//
	//"{USER}?memberOf?base"
	//Your LDAP configuration may not include the memberOf attribute as part of the user schema, may possess a different attribute for reporting group membership, or may not track group membership through attributes. Configure your query with respect to your own unique LDAP configuration.
	//
	//If unset, mongod cannot authorize users using LDAP.
	//
	//This setting can be configured on a running mongod using the setParameter database command.
	//
	//NOTE
	//
	//An explanation of RFC4515, RFC4516 or LDAP queries is out of scope for the MongoDB Documentation. Please review the RFC directly or use your preferred LDAP resource.
	QueryTemplate string `yaml:"queryTemplate"`
}

type Sasl struct {
	//A fully qualified server domain name for the purpose of configuring SASL and Kerberos authentication. The SASL hostname overrides the hostname only for the configuration of SASL and Kerberos.
	//
	//For mongo shell and other MongoDB tools to connect to the new hostName, see the gssapiHostName option in the mongo shell and other tools.
	HostName string `yaml:"hostName"`

	//Registered name of the service using SASL. This option allows you to override the default Kerberos service name component of the Kerberos principal name, on a per-instance basis. If unspecified, the default value is mongodb.
	//
	//MongoDB permits setting this option only at startup. The setParameter can not change this setting.
	//
	//This option is available only in MongoDB Enterprise.
	//
	//IMPORTANT: Ensure that your driver supports alternate service names. For mongo shell and other MongoDB tools to connect to the new serviceName, see the gssapiServiceName option.
	ServiceName string `yaml:"serviceName"`

	//The path to the UNIX domain socket file for saslauthd.
	SaslauthdSocketPath string `yaml:"saslauthdSocketPath"`
}

type Kmip struct {
	//New in version 3.2.
	//
	//Unique KMIP identifier for an existing key within the KMIP server. Include to use the key associated with the identifier as the system key. You can only use the setting the first time you enable encryption for the mongod instance. Requires security.enableEncryption to be true.
	//
	//If unspecified, MongoDB will request that the KMIP server create a new key to utilize as the system key.
	//
	//If the KMIP server cannot locate a key with the specified identifier or the data is already encrypted with a key, MongoDB will throw an error.
	//
	//ENTERPRISE FEATURE
	//
	//Available in MongoDB Enterprise only.
	KeyIdentifier string `yaml:"keyIdentifier"`

	//Default: false
	//
	//New in version 3.2.
	//
	//If true, rotate the master key and re-encrypt the internal keystore.
	//
	//ENTERPRISE FEATURE
	//
	//Available in MongoDB Enterprise only.
	//
	//SEE ALSO: KMIP Master Key Rotation https://docs.mongodb.com/manual/tutorial/rotate-encryption-key/#kmip-master-key-rotation
	RotateMasterKey LogicBoolean `yaml:"rotateMasterKey"`

	//New in version 3.2.
	//
	//Hostname or IP address of key management solution running a KMIP server. Requires security.enableEncryption to be true.
	//
	//ENTERPRISE FEATURE: Available in MongoDB Enterprise only.
	ServerName string `yaml:"serverName"`

	//Default: 5696
	//
	//New in version 3.2.
	//
	//Port number the KMIP server is listening on. Requires that a security.kmip.serverName be provided. Requires security.enableEncryption to be true.
	//
	//ENTERPRISE FEATURE: Available in MongoDB Enterprise only.
	Port int `yaml:"port"`

	//New in version 3.2.
	//
	//String containing the path to the client certificate used for authenticating MongoDB to the KMIP server. Requires that a security.kmip.serverName be provided.
	//
	//NOTE
	//
	//Starting in 4.0, on macOS or Windows, you can use a certificate from the operating system’s secure store instead of a PEM key file. See security.kmip.clientCertificateSelector.
	//
	//ENTERPRISE FEATURE: Available in MongoDB Enterprise only.
	ClientCertificateFile string `yaml:"clientCertificateFile"`

	//New in version 3.2.
	//
	//The password to decrypt the client certificate (i.e. security.kmip.clientCertificateFile), used to authenticate MongoDB to the KMIP server. Use the option only if the certificate is encrypted.
	//
	//ENTERPRISE FEATURE: Available in MongoDB Enterprise only.
	ClientCertificatePassword string `yaml:"clientCertificatePassword"`

	//New in version 4.0: Available on Windows and macOS as an alternative to security.kmip.clientCertificateFile.
	//
	//security.kmip.clientCertificateFile and security.kmip.clientCertificateSelector options are mutually exclusive. You can only specify one.
	//
	//Specifies a certificate property in order to select a matching certificate from the operating system’s certificate store to authenticate MongoDB to the KMIP server.
	//
	//security.kmip.clientCertificateSelector accepts an argument of the format <property>=<value> where the property can be one of the following:
	//
	//Property	Value type	Description
	//subject	ASCII string	Subject name or common name on certificate
	//thumbprint	hex string
	//A sequence of bytes, expressed as hexadecimal, used to identify a public key by its SHA-1 digest.
	//
	//The thumbprint is sometimes referred to as a fingerprint.
	//
	//ENTERPRISE FEATURE: Available in MongoDB Enterprise only.
	ClientCertificateSelector string `yaml:"clientCertificateSelector"`

	//New in version 3.2.
	//
	//Path to CA File. Used for validating secure client connection to KMIP server.
	//
	//NOTE
	//
	//Starting in 4.0, on macOS or Windows, you can use a certificate from the operating system’s secure store instead of a PEM key file. See security.kmip.clientCertificateSelector. When using the secure store, you do not need to, but can, also specify the security.kmip.serverCAFile.
	//
	//ENTERPRISE FEATURE: Available in MongoDB Enterprise only.
	ServerCAFile string `yaml:"serverCAFile"`
}

type Net struct {
	//Default:
	//
	//27017 for mongod (if not a shard member or a config server member) or mongos instance
	//27018 if mongod is a shard member
	//27019 if mongod is a config server member
	//The TCP port on which the MongoDB instance listens for client connections.
	Port NetPort `yaml:"port"`

	//Default: localhost
	//
	//NOTE: Starting in MongoDB 3.6, mongos or mongod bind to localhost by default. See Default Bind to Localhost.
	//
	//The hostnames and/or IP addresses and/or full Unix domain socket paths on which mongos or mongod should listen for client connections. You may attach mongos or mongod to any interface. To bind to multiple addresses, enter a list of comma-separated values.
	//
	//EXAMPLE: localhost,/tmp/mongod.sock
	//
	//You can specify both IPv4 and IPv6 addresses, or hostnames that resolve to an IPv4 or IPv6 address.
	//
	//EXAMPLE: localhost, 2001:0DB8:e132:ba26:0d5c:2774:e7f9:d513
	//
	//NOTE: If specifying an IPv6 address or a hostname that resolves to an IPv6 address to net.bindIp, you must start mongos or mongod with net.ipv6 : true to enable IPv6 support. Specifying an IPv6 address to net.bindIp does not enable IPv6 support.
	//
	//If specifying a link-local IPv6 address (fe80::/10), you must append the zone index to that address (i.e. fe80::<address>%<adapter-name>).
	//
	//EXAMPLE: localhost,fe80::a00:27ff:fee0:1fcf%enp0s3
	//
	//TIP: When possible, use a logical DNS hostname instead of an ip address, particularly when configuring replica set members or sharded cluster members. The use of logical DNS hostnames avoids configuration changes due to ip address changes.
	//
	//WARNING: Before binding to a non-localhost (e.g. publicly accessible) IP address, ensure you have secured your cluster from unauthorized access. For a complete list of security recommendations, see Security Checklist. At minimum, consider enabling authentication and hardening network infrastructure.
	//
	//For more information about IP Binding, refer to the IP Binding documentation.
	//
	//To bind to all IPv4 addresses, enter 0.0.0.0.
	//
	//To bind to all IPv4 and IPv6 addresses, enter ::,0.0.0.0 or starting in MongoDB 4.2, an asterisk "*" (enclose the asterisk in quotes to distinguish from YAML alias nodes). Alternatively, use the net.bindIpAll setting.
	//
	//NOTE: net.bindIp and net.bindIpAll are mutually exclusive. That is, you can specify one or the other, but not both.
	//		The command-line option --bind_ip overrides the configuration file setting net.bindIp.
	BindIp []string `yaml:"bindIp"` //fixme: array to coma list

	//Default: false
	//
	//New in version 3.6.
	//
	//If true, the mongos or mongod instance binds to all IPv4 addresses (i.e. 0.0.0.0). If mongos or mongod starts with net.ipv6 : true, net.bindIpAll also binds to all IPv6 addresses (i.e. ::).
	//
	//mongos or mongod only supports IPv6 if started with net.ipv6 : true. Specifying net.bindIpAll alone does not enable IPv6 support.
	//
	//WARNING: Before binding to a non-localhost (e.g. publicly accessible) IP address, ensure you have secured your cluster from unauthorized access. For a complete list of security recommendations, see Security Checklist. At minimum, consider enabling authentication and hardening network infrastructure.
	//
	//For more information about IP Binding, refer to the IP Binding documentation.
	//
	//Alternatively, set net.bindIp to ::,0.0.0.0 or, starting in MongoDB 4.2, to an asterisk "*" (enclose the asterisk in quotes to distinguish from YAML alias nodes) to bind to all IP addresses.
	//
	//NOTE: net.bindIp and net.bindIpAll are mutually exclusive. Specifying both options causes mongos or mongod to throw an error and terminate.
	BindIpAll LogicBoolean `yaml:"bindIpAll"`

	//Default: 65536
	//
	//The maximum number of simultaneous connections that mongos or mongod will accept. This setting has no effect if it is higher than your operating system’s configured maximum connection tracking threshold.
	//
	//Do not assign too low of a value to this option, or you will encounter errors during normal application operation.
	//
	//This is particularly useful for a mongos if you have a client that creates multiple connections and allows them to timeout rather than closing them.
	//
	//In this case, set maxIncomingConnections to a value slightly higher than the maximum number of connections that the client creates, or the maximum size of the connection pool.
	//
	//This setting prevents the mongos from causing connection spikes on the individual shards. Spikes like these may disrupt the operation and memory allocation of the sharded cluster.
	MaxIncomingConnections int `yaml:"maxIncomingConnections"`

	//Default: true
	//
	//When true, the mongod or mongos instance validates all requests from clients upon receipt to prevent clients from inserting malformed or invalid BSON into a MongoDB database.
	//
	//For objects with a high degree of sub-document nesting, net.wireObjectCheck can have a small impact on performance.
	WireObjectCheck LogicBoolean `yaml:"wireObjectCheck"`

	//Default: false
	//
	//Set net.ipv6 to true to enable IPv6 support. mongos/mongod disables IPv6 support by default.
	//
	//Setting net.ipv6 does not direct the mongos/mongod to listen on any local IPv6 addresses or interfaces. To configure the mongos/mongod to listen on an IPv6 interface, you must either:
	//
	//Configure net.bindIp with one or more IPv6 addresses or hostnames that resolve to IPv6 addresses, or
	//Set net.bindIpAll to true.
	Ipv6 LogicBoolean `yaml:"ipv6"`

	UnixDomainSocket UnixDomainSocket `yaml:"unixDomainSocket"`
	Tls              Tls              `yaml:"tls"`

	Compression Compression `yaml:"compression"`

	//Default: synchronous
	//
	//New in version 3.6.
	//
	//Determines the threading and execution model mongos or mongod uses to execute client requests. The --serviceExecutor option accepts one of the following values:
	//
	//Value	Description
	//synchronous	The mongos or mongod uses synchronous networking and manages its networking thread pool on a per connection basis. Previous versions of MongoDB managed threads in this way.
	//adaptive	The mongos or mongod uses the new experimental asynchronous networking mode with an adaptive thread pool which manages threads on a per request basis. This mode should have more consistent performance and use less resources when there are more inactive connections than database requests.
	ServiceExecutor string `yaml:"serviceExecutor"`
}

type Compression struct {
	//Default: snappy,zstd,zlib
	//
	//New in version 3.4.
	//
	//Specifies the default compressor(s) to use for communication between this mongod or mongos instance and:
	//
	//other members of the deployment if the instance is part of a replica set or a sharded cluster
	//a mongo shell
	//drivers that support the OP_COMPRESSED message format.
	//MongoDB supports the following compressors:
	//
	//snappy
	//zlib (Available starting in MongoDB 3.6)
	//zstd (Available starting in MongoDB 4.2)
	//In versions 3.6 and 4.0, mongod and mongos enable network compression by default with snappy as the compressor.
	//
	//Starting in version 4.2, mongod and mongos instances default to both snappy,zstd,zlib compressors, in that order.
	//
	//To disable network compression, set the value to disabled.
	//
	//IMPORTANT: Messages are compressed when both parties enable network compression. Otherwise, messages between the parties are uncompressed.
	//
	//If you specify multiple compressors, then the order in which you list the compressors matter as well as the communication initiator. For example, if a mongo shell specifies the following network compressors zlib,snappy and the mongod specifies snappy,zlib, messages between mongo shell and mongod uses zlib.
	//
	//If the parties do not share at least one common compressor, messages between the parties are uncompressed. For example, if a mongo shell specifies the network compressor zlib and mongod specifies snappy, messages between mongo shell and mongod are not compressed.
	Compressors string `yaml:"compressors"`
}

type Tls struct {
	//New in version 4.2.
	//
	//Enables TLS used for all network connections. The argument to the net.tls.mode setting can be one of the following:
	//
	//Value	Description
	//disabled	The server does not use TLS.
	//allowTLS	Connections between servers do not use TLS. For incoming connections, the server accepts both TLS and non-TLS.
	//preferTLS	Connections between servers use TLS. For incoming connections, the server accepts both TLS and non-TLS.
	//requireTLS	The server uses and accepts only TLS encrypted connections.
	//If --tlsCAFile or tls.CAFile is not specified and you are not using x.509 authentication, the system-wide CA certificate store will be used when connecting to an TLS-enabled server.
	//
	//If using x.509 authentication, --tlsCAFile or tls.CAFile must be specified unless using --tlsCertificateSelector.
	//
	//For more information about TLS and MongoDB, see Configure mongod and mongos for TLS/SSL and TLS/SSL Configuration for Clients .
	Mode string `yaml:"mode"`

	//New in version 4.2: The .pem file that contains both the TLS certificate and key.
	//
	//Starting with MongoDB 4.0 on macOS or Windows, you can use the net.tls.certificateSelector setting to specify a certificate from the operating system’s secure certificate store instead of a PEM key file. certificateKeyFile and net.tls.certificateSelector are mutually exclusive. You can only specify one.
	//
	//On Linux/BSD, you must specify net.tls.certificateKeyFile when TLS is enabled.
	//
	//On Windows or macOS, you must specify either net.tls.certificateKeyFile or net.tls.certificateSelector when TLS is enabled.
	//
	//IMPORTANT: For Windows only, MongoDB 4.0 and later do not support encrypted PEM files. The mongod fails to start if it encounters an encrypted PEM file. To securely store and access a certificate for use with TLS on Windows, use net.tls.certificateSelector.
	//
	//For more information about TLS and MongoDB, see Configure mongod and mongos for TLS/SSL and TLS/SSL Configuration for Clients .
	CertificateKeyFile string `yaml:"certificateKeyFile"`

	//New in version 4.2: The password to de-crypt the certificate-key file (i.e. certificateKeyFile). Use the net.tls.certificateKeyPassword option only if the certificate-key file is encrypted. In all cases, the mongos or mongod will redact the password from all logging and reporting output.
	//
	//Starting in MongoDB 4.0:
	//
	//On Linux/BSD, if the private key in the PEM file is encrypted and you do not specify the net.tls.certificateKeyFukePassword option, MongoDB will prompt for a passphrase. See TLS/SSL Certificate Passphrase.
	//On macOS, if the private key in the PEM file is encrypted, you must explicitly specify the net.tls.certificateKeyFilePassword option. Alternatively, you can use a certificate from the secure system store (see net.tls.certificateSelector) instead of a PEM key file or use an unencrypted PEM file.
	//On Windows, MongoDB does not support encrypted certificates. The mongod fails if it encounters an encrypted PEM file. Use net.tls.certificateSelector instead.
	//For more information about TLS and MongoDB, see Configure mongod and mongos for TLS/SSL and TLS/SSL Configuration for Clients .
	CertificateKeyFilePassword string `yaml:"certificateKeyFilePassword"`

	//New in version 4.2: Available on Windows and macOS as an alternative to net.tls.certificateKeyFile.
	//
	//Specifies a certificate property in order to select a matching certificate from the operating system’s certificate store to use for TLS/SSL.
	//
	//net.tls.certificateKeyFile and net.tls.certificateSelector options are mutually exclusive. You can only specify one.
	//
	//net.tls.certificateSelector accepts an argument of the format <property>=<value> where the property can be one of the following:
	//
	//Property	Value type	Description
	//subject	ASCII string	Subject name or common name on certificate
	//thumbprint	hex string
	//A sequence of bytes, expressed as hexadecimal, used to identify a public key by its SHA-1 digest.
	//
	//The thumbprint is sometimes referred to as a fingerprint.
	//
	//When using the system SSL certificate store, OCSP (Online Certificate Status Protocol) is used to validate the revocation status of certificates.
	//
	//The mongod searches the operating system’s secure certificate store for the CA certificates required to validate the full certificate chain of the specified TLS certificate. Specifically, the secure certificate store must contain the root CA and any intermediate CA certificates required to build the full certificate chain to the TLS certificate. Do not use net.tls.CAFile or net.tls.clusterFile to specify the root and intermediate CA certificate
	//
	//For example, if the TLS certificate was signed with a single root CA certificate, the secure certificate store must contain that root CA certificate. If the TLS certificate was signed with an intermediate CA certificate, the secure certificate store must contain the intermedia CA certificate and the root CA certificate.
	CertificateSelector string `yaml:"certificateSelector"`

	//New in version 4.2: Available on Windows and macOS as an alternative to net.tls.clusterFile.
	//
	//Specifies a certificate property to select a matching certificate from the operating system’s secure certificate store to use for internal x.509 membership authentication.
	//
	//net.tls.clusterFile and net.tls.clusterCertificateSelector options are mutually exclusive. You can only specify one.
	//
	//net.tls.clusterCertificateSelector accepts an argument of the format <property>=<value> where the property can be one of the following:
	//
	//Property	Value type	Description
	//subject	ASCII string	Subject name or common name on certificate
	//thumbprint	hex string
	//A sequence of bytes, expressed as hexadecimal, used to identify a public key by its SHA-1 digest.
	//
	//The thumbprint is sometimes referred to as a fingerprint.
	//
	//The mongod searches the operating system’s secure certificate store for the CA certificates required to validate the full certificate chain of the specified cluster certificate. Specifically, the secure certificate store must contain the root CA and any intermediate CA certificates required to build the full certificate chain to the cluster certificate. Do not use net.tls.CAFile or net.tls.clusterFile to specify the root and intermediate CA certificate.
	//
	//For example, if the cluster certificate was signed with a single root CA certificate, the secure certificate store must contain that root CA certificate. If the cluster certificate was signed with an intermediate CA certificate, the secure certificate store must contain the intermedia CA certificate and the root CA certificate.
	ClusterCertificateSelector string `yaml:"clusterCertificateSelector"`

	//New in version 4.2: The .pem file that contains the x.509 certificate-key file for membership authentication for the cluster or replica set.
	//
	//Starting with MongoDB 4.0 on macOS or Windows, you can use the net.tls.clusterCertificateSelector option to specify a certificate from the operating system’s secure certificate store instead of a PEM key file. net.tls.clusterFile and net.tls.clusterCertificateSelector options are mutually exclusive. You can only specify one.
	//
	//If net.tls.clusterFile does not specify the .pem file for internal cluster authentication or the alternative net.tls.clusterCertificateSelector, the cluster uses the .pem file specified in the certificateKeyFile setting or the certificate returned by the net.tls.certificateSelector.
	//
	//If using x.509 authentication, --tlsCAFile or tls.CAFile must be specified unless using --tlsCertificateSelector.
	//
	//For more information about TLS and MongoDB, see Configure mongod and mongos for TLS/SSL and TLS/SSL Configuration for Clients .
	//
	//IMPORTANT: For Windows only, MongoDB 4.0 and later do not support encrypted PEM files. The mongod fails to start if it encounters an encrypted PEM file. To securely store and access a certificate for use with membership authentication on Windows, use net.tls.clusterCertificateSelector.
	ClusterFile string `yaml:"clusterFile"`

	//New in version 4.2: The password to de-crypt the x.509 certificate-key file specified with --sslClusterFile. Use the net.tls.clusterPassword option only if the certificate-key file is encrypted. In all cases, the mongos or mongod will redact the password from all logging and reporting output.
	//
	//Starting in MongoDB 4.0:
	//
	//On Linux/BSD, if the private key in the x.509 file is encrypted and you do not specify the net.tls.clusterPassword option, MongoDB will prompt for a passphrase. See TLS/SSL Certificate Passphrase.
	//On macOS, if the private key in the x.509 file is encrypted, you must explicitly specify the net.tls.clusterPassword option. Alternatively, you can either use a certificate from the secure system store (see net.tls.clusterCertificateSelector) instead of a cluster PEM file or use an unencrypted PEM file.
	//On Windows, MongoDB does not support encrypted certificates. The mongod fails if it encounters an encrypted PEM file. Use net.tls.clusterCertificateSelector.
	//For more information about TLS and MongoDB, see Configure mongod and mongos for TLS/SSL and TLS/SSL Configuration for Clients .
	ClusterPassword string `yaml:"clusterPassword"`

	//New in version 4.2: The .pem file that contains the root certificate chain from the Certificate Authority. Specify the file name of the .pem file using relative or absolute paths.
	//
	//Windows/macOS Only: If using net.tls.certificateSelector and/or net.tls.clusterCertificateSelector, do not use net.tls.CAFile to specify the root and intermediate CA certificates. Store all CA certificates required to validate the full trust chain of the net.tls.certificateSelector and/or net.tls.clusterCertificateSelector certificates in the secure certificate store.
	//
	//For more information about TLS and MongoDB, see Configure mongod and mongos for TLS/SSL and TLS/SSL Configuration for Clients .
	CAFile string `yaml:"CAFile"`

	//New in version 4.2: The .pem file that contains the root certificate chain from the Certificate Authority used to validate the certificate presented by a client establishing a connection. Specify the file name of the .pem file using relative or absolute paths. net.tls.clusterCAFile requires that net.tls.CAFile is set.
	//
	//If net.tls.clusterCAFile does not specify the .pem file for validating the certificate from a client establishing a connection, the cluster uses the .pem file specified in the net.tls.CAFile option.
	//
	//net.tls.clusterCAFile lets you use separate Certificate Authorities to verify the client to server and server to client portions of the TLS handshake.
	//
	//Starting in 4.0, on macOS or Windows, you can use a certificate from the operating system’s secure store instead of a PEM key file. See net.tls.clusterCertificateSelector. When using the secure store, you do not need to, but can, also specify the net.tls.clusterCAFile.
	//
	//Windows/macOS Only: If using net.tls.certificateSelector and/or net.tls.clusterCertificateSelector, do not use net.tls.clusterCAFile to specify the root and intermediate CA certificates. Store all CA certificates required to validate the full trust chain of the net.tls.certificateSelector and/or net.tls.clusterCertificateSelector certificates in the secure certificate store.
	//
	//For more information about TLS and MongoDB, see Configure mongod and mongos for TLS/SSL and TLS/SSL Configuration for Clients .
	ClusterCAFile string `yaml:"clusterCAFile"`

	//New in version 4.2.
	//
	//The the .pem file that contains the Certificate Revocation List. Specify the file name of the .pem file using relative or absolute paths.
	//
	//NOTE: Starting in MongoDB 4.0, you cannot specify net.tls.CRLFile on macOS. Use net.tls.certificateSelector instead.
	//
	//For more information about TLS and MongoDB, see Configure mongod and mongos for TLS/SSL and TLS/SSL Configuration for Clients .
	CRLFile string `yaml:"CRLFile"`

	//New in version 4.2.
	//
	//For clients that do not present certificates, mongos or mongod bypasses TLS/SSL certificate validation when establishing the connection.
	//
	//For clients that present a certificate, however, mongos or mongod performs certificate validation using the root certificate chain specified by CAFile and reject clients with invalid certificates.
	//
	//Use the net.tls.allowConnectionsWithoutCertificates option if you have a mixed deployment that includes clients that do not or cannot present certificates to the mongos or mongod.
	//
	//For more information about TLS and MongoDB, see Configure mongod and mongos for TLS/SSL and TLS/SSL Configuration for Clients .
	AllowConnectionsWithoutCertificates LogicBoolean `yaml:"allowConnectionsWithoutCertificates"`

	//New in version 4.2.
	//
	//Enable or disable the validation checks for TLS certificates on other servers in the cluster and allows the use of invalid certificates to connect.
	//
	//NOTE
	//
	//If you specify --tlsAllowInvalidCertificates or tls.allowInvalidCertificates: true when using x.509 authentication, an invalid certificate is only sufficient to establish a TLS connection but is insufficient for authentication.
	//
	//When using the net.tls.allowInvalidCertificates setting, MongoDB logs a warning regarding the use of the invalid certificate.
	//
	//For more information about TLS and MongoDB, see Configure mongod and mongos for TLS/SSL and TLS/SSL Configuration for Clients .
	AllowInvalidCertificates LogicBoolean `yaml:"allowInvalidCertificates"`

	//Default: false
	//
	//When net.tls.allowInvalidHostnames is true, MongoDB disables the validation of the hostnames in TLS certificates, allowing mongod to connect to MongoDB instances if the hostname their certificates do not match the specified hostname.
	//
	//For more information about TLS and MongoDB, see Configure mongod and mongos for TLS/SSL and TLS/SSL Configuration for Clients .
	AllowInvalidHostnames LogicBoolean `yaml:"allowInvalidHostnames"`

	//New in version 4.2.
	//
	//Prevents a MongoDB server running with TLS from accepting incoming connections that use a specific protocol or protocols. To specify multiple protocols, use a comma separated list of protocols.
	//
	//net.tls.disabledProtocols recognizes the following protocols: TLS1_0, TLS1_1, TLS1_2, and starting in version 4.0.4 (and 3.6.9), TLS1_3.
	//
	//On macOS, you cannot disable TLS1_1 and leave both TLS1_0 and TLS1_2 enabled. You must disable at least one of the other two, for example, TLS1_0,TLS1_1.
	//To list multiple protocols, specify as a comma separated list of protocols. For example TLS1_0,TLS1_1.
	//Specifying an unrecognized protocol will prevent the server from starting.
	//The specified disabled protocols overrides any default disabled protocols.
	//Starting in version 4.0, MongoDB disables the use of TLS 1.0 if TLS 1.1+ is available on the system. To enable the disabled TLS 1.0, specify none to net.tls.disabledProtocols. See Disable TLS 1.0.
	//
	//Members of replica sets and sharded clusters must speak at least one protocol in common.
	//
	//SEE ALSO: Disallow Protocols https://docs.mongodb.com/manual/tutorial/configure-ssl/#ssl-disallow-protocols
	DisabledProtocols string `yaml:"disabledProtocols"`

	//New in version 4.2.
	//
	//Enable or disable the use of the FIPS mode of the TLS library for the mongos or mongod. Your system must have a FIPS compliant library to use the net.tls.FIPSMode option.
	//
	//NOTE: FIPS-compatible TLS/SSL is available only in MongoDB Enterprise. See Configure MongoDB for FIPS for more information.
	FIPSMode LogicBoolean `yaml:"FIPSMode"`
}

type Free struct {
	//New in version 4.0: Available for MongoDB Community Edition.
	//
	//Enables or disables free MongoDB Cloud monitoring. cloud.monitoring.free.state accepts the following values:
	//
	//runtime
	//Default. You can enable or disable free monitoring during runtime.
	//
	//To enable or disable free monitoring during runtime, see db.enableFreeMonitoring() and db.disableFreeMonitoring().
	//
	//To enable or disable free monitoring during runtime when running with access control, users must have required privileges. See db.enableFreeMonitoring() and db.disableFreeMonitoring() for details.
	//
	//on	Enables free monitoring at startup; i.e. registers for free monitoring. When enabled at startup, you cannot disable free monitoring during runtime.
	//off	Disables free monitoring at startup, regardless of whether you have previously registered for free monitoring. When disabled at startup, you cannot enable free monitoring during runtime.
	//Once enabled, the free monitoring state remains enabled until explicitly disabled. That is, you do not need to re-enable each time you start the server.
	//
	//For the corresponding command-line option, see --enableFreeMonitoring.
	State FreeState `yaml:"state"`

	//New in version 4.0: Available for MongoDB Community Edition.
	//
	//Optional tag to describe environment context. The tag can be sent as part of the free MongoDB Cloud monitoring registration at start up.
	//
	//For the corresponding command-line option, see --freeMonitoringTag.
	Tags string `yaml:"tags"`
}

type Monitoring struct {
	Free Free `yaml:"free"`
}

type Cloud struct {
	Monitoring Monitoring `yaml:"monitoring"`
}

type UnixDomainSocket struct {
	//Default: true
	//
	//Enable or disable listening on the UNIX domain socket. net.unixDomainSocket.enabled applies only to Unix-based systems.
	//
	//When net.unixDomainSocket.enabled is true, mongos or mongod listens on the UNIX socket.
	//
	//The mongos or mongod process always listens on the UNIX socket unless one of the following is true:
	//
	//net.unixDomainSocket.enabled is false
	//--nounixsocket is set. The command line option takes precedence over the configuration file setting.
	//net.bindIp is not set
	//net.bindIp does not specify localhost or its associated IP address
	//mongos or mongod installed from official .deb and .rpm packages have the bind_ip configuration set to 127.0.0.1 by default.
	Enabled LogicBoolean `yaml:"enabled"`

	//Default: /tmp
	//
	//The path for the UNIX socket. net.unixDomainSocket.pathPrefix applies only to Unix-based systems.
	//
	//If this option has no value, the mongos or mongod process creates a socket with /tmp as a prefix. MongoDB creates and listens on a UNIX socket unless one of the following is true:
	//
	//net.unixDomainSocket.enabled is false
	//--nounixsocket is set
	//net.bindIp is not set
	//net.bindIp does not specify localhost or its associated IP address
	PathPrefix LogicBoolean `yaml:"pathPrefix"`

	//Default: 0700
	//
	//Sets the permission for the UNIX domain socket file.
	//
	//net.unixDomainSocket.filePermissions applies only to Unix-based systems.
	FilePermissions int `yaml:"filePermissions"`
}

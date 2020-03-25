package iotmaker_db_mongodb_config

import (
	"bytes"
	"errors"
	"fmt"
	"reflect"
	"strconv"
)

//https://docs.mongodb.com/manual/reference/configuration-options/
//todo:https://docs.mongodb.com/manual/reference/configuration-options/#mongos-only-options

type Configuration struct {
	//Default: 0
	//
	//The default log message verbosity level for components. The verbosity level
	//determines the amount of Informational and Debug messages MongoDB outputs. [2]
	//
	//The verbosity level can range from 0 to 5:
	//    0 is the MongoDB’s default log verbosity level, to include Informational
	//      messages.
	//    1 to 5 increases the verbosity level to include Debug messages.
	//To use a different verbosity level for a named component, use the component’s
	//verbosity setting. For example, use the systemLog.component.accessControl.verbosity
	//to set the verbosity level specifically for ACCESS components.
	//
	//See the systemLog.component.<name>.verbosity settings for specific component
	//verbosity settings.
	//
	//For various ways to set the log verbosity level, see Configure Log Verbosity Levels.
	//
	//[2]	Starting in version 4.2, MongoDB includes the Debug verbosity level (1-5) in the
	//log messages. For example, if the verbosity level is 2, MongoDB logs D2. In previous
	//versions, MongoDB log messages only specified D for Debug level.
	Verbosity Verbosity `yaml:"verbosity"`

	//Run mongos or mongod in a quiet mode that attempts to limit the amount of output.
	//
	//systemLog.quiet is not recommended for production systems as it may make tracking
	//problems during particular connections much more difficult.
	Quiet LogicBoolean `yaml:"quiet"`

	//Print verbose information for debugging. Use for additional logging for
	//support-related troubleshooting.
	TraceAllExceptions LogicBoolean `yaml:"traceAllExceptions"`

	//Default: user
	//
	//The facility level used when logging messages to syslog. The value you specify must
	//be supported by your operating system’s implementation of syslog. To use this option,
	//you must set systemLog.destination to syslog.
	SyslogFacility int `yaml:"syslogFacility"`

	//The path of the log file to which mongod or mongos should send all diagnostic logging
	//information, rather than the standard output or the host’s syslog. MongoDB creates
	//the log file at the specified path.
	//
	//The Linux package init scripts do not expect systemLog.path to change from the
	//defaults.
	//If you use the Linux packages and change systemLog.path, you will have to use your
	//own init scripts and disable the built-in scripts.
	Path string `yaml:"path"`

	//Default: false
	//
	//When true, mongos or mongod appends new entries to the end of the existing log file
	//when the mongos or mongod instance restarts. Without this option, mongod will back up
	//the existing log and create a new file.
	LogAppend LogicBoolean `yaml:"logAppend"`

	//Default: rename
	//
	//The behavior for the logRotate command. Specify either rename or reopen:
	//
	//    Value  |  Description
	//    ---------------------------------------------------------------------------------
	//    rename |  renames the log file.
	//    ---------------------------------------------------------------------------------
	//    reopen |  closes and reopens the log file following the typical Linux/Unix log
	//           |  rotate behavior. Use reopen when using the Linux/Unix logrotate utility
	//           |  to avoid log loss.
	//
	//If you specify reopen, you must also set systemLog.logAppend to true.
	LogRotate LogRotate `yaml:"logRotate"`

	//The destination to which MongoDB sends all log output. Specify either file or syslog.
	//If you specify file, you must also specify systemLog.path.
	//
	//If you do not specify systemLog.destination, MongoDB sends all log output to standard
	//output.
	//
	//WARNING: The syslog daemon generates timestamps when it logs a message, not when
	//MongoDB issues the message. This can lead to misleading timestamps for log entries,
	//especially when the system is under heavy load. We recommend using the file option
	//for production systems to ensure accurate timestamps.
	Destination string `yaml:"destination"`

	//Default: iso8601-local
	//
	//The time format for timestamps in log messages. Specify one of the following values:
	//
	//    Value         |  Description
	//    ---------------------------------------------------------------------------------
	//    ctime         |  Displays timestamps as Wed Dec 31 18:17:54.811.
	//    ---------------------------------------------------------------------------------
	//    iso8601-utc   |  Displays timestamps in Coordinated Universal Time (UTC) in the
	//                  |  ISO-8601 format. For example, for New York at the start of the
	//                  |  Epoch: 1970-01-01T00:00:00.000Z
	//    ---------------------------------------------------------------------------------
	//    iso8601-local |  Displays timestamps in local time in the ISO-8601 format.
	//                  |  For example, for New York at the start of the Epoch:
	//                  |  1969-12-31T19:00:00.000-0500
	//
	TimeStampFormat TimeStampFormat `yaml:"timeStampFormat"`

	//NOTE: Starting in version 4.2, MongoDB includes the Debug verbosity level (1-5) in
	//the log messages. For example, if the verbosity level is 2, MongoDB logs D2. In
	//previous versions, MongoDB log messages only specified D for Debug level.
	Component         Component         `yaml:"component"`
	ProcessManagement ProcessManagement `yaml:"processManagement"`
	Cloud             Cloud             `yaml:"cloud"`

	//Changed in version 4.2: MongoDB 4.2 deprecates ssl options in favor of tls options
	//with identical functionality.
	Net Net `yaml:"net"`

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
	//The interval (in seconds) mongod waits between external user cache flushes. After
	//mongod flushes the external user cache, MongoDB reacquires authorization data from
	//the LDAP server the next time an LDAP-authorized user issues an operation.
	//
	//Increasing the value specified increases the amount of time mongod and the LDAP
	//server can be out of sync, but reduces the load on the LDAP server. Conversely,
	//decreasing the value specified decreases the time mongod and the LDAP server can be
	//out of sync while increasing the load on the LDAP server.
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

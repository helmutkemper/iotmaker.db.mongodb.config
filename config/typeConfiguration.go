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

type Configuration struct {
	SystemLog         SystemLog         `yaml:"systemLog"`
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
	SetParameter interface{} `yaml:"-"` //todo: fazer

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
	LdapUserCacheInvalidationInterval int `yaml:"-"`

	Storage            Storage            `yaml:"-"`
	OperationProfiling OperationProfiling `yaml:"-"`
	Replication        Replication        `yaml:"-"`
	Sharding           Sharding           `yaml:"-"`
	AuditLog           AuditLog           `yaml:"-"`
	Snmp               Snmp               `yaml:"-"`
}

func (el *Configuration) getTagData(tag reflect.StructTag) (string, string) {
	var tagName, tagValue string

	tagName = "yaml"
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

func (el *Configuration) writeDoubleQuotes(buffer *bytes.Buffer) {
	buffer.WriteString(`"`)
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

func (el *Configuration) writeKey(buffer *bytes.Buffer, key string) {
	buffer.WriteString(key)
	buffer.WriteString(": ")
	buffer.WriteString("\r\n")
}

func (el *Configuration) writeKeyValueForString(buffer *bytes.Buffer, key, value string) {
	buffer.WriteString(key)
	buffer.WriteString(": ")
	buffer.WriteString(`"`)
	buffer.WriteString(value)
	buffer.WriteString(`"`)
	buffer.WriteString("\r\n")
}

func (el *Configuration) ToYaml() (error, []byte) {
	var buffer bytes.Buffer
	var err = el.ToText(2, reflect.ValueOf(*el), &buffer)
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
		case "iotmaker_db_mongodb_config.Verbosity":

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

		case "iotmaker_db_mongodb_config.LogicBoolean":

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
			el.writeKeyValueForString(buffer, tagValue, field.Interface().(string))

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

		case "iotmaker_db_mongodb_config.TimeStampFormat":
			if field.Interface().(TimeStampFormat).String() == "" {
				continue
			}

			el.writeSpaces(buffer, level)
			str := field.Interface().(TimeStampFormat).String()
			el.writeKeyValue(buffer, tagValue, str)

		case "iotmaker_db_mongodb_config.LogRotate":
			if field.Interface().(LogRotate).String() == "" {
				continue
			}

			el.writeSpaces(buffer, level)
			str := field.Interface().(LogRotate).String()
			el.writeKeyValue(buffer, tagValue, str)

		case "iotmaker_db_mongodb_config.SystemLog":

			if reflect.DeepEqual(SystemLog{}, field.Interface().(SystemLog)) == true {
				continue
			}

			el.writeSpaces(buffer, level)
			el.writeKey(buffer, tagValue)
			err := el.ToText(level+2, field, buffer)
			if err != nil {
				return err
			}

		case "iotmaker_db_mongodb_config.Component":

			if reflect.DeepEqual(Component{}, field.Interface().(Component)) == true {
				continue
			}

			el.writeSpaces(buffer, level)
			el.writeKey(buffer, tagValue)
			err := el.ToText(level+2, field, buffer)
			if err != nil {
				return err
			}

		case "iotmaker_db_mongodb_config.Cloud":

			if reflect.DeepEqual(Cloud{}, field.Interface().(Cloud)) == true {
				continue
			}

			el.writeSpaces(buffer, level)
			el.writeKey(buffer, tagValue)
			err := el.ToText(level+2, field, buffer)
			if err != nil {
				return err
			}

		case "iotmaker_db_mongodb_config.ProcessManagement":
			fallthrough
		case "iotmaker_db_mongodb_config.AccessControlLog":
			fallthrough
		case "iotmaker_db_mongodb_config.CommandLog":
			fallthrough
		case "iotmaker_db_mongodb_config.ControlLog":
			fallthrough
		case "iotmaker_db_mongodb_config.FtdcLog":
			fallthrough
		case "iotmaker_db_mongodb_config.GeoLog":
			fallthrough
		case "iotmaker_db_mongodb_config.IndexLog":
			fallthrough
		case "iotmaker_db_mongodb_config.NetworkLog":
			fallthrough
		case "iotmaker_db_mongodb_config.QueryLog":
			fallthrough
		case "iotmaker_db_mongodb_config.ReplicationLog":
			fallthrough
		case "iotmaker_db_mongodb_config.TransactionLog":
			fallthrough
		case "iotmaker_db_mongodb_config.WriteLog":
			fallthrough
		case "iotmaker_db_mongodb_config.ElectionLog":
			fallthrough
		case "iotmaker_db_mongodb_config.HeartbeatsLog":
			fallthrough
		case "iotmaker_db_mongodb_config.InitialSyncLog":
			fallthrough
		case "iotmaker_db_mongodb_config.RollbackLog":
			fallthrough
		case "iotmaker_db_mongodb_config.ShardingLog":
			fallthrough
		case "iotmaker_db_mongodb_config.StorageLog":
			fallthrough
		case "iotmaker_db_mongodb_config.Net":
			fallthrough
		case "iotmaker_db_mongodb_config.Security":

			el.writeSpaces(buffer, level)
			el.writeKey(buffer, tagValue)
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

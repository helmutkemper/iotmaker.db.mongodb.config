package iotmaker_db_mongodb_config

import (
	"bytes"
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

//https://docs.mongodb.com/manual/reference/configuration-options/
//todo:https://docs.mongodb.com/manual/reference/configuration-options/#mongos-only-options

type Configuration struct {
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
	//    setParameter:
	//       <parameter1>: <value1>
	//       <parameter2>: <value2>
	//
	//For example, to specify the enableLocalhostAuthBypass in the configuration file:
	//
	//    setParameter:
	//       enableLocalhostAuthBypass: false
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
	SystemLog          SystemLog          `yaml:"systemLog"`
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

func (el *Configuration) writeSpaces(buffer *bytes.Buffer, spaces int) {
	for i := 0; i != spaces; i += 1 {
		buffer.WriteString(" ")
	}
}

func (el *Configuration) writeKey(buffer *bytes.Buffer, key string) {
	buffer.WriteString(key)
	buffer.WriteString(":")
	buffer.WriteString("\n")
}

func (el *Configuration) writeValue(buffer *bytes.Buffer, key string) {
	buffer.WriteString(key)
	buffer.WriteString("\n")
}

func (el *Configuration) writeKeyValue(buffer *bytes.Buffer, key, value string) {
	buffer.WriteString(key)
	buffer.WriteString(": ")
	buffer.WriteString(value)
	buffer.WriteString("\n")
}

func (el *Configuration) ToYaml(level int) (error, []byte) {
	var buffer bytes.Buffer
	var err = el.ToText(level, reflect.ValueOf(*el), &buffer)
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

		case "interface {}":

			switch tagValue {
			case "setParameter":

				switch converted := el.SetParameter.(type) {
				case EnableLocalhostAuthBypass:
					if converted == 0 {
						continue
					}

					el.writeSpaces(buffer, level)
					el.writeKey(buffer, tagValue)

					if converted == -1 {
						el.writeSpaces(buffer, level+2)
						el.writeKeyValue(buffer, "enableLocalhostAuthBypass", "false")
						continue
					}

					if converted == 1 {
						el.writeSpaces(buffer, level+2)
						el.writeKeyValue(buffer, "enableLocalhostAuthBypass", "true")
						continue
					}
				}

			default:
				fmt.Printf("todo: %v\n\n\n", tagValue)
			}

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

			if field.Interface().(LogicBoolean) == 0 {
				continue
			}

			if field.Interface().(LogicBoolean) == -1 {
				el.writeSpaces(buffer, level)
				el.writeKeyValue(buffer, tagValue, "false")
				continue
			}

			if field.Interface().(LogicBoolean) == 1 {
				el.writeSpaces(buffer, level)
				el.writeKeyValue(buffer, tagValue, "true")
				continue
			}

			return errors.New("logical must be -1 for false, 0 for not set and 1 for true")

		case "iotmaker_db_mongodb_config.ComaList":

			if len(field.Interface().(ComaList)) == 0 {
				continue
			}

			if len(field.Interface().(ComaList)) == 1 {
				el.writeSpaces(buffer, level)
				el.writeKeyValue(buffer, tagValue, `"`+strings.Join(field.Interface().(ComaList), `, `)+`"`)
				continue
			}

			el.writeSpaces(buffer, level)
			el.writeKey(buffer, tagValue)
			for _, value := range field.Interface().(ComaList) {
				el.writeSpaces(buffer, level+2)
				el.writeValue(buffer, `- "`+value+`"`)
			}

		case "string":
			if field.Interface().(string) == "" {
				continue
			}

			el.writeSpaces(buffer, level)
			el.writeKeyValue(buffer, tagValue, `"`+field.Interface().(string)+`"`)

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

		case "iotmaker_db_mongodb_config.SystemLogDestination":
			if field.Interface().(SystemLogDestination) == 0 {
				continue
			}

			//todo: you must also specify systemLog.path

			el.writeSpaces(buffer, level)
			str := field.Interface().(SystemLogDestination).String()
			el.writeKeyValue(buffer, tagValue, str)

		case "iotmaker_db_mongodb_config.TimeStampFormat":
			if field.Interface().(TimeStampFormat) == 0 {
				continue
			}

			el.writeSpaces(buffer, level)
			str := field.Interface().(TimeStampFormat).String()
			el.writeKeyValue(buffer, tagValue, str)

		case "iotmaker_db_mongodb_config.FreeState":
			if field.Interface().(FreeState) == 0 {
				continue
			}

			el.writeSpaces(buffer, level)
			str := field.Interface().(FreeState).String()
			el.writeKeyValue(buffer, tagValue, str)

		case "iotmaker_db_mongodb_config.Compressor":
			if field.Interface().(Compressor) == 0 {
				continue
			}

			el.writeSpaces(buffer, level)
			str := field.Interface().(Compressor).String()
			el.writeKeyValue(buffer, tagValue, str)

		case "iotmaker_db_mongodb_config.LogRotate":
			if field.Interface().(LogRotate) == 0 {
				continue
			}

			el.writeSpaces(buffer, level)
			str := field.Interface().(LogRotate).String()
			el.writeKeyValue(buffer, tagValue, str)

		case "iotmaker_db_mongodb_config.Journal":

			if reflect.DeepEqual(Journal{}, field.Interface().(Journal)) == true {
				continue
			}

			el.writeSpaces(buffer, level)
			el.writeKey(buffer, tagValue)
			err := el.ToText(level+2, field, buffer)
			if err != nil {
				return err
			}

		case "iotmaker_db_mongodb_config.WiredTiger":

			if reflect.DeepEqual(WiredTiger{}, field.Interface().(WiredTiger)) == true {
				continue
			}

			el.writeSpaces(buffer, level)
			el.writeKey(buffer, tagValue)
			err := el.ToText(level+2, field, buffer)
			if err != nil {
				return err
			}

		case "iotmaker_db_mongodb_config.InMemory":

			if reflect.DeepEqual(InMemory{}, field.Interface().(InMemory)) == true {
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

			if reflect.DeepEqual(ProcessManagement{}, field.Interface().(ProcessManagement)) == true {
				continue
			}

			el.writeSpaces(buffer, level)
			el.writeKey(buffer, tagValue)
			err := el.ToText(level+2, field, buffer)
			if err != nil {
				return err
			}

		case "iotmaker_db_mongodb_config.AccessControlLog":

			if reflect.DeepEqual(AccessControlLog{}, field.Interface().(AccessControlLog)) == true {
				continue
			}

			el.writeSpaces(buffer, level)
			el.writeKey(buffer, tagValue)
			err := el.ToText(level+2, field, buffer)
			if err != nil {
				return err
			}

		case "iotmaker_db_mongodb_config.CommandLog":

			if reflect.DeepEqual(CommandLog{}, field.Interface().(CommandLog)) == true {
				continue
			}

			el.writeSpaces(buffer, level)
			el.writeKey(buffer, tagValue)
			err := el.ToText(level+2, field, buffer)
			if err != nil {
				return err
			}

		case "iotmaker_db_mongodb_config.ControlLog":

			if reflect.DeepEqual(ControlLog{}, field.Interface().(ControlLog)) == true {
				continue
			}

			el.writeSpaces(buffer, level)
			el.writeKey(buffer, tagValue)
			err := el.ToText(level+2, field, buffer)
			if err != nil {
				return err
			}

		case "iotmaker_db_mongodb_config.FtdcLog":

			if reflect.DeepEqual(FtdcLog{}, field.Interface().(FtdcLog)) == true {
				continue
			}

			el.writeSpaces(buffer, level)
			el.writeKey(buffer, tagValue)
			err := el.ToText(level+2, field, buffer)
			if err != nil {
				return err
			}

		case "iotmaker_db_mongodb_config.GeoLog":

			if reflect.DeepEqual(GeoLog{}, field.Interface().(GeoLog)) == true {
				continue
			}

			el.writeSpaces(buffer, level)
			el.writeKey(buffer, tagValue)
			err := el.ToText(level+2, field, buffer)
			if err != nil {
				return err
			}

		case "iotmaker_db_mongodb_config.IndexLog":

			if reflect.DeepEqual(IndexLog{}, field.Interface().(IndexLog)) == true {
				continue
			}

			el.writeSpaces(buffer, level)
			el.writeKey(buffer, tagValue)
			err := el.ToText(level+2, field, buffer)
			if err != nil {
				return err
			}

		case "iotmaker_db_mongodb_config.NetworkLog":

			if reflect.DeepEqual(NetworkLog{}, field.Interface().(NetworkLog)) == true {
				continue
			}

			el.writeSpaces(buffer, level)
			el.writeKey(buffer, tagValue)
			err := el.ToText(level+2, field, buffer)
			if err != nil {
				return err
			}

		case "iotmaker_db_mongodb_config.QueryLog":

			if reflect.DeepEqual(QueryLog{}, field.Interface().(QueryLog)) == true {
				continue
			}

			el.writeSpaces(buffer, level)
			el.writeKey(buffer, tagValue)
			err := el.ToText(level+2, field, buffer)
			if err != nil {
				return err
			}

		case "iotmaker_db_mongodb_config.ReplicationLog":

			if reflect.DeepEqual(ReplicationLog{}, field.Interface().(ReplicationLog)) == true {
				continue
			}

			el.writeSpaces(buffer, level)
			el.writeKey(buffer, tagValue)
			err := el.ToText(level+2, field, buffer)
			if err != nil {
				return err
			}

		case "iotmaker_db_mongodb_config.TransactionLog":

			if reflect.DeepEqual(TransactionLog{}, field.Interface().(TransactionLog)) == true {
				continue
			}

			el.writeSpaces(buffer, level)
			el.writeKey(buffer, tagValue)
			err := el.ToText(level+2, field, buffer)
			if err != nil {
				return err
			}

		case "iotmaker_db_mongodb_config.WriteLog":

			if reflect.DeepEqual(WriteLog{}, field.Interface().(WriteLog)) == true {
				continue
			}

			el.writeSpaces(buffer, level)
			el.writeKey(buffer, tagValue)
			err := el.ToText(level+2, field, buffer)
			if err != nil {
				return err
			}

		case "iotmaker_db_mongodb_config.ElectionLog":

			if reflect.DeepEqual(ElectionLog{}, field.Interface().(ElectionLog)) == true {
				continue
			}

			el.writeSpaces(buffer, level)
			el.writeKey(buffer, tagValue)
			err := el.ToText(level+2, field, buffer)
			if err != nil {
				return err
			}

		case "iotmaker_db_mongodb_config.HeartbeatsLog":

			if reflect.DeepEqual(HeartbeatsLog{}, field.Interface().(HeartbeatsLog)) == true {
				continue
			}

			el.writeSpaces(buffer, level)
			el.writeKey(buffer, tagValue)
			err := el.ToText(level+2, field, buffer)
			if err != nil {
				return err
			}

		case "iotmaker_db_mongodb_config.InitialSyncLog":

			if reflect.DeepEqual(InitialSyncLog{}, field.Interface().(InitialSyncLog)) == true {
				continue
			}

			el.writeSpaces(buffer, level)
			el.writeKey(buffer, tagValue)
			err := el.ToText(level+2, field, buffer)
			if err != nil {
				return err
			}

		case "iotmaker_db_mongodb_config.RollbackLog":

			if reflect.DeepEqual(RollbackLog{}, field.Interface().(RollbackLog)) == true {
				continue
			}

			el.writeSpaces(buffer, level)
			el.writeKey(buffer, tagValue)
			err := el.ToText(level+2, field, buffer)
			if err != nil {
				return err
			}

		case "iotmaker_db_mongodb_config.ShardingLog":

			if reflect.DeepEqual(ShardingLog{}, field.Interface().(ShardingLog)) == true {
				continue
			}

			el.writeSpaces(buffer, level)
			el.writeKey(buffer, tagValue)
			err := el.ToText(level+2, field, buffer)
			if err != nil {
				return err
			}

		case "iotmaker_db_mongodb_config.StorageLog":

			if reflect.DeepEqual(StorageLog{}, field.Interface().(StorageLog)) == true {
				continue
			}

			el.writeSpaces(buffer, level)
			el.writeKey(buffer, tagValue)
			err := el.ToText(level+2, field, buffer)
			if err != nil {
				return err
			}

		case "iotmaker_db_mongodb_config.Net":

			if reflect.DeepEqual(Net{}, field.Interface().(Net)) == true {
				continue
			}

			el.writeSpaces(buffer, level)
			el.writeKey(buffer, tagValue)
			err := el.ToText(level+2, field, buffer)
			if err != nil {
				return err
			}

		case "iotmaker_db_mongodb_config.Snmp":

			if reflect.DeepEqual(Snmp{}, field.Interface().(Snmp)) == true {
				continue
			}

			el.writeSpaces(buffer, level)
			el.writeKey(buffer, tagValue)
			err := el.ToText(level+2, field, buffer)
			if err != nil {
				return err
			}

		case "iotmaker_db_mongodb_config.AuditLog":

			if reflect.DeepEqual(AuditLog{}, field.Interface().(AuditLog)) == true {
				continue
			}

			el.writeSpaces(buffer, level)
			el.writeKey(buffer, tagValue)
			err := el.ToText(level+2, field, buffer)
			if err != nil {
				return err
			}

		case "iotmaker_db_mongodb_config.Sharding":

			if reflect.DeepEqual(Sharding{}, field.Interface().(Sharding)) == true {
				continue
			}

			el.writeSpaces(buffer, level)
			el.writeKey(buffer, tagValue)
			err := el.ToText(level+2, field, buffer)
			if err != nil {
				return err
			}

		case "iotmaker_db_mongodb_config.Replication":

			if reflect.DeepEqual(Replication{}, field.Interface().(Replication)) == true {
				continue
			}

			el.writeSpaces(buffer, level)
			el.writeKey(buffer, tagValue)
			err := el.ToText(level+2, field, buffer)
			if err != nil {
				return err
			}

		case "iotmaker_db_mongodb_config.OperationProfiling":

			if reflect.DeepEqual(OperationProfiling{}, field.Interface().(OperationProfiling)) == true {
				continue
			}

			el.writeSpaces(buffer, level)
			el.writeKey(buffer, tagValue)
			err := el.ToText(level+2, field, buffer)
			if err != nil {
				return err
			}

		case "iotmaker_db_mongodb_config.Storage":

			if reflect.DeepEqual(Storage{}, field.Interface().(Storage)) == true {
				continue
			}

			el.writeSpaces(buffer, level)
			el.writeKey(buffer, tagValue)
			err := el.ToText(level+2, field, buffer)
			if err != nil {
				return err
			}

		case "iotmaker_db_mongodb_config.Ldap":

			if reflect.DeepEqual(Ldap{}, field.Interface().(Ldap)) == true {
				continue
			}

			el.writeSpaces(buffer, level)
			el.writeKey(buffer, tagValue)
			err := el.ToText(level+2, field, buffer)
			if err != nil {
				return err
			}

		case "iotmaker_db_mongodb_config.Sasl":

			if reflect.DeepEqual(Sasl{}, field.Interface().(Sasl)) == true {
				continue
			}

			el.writeSpaces(buffer, level)
			el.writeKey(buffer, tagValue)
			err := el.ToText(level+2, field, buffer)
			if err != nil {
				return err
			}

		case "iotmaker_db_mongodb_config.Kmip":

			if reflect.DeepEqual(Kmip{}, field.Interface().(Kmip)) == true {
				continue
			}

			el.writeSpaces(buffer, level)
			el.writeKey(buffer, tagValue)
			err := el.ToText(level+2, field, buffer)
			if err != nil {
				return err
			}

		case "iotmaker_db_mongodb_config.Mode":
			if field.Interface().(Mode) == 0 {
				continue
			}

			el.writeSpaces(buffer, level)
			str := field.Interface().(Mode).String()
			el.writeKeyValue(buffer, tagValue, str)

		case "iotmaker_db_mongodb_config.ServiceExecutor":
			if field.Interface().(ServiceExecutor) == 0 {
				continue
			}

			el.writeSpaces(buffer, level)
			str := field.Interface().(ServiceExecutor).String()
			el.writeKeyValue(buffer, tagValue, str)

		case "iotmaker_db_mongodb_config.Compression":

			if reflect.DeepEqual(Compression{}, field.Interface().(Compression)) == true {
				continue
			}

			el.writeSpaces(buffer, level)
			el.writeKey(buffer, tagValue)
			err := el.ToText(level+2, field, buffer)
			if err != nil {
				return err
			}

		case "iotmaker_db_mongodb_config.Tls":

			if reflect.DeepEqual(Tls{}, field.Interface().(Tls)) == true {
				continue
			}

			el.writeSpaces(buffer, level)
			el.writeKey(buffer, tagValue)
			err := el.ToText(level+2, field, buffer)
			if err != nil {
				return err
			}

		case "iotmaker_db_mongodb_config.UnixDomainSocket":

			if reflect.DeepEqual(UnixDomainSocket{}, field.Interface().(UnixDomainSocket)) == true {
				continue
			}

			el.writeSpaces(buffer, level)
			el.writeKey(buffer, tagValue)
			err := el.ToText(level+2, field, buffer)
			if err != nil {
				return err
			}

		case "iotmaker_db_mongodb_config.NetPort":
			if field.Interface().(NetPort) == 0 {
				continue
			}

			el.writeSpaces(buffer, level)
			str := strconv.FormatInt(int64(field.Interface().(NetPort)), 10)
			el.writeKeyValue(buffer, tagValue, str)

		case "iotmaker_db_mongodb_config.Monitoring":

			if reflect.DeepEqual(Monitoring{}, field.Interface().(Monitoring)) == true {
				continue
			}

			el.writeSpaces(buffer, level)
			el.writeKey(buffer, tagValue)
			err := el.ToText(level+2, field, buffer)
			if err != nil {
				return err
			}

		case "iotmaker_db_mongodb_config.RecoveryLog":

			if reflect.DeepEqual(RecoveryLog{}, field.Interface().(RecoveryLog)) == true {
				continue
			}

			el.writeSpaces(buffer, level)
			el.writeKey(buffer, tagValue)
			err := el.ToText(level+2, field, buffer)
			if err != nil {
				return err
			}

		case "iotmaker_db_mongodb_config.JournalLog":

			if reflect.DeepEqual(JournalLog{}, field.Interface().(JournalLog)) == true {
				continue
			}

			el.writeSpaces(buffer, level)
			el.writeKey(buffer, tagValue)
			err := el.ToText(level+2, field, buffer)
			if err != nil {
				return err
			}

		case "iotmaker_db_mongodb_config.Free":

			if reflect.DeepEqual(Free{}, field.Interface().(Free)) == true {
				continue
			}

			el.writeSpaces(buffer, level)
			el.writeKey(buffer, tagValue)
			err := el.ToText(level+2, field, buffer)
			if err != nil {
				return err
			}

		case "iotmaker_db_mongodb_config.Security":

			if reflect.DeepEqual(Security{}, field.Interface().(Security)) == true {
				continue
			}

			el.writeSpaces(buffer, level)
			el.writeKey(buffer, tagValue)
			err := el.ToText(level+2, field, buffer)
			if err != nil {
				return err
			}

		default:
			fmt.Printf("\nyaml(): %v[ %d ]: %s - %s = %v\n", tagValue, i, field.Type(), field.Interface(), tagName)
		}
	}

	return nil
}

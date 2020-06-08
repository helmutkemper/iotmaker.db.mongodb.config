package factoryMongoDBConfig

import iotmaker_db_mongodb_config "github.com/helmutkemper/iotmaker.db.mongodb.config/config"

func NewBasicConfigWithPortAndEphemeralData(port int) *iotmaker_db_mongodb_config.Configuration {
	return &iotmaker_db_mongodb_config.Configuration{
		SystemLog: iotmaker_db_mongodb_config.SystemLog{
			Destination: iotmaker_db_mongodb_config.KSystemLogDestinationFile,
			Path:        "/var/log/mongodb/mongod.log",
			LogAppend:   iotmaker_db_mongodb_config.TRUE,
		},
		Storage: iotmaker_db_mongodb_config.Storage{
			Journal: iotmaker_db_mongodb_config.Journal{
				Enabled: iotmaker_db_mongodb_config.TRUE,
			},
		},
		ProcessManagement: iotmaker_db_mongodb_config.ProcessManagement{
			Fork: iotmaker_db_mongodb_config.FALSE,
		},
		Net: iotmaker_db_mongodb_config.Net{
			BindIp: iotmaker_db_mongodb_config.ComaList{
				"127.0.0.1",
				"0.0.0.0",
			},
			Port: iotmaker_db_mongodb_config.NetPort(port),
		},
		SetParameter: iotmaker_db_mongodb_config.EnableLocalhostAuthBypass(iotmaker_db_mongodb_config.TRUE),
	}
}

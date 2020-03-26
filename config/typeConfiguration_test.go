package iotmaker_db_mongodb_config

import "fmt"

func ExampleConfiguration_ToYaml() {
	conf := &Configuration{
		SystemLog: SystemLog{
			Verbosity:          1,
			Quiet:              FALSE,
			TraceAllExceptions: TRUE,
			SyslogFacility:     "user",
			Path:               "/var/log/mongodb/mongod.log",
			LogAppend:          TRUE,
			LogRotate:          KLogRotateReopen,
			Destination:        KSystemLogDestinationSyslog,
			TimeStampFormat:    KTimeStampFormatISO8601UTC,
			Component: Component{
				AccessControl: AccessControlLog{
					Verbosity: 1,
				},
				Command: CommandLog{
					Verbosity: 1,
				},
				Control: ControlLog{
					Verbosity: 1,
				},
				Ftdc: FtdcLog{
					Verbosity: 1,
				},
				Geo: GeoLog{
					Verbosity: 1,
				},
				Index: IndexLog{
					Verbosity: 1,
				},
				Network: NetworkLog{
					Verbosity: 1,
				},
				Query: QueryLog{
					Verbosity: 1,
				},
				Replication: ReplicationLog{
					Verbosity: 1,
					Election: ElectionLog{
						Verbosity: 1,
					},
					Heartbeats: HeartbeatsLog{
						Verbosity: 1,
					},
					InitialSync: InitialSyncLog{
						Verbosity: 1,
					},
					Rollback: RollbackLog{
						Verbosity: 1,
					},
					Sharding: ShardingLog{
						Verbosity: 1,
					},
					Storage: StorageLog{
						Verbosity: 1,
						Journal: JournalLog{
							Verbosity: 1,
						},
						Recovery: RecoveryLog{
							Verbosity: 1,
						},
					},
				},
				Transaction: TransactionLog{
					Verbosity: 1,
				},
				Write: WriteLog{
					Verbosity: 1,
				},
			},
		},
	}

	err, yml := conf.ToYaml(0)
	if err != nil {
		fmt.Printf("error: %v\n", err.Error())
		return
	}

	fmt.Printf("%s\n", yml)
	//output:
	//
}

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
		ProcessManagement: ProcessManagement{
			Fork:         FALSE,
			PidFilePath:  "/etc/init.d",
			TimeZoneInfo: "/usr/share/zoneinfo",
		},
		Cloud: Cloud{
			Monitoring: Monitoring{
				Free: Free{
					State: KFreeStateOn,
					Tags:  "tags",
				},
			},
		},
		Net: Net{
			Port:                   KNetPortMongodOrMongosInstance,
			BindIp:                 ComaList{"0.0.0.0", "/tmp/mongod.sock"},
			BindIpAll:              FALSE,
			MaxIncomingConnections: 65536,
			WireObjectCheck:        TRUE,
			Ipv6:                   FALSE,
			UnixDomainSocket: UnixDomainSocket{
				Enabled:         FALSE,
				PathPrefix:      "/tmp",
				FilePermissions: 0700,
			},
			Tls: Tls{
				Mode:                                KModePreferTLS,
				CertificateKeyFile:                  "/CertificateKeyFile",
				CertificateKeyFilePassword:          "******",
				CertificateSelector:                 "subject=ffffffff",
				ClusterCertificateSelector:          "subject=ffffffff",
				ClusterFile:                         "/ClusterFile",
				ClusterPassword:                     "******",
				CAFile:                              "/CAFile",
				ClusterCAFile:                       "/ClusterCAFile",
				CRLFile:                             "/CRLFile",
				AllowConnectionsWithoutCertificates: TRUE,
				AllowInvalidCertificates:            FALSE,
				AllowInvalidHostNames:               TRUE,
				DisabledProtocols:                   nil,
				FIPSMode:                            FALSE,
			},
			Compression: Compression{
				Compressors: KCompressorZStd,
			},
			ServiceExecutor: KServiceExecutorSynchronous,
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

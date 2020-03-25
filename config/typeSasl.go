package iotmaker_db_mongodb_config

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

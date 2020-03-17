package iotmaker_db_mongodb_config

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

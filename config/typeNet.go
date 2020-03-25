package iotmaker_db_mongodb_config

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
	BindIp []string `yaml:"bindIp"`

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
	Compression      Compression      `yaml:"compression"`

	//Default: synchronous
	//
	//New in version 3.6.
	//
	//Determines the threading and execution model mongos or mongod uses to execute client requests. The --serviceExecutor option accepts one of the following values:
	//
	//Value	Description
	//synchronous	The mongos or mongod uses synchronous networking and manages its networking thread pool on a per connection basis. Previous versions of MongoDB managed threads in this way.
	//adaptive	The mongos or mongod uses the new experimental asynchronous networking mode with an adaptive thread pool which manages threads on a per request basis. This mode should have more consistent performance and use less resources when there are more inactive connections than database requests.
	ServiceExecutor ServiceExecutor `yaml:"serviceExecutor"`
}

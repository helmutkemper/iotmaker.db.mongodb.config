package iotmaker_db_mongodb_config

import iotmaker_db_mongodb_config "github.com/helmut.kemper/iotmaker.db.mongodb.config/config"

//Default: localhost
//
//NOTE: Starting in MongoDB 3.6, mongos or mongod bind to localhost by default. See Default Bind to Localhost.
//
//The hostnames and/or IP addresses and/or full Unix domain socket paths on which mongos or mongod should listen for client connections. You may attach mongos or mongod to any interface. To bind to multiple addresses, enter a list of comma-separated values.
//
//EXAMPLE: "localhost", "/tmp/mongod.sock"
//
//You can specify both IPv4 and IPv6 addresses, or hostnames that resolve to an IPv4 or IPv6 address.
//
//EXAMPLE: "localhost", "2001:0DB8:e132:ba26:0d5c:2774:e7f9:d513"
//
//NOTE: If specifying an IPv6 address or a hostname that resolves to an IPv6 address to net.bindIp, you must start mongos or mongod with net.ipv6 : true to enable IPv6 support. Specifying an IPv6 address to net.bindIp does not enable IPv6 support.
//
//If specifying a link-local IPv6 address (fe80::/10), you must append the zone index to that address (i.e. fe80::<address>%<adapter-name>).
//
//EXAMPLE: "localhost","fe80::a00:27ff:fee0:1fcf%enp0s3"
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
//
//The command-line option --bind_ip overrides the configuration file setting net.bindIp.
func SetBindIpList(configuration *iotmaker_db_mongodb_config.Configuration, list ...string) {
	configuration.Net.BindIp = list
}

func SetBindAll(configuration *iotmaker_db_mongodb_config.Configuration, value bool) {
	configuration.Net.BindIpAll = value
}

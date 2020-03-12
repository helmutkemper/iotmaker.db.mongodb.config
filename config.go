package main

import (
	"fmt"
)

type AccessControl struct {
	//Default: 0
	//
	//The log message verbosity level for components related to access control. See ACCESS components.
	//
	//The verbosity level can range from 0 to 5:
	//
	//0 is the MongoDB’s default log verbosity level, to include Informational messages.
	//1 to 5 increases the verbosity level to include Debug messages.
	Verbosity int
}

type Command struct {
	//Default: 0
	//
	//The log message verbosity level for components related to commands. See COMMAND components.
	//
	//The verbosity level can range from 0 to 5:
	//
	//0 is the MongoDB’s default log verbosity level, to include Informational messages.
	//1 to 5 increases the verbosity level to include Debug messages.
	Verbosity int
}

type Control struct {
	//Default: 0
	//
	//The log message verbosity level for components related to control operations. See CONTROL components.
	//
	//The verbosity level can range from 0 to 5:
	//
	//0 is the MongoDB’s default log verbosity level, to include Informational messages.
	//1 to 5 increases the verbosity level to include Debug messages.
	Verbosity int
}

type Ftdc struct {
	//Default: 0
	//
	//New in version 3.2.
	//
	//The log message verbosity level for components related to diagnostic data collection operations. See FTDC components.
	//
	//The verbosity level can range from 0 to 5:
	//
	//0 is the MongoDB’s default log verbosity level, to include Informational messages.
	//1 to 5 increases the verbosity level to include Debug messages.
	Verbosity int
}

type Geo struct {
	//Default: 0
	//
	//The log message verbosity level for components related to geospatial parsing operations. See GEO components.
	//
	//The verbosity level can range from 0 to 5:
	//
	//0 is the MongoDB’s default log verbosity level, to include Informational messages.
	//1 to 5 increases the verbosity level to include Debug messages.
	Verbosity int
}

type Index struct {
	//Default: 0
	//
	//The log message verbosity level for components related to indexing operations. See INDEX components.
	//
	//The verbosity level can range from 0 to 5:
	//
	//0 is the MongoDB’s default log verbosity level, to include Informational messages.
	//1 to 5 increases the verbosity level to include Debug messages.
	Verbosity int
}

type Network struct {
	//Default: 0
	//
	//The log message verbosity level for components related to networking operations. See NETWORK components.
	//
	//The verbosity level can range from 0 to 5:
	//
	//0 is the MongoDB’s default log verbosity level, to include Informational messages.
	//1 to 5 increases the verbosity level to include Debug messages.
	Verbosity int
}

type Query struct {
	// Default: 0
	//
	//The log message verbosity level for components related to query operations. See QUERY components.
	//
	//The verbosity level can range from 0 to 5:
	//
	//0 is the MongoDB’s default log verbosity level, to include Informational messages.
	//1 to 5 increases the verbosity level to include Debug messages.
	Verbosity int
}

type Replication struct {
	//Default: 0
	//
	//The log message verbosity level for components related to replication. See REPL components.
	//
	//The verbosity level can range from 0 to 5:
	//
	//0 is the MongoDB’s default log verbosity level, to include Informational messages.
	//1 to 5 increases the verbosity level to include Debug messages.
	Verbosity int

	//Default: 0
	//
	//New in version 4.2.
	//
	//The log message verbosity level for components related to election. See ELECTION components.
	//
	//If systemLog.component.replication.election.verbosity is unset, systemLog.component.replication.verbosity level also applies to election components.
	//
	//The verbosity level can range from 0 to 5:
	//
	//0 is the MongoDB’s default log verbosity level, to include Informational messages.
	//1 to 5 increases the verbosity level to include Debug messages.
	Election int

	Heartbeats  Heartbeats
	InitialSync InitialSync
	Rollback    Rollback
	Sharding    Sharding
	Storage     Storage
}

type Storage struct {
	//Default: 0
	//
	//The log message verbosity level for components related to storage. See STORAGE components.
	//
	//If systemLog.component.storage.journal.verbosity is unset, systemLog.component.storage.verbosity level also applies to journaling components.
	//
	//The verbosity level can range from 0 to 5:
	//
	//0 is the MongoDB’s default log verbosity level, to include Informational messages.
	//1 to 5 increases the verbosity level to include Debug messages.
	Verbosity int

	Journal  Journal
	Recovery Recovery
}

type Recovery struct {
	//Default: 0
	//
	//New in version 4.0.
	//
	//The log message verbosity level for components related to recovery. See RECOVERY components.
	//
	//If systemLog.component.storage.recovery.verbosity is unset, systemLog.component.storage.verbosity level also applies to recovery components.
	//
	//The verbosity level can range from 0 to 5:
	//
	//0 is the MongoDB’s default log verbosity level, to include Informational messages.
	//1 to 5 increases the verbosity level to include Debug messages.
	Verbosity int
}

type Journal struct {
	//Default: 0
	//
	//The log message verbosity level for components related to journaling. See JOURNAL components.
	//
	//If systemLog.component.storage.journal.verbosity is unset, the journaling components have the same verbosity level as the parent storage components: i.e. either the systemLog.component.storage.verbosity level if set or the default verbosity level.
	//
	//The verbosity level can range from 0 to 5:
	//
	//0 is the MongoDB’s default log verbosity level, to include Informational messages.
	//1 to 5 increases the verbosity level to include Debug messages.
	Verbosity int
}

type Sharding struct {
	//Default: 0
	//
	//The log message verbosity level for components related to sharding. See SHARDING components.
	//
	//The verbosity level can range from 0 to 5:
	//
	//0 is the MongoDB’s default log verbosity level, to include Informational messages.
	//1 to 5 increases the verbosity level to include Debug messages.
	Verbosity int
}

type Rollback struct {
	//Default: 0
	//
	//New in version 3.6.
	//
	//The log message verbosity level for components related to rollback. See ROLLBACK components.
	//
	//If systemLog.component.replication.rollback.verbosity is unset, systemLog.component.replication.verbosity level also applies to rollback components.
	//
	//The verbosity level can range from 0 to 5:
	//
	//0 is the MongoDB’s default log verbosity level, to include Informational messages.
	//1 to 5 increases the verbosity level to include Debug messages.
	Verbosity int
}

type Heartbeats struct {
	//Default: 0
	//
	//New in version 3.6.
	//
	//The log message verbosity level for components related to heartbeats. See REPL_HB components.
	//
	//If systemLog.component.replication.heartbeats.verbosity is unset, systemLog.component.replication.verbosity level also applies to heartbeats components.
	//
	//The verbosity level can range from 0 to 5:
	//
	//0 is the MongoDB’s default log verbosity level, to include Informational messages.
	//1 to 5 increases the verbosity level to include Debug messages.
	Verbosity int
}

type InitialSync struct {
	//Default: 0
	//
	//New in version 4.2.
	//
	//The log message verbosity level for components related to initialSync. See INITSYNC components.
	//
	//If systemLog.component.replication.initialSync.verbosity is unset, systemLog.component.replication.verbosity level also applies to initialSync components.
	//
	//The verbosity level can range from 0 to 5:
	//
	//0 is the MongoDB’s default log verbosity level, to include Informational messages.
	//1 to 5 increases the verbosity level to include Debug messages.
	Verbosity int
}

type Component struct {
	AccessControl AccessControl
	Command       Command
	Control       Control
	Ftdc          Ftdc
	Geo           Geo
	Index         Index
	Network       Network
	Query         Query
	Replication   Replication
	Transaction   Transaction
	Write         Write
}

type Write struct {
	//Default: 0
	//
	//The log message verbosity level for components related to write operations. See WRITE components.
	//
	//The verbosity level can range from 0 to 5:
	//
	//0 is the MongoDB’s default log verbosity level, to include Informational messages.
	//1 to 5 increases the verbosity level to include Debug messages.
	Verbosity int
}

type Transaction struct {
	//Default: 0
	//
	//New in version 4.0.2.
	//
	//The log message verbosity level for components related to transaction. See TXN components.
	//
	//The verbosity level can range from 0 to 5:
	//
	//0 is the MongoDB’s default log verbosity level, to include Informational messages.
	//1 to 5 increases the verbosity level to include Debug messages.
	Verbosity int
}

type ProcessManagement struct {
	//Default: false
	//
	//Enable a daemon mode that runs the mongos or mongod process in the background. By default mongos or mongod does not run as a daemon: typically you will run mongos or mongod as a daemon, either by using processManagement.fork or by using a controlling process that handles the daemonization process (e.g. as with upstart and systemd).
	//
	//The processManagement.fork option is not supported on Windows.
	//
	//The Linux package init scripts do not expect processManagement.fork to change from the defaults. If you use the Linux packages and change processManagement.fork, you will have to use your own init scripts and disable the built-in scripts.
	Fork bool

	//Specifies a file location to store the process ID (PID) of the mongos or mongod process . The user running the the mongod or mongos process must be able to write to this path. If the processManagement.pidFilePath option is not specified, the process does not create a PID file. This option is generally only useful in combination with the processManagement.fork setting.
	//
	//LINUX: On Linux, PID file management is generally the responsibility of your distro’s init system: usually a service file in the /etc/init.d directory, or a systemd unit file registered with systemctl. Only use the processManagement.pidFilePath option if you are not using one of these init systems. For more information, please see the respective Installation Guide for your operating system.
	//
	//MACOS: On macOS, PID file management is generally handled by brew. Only use the processManagement.pidFilePath option if you are not using brew on your macOS system. For more information, please see the respective Installation Guide for your operating system.
	PidFilePath string

	//The full path from which to load the time zone database. If this option is not provided, then MongoDB will use its built-in time zone database.
	//
	//The configuration file included with Linux and macOS packages sets the time zone database path to /usr/share/zoneinfo by default.
	//
	//The built-in time zone database is a copy of the Olson/IANA time zone database. It is updated along with MongoDB releases, but the release cycle of the time zone database differs from the release cycle of MongoDB. A copy of the most recent release of the time zone database can be downloaded from https://downloads.mongodb.org/olson_tz_db/timezonedb-latest.zip.
	TimeZoneInfo string
}

type Configuration struct {
	//Default: 0
	//
	//The default log message verbosity level for components. The verbosity level determines the amount of Informational and Debug messages MongoDB outputs. [2]
	//
	//The verbosity level can range from 0 to 5:
	//
	//0 is the MongoDB’s default log verbosity level, to include Informational messages.
	//1 to 5 increases the verbosity level to include Debug messages.
	//To use a different verbosity level for a named component, use the component’s verbosity setting. For example, use the systemLog.component.accessControl.verbosity to set the verbosity level specifically for ACCESS components.
	//
	//See the systemLog.component.<name>.verbosity settings for specific component verbosity settings.
	//
	//For various ways to set the log verbosity level, see Configure Log Verbosity Levels.
	//
	//[2]	Starting in version 4.2, MongoDB includes the Debug verbosity level (1-5) in the log messages. For example, if the verbosity level is 2, MongoDB logs D2. In previous versions, MongoDB log messages only specified D for Debug level.
	Verbosity int

	//Run mongos or mongod in a quiet mode that attempts to limit the amount of output.
	//
	//systemLog.quiet is not recommended for production systems as it may make tracking problems during particular connections much more difficult.
	Quiet bool

	//Print verbose information for debugging. Use for additional logging for support-related troubleshooting.
	TraceAllExceptions bool

	//Default: user
	//
	//The facility level used when logging messages to syslog. The value you specify must be supported by your operating system’s implementation of syslog. To use this option, you must set systemLog.destination to syslog.
	SyslogFacility string

	//The path of the log file to which mongod or mongos should send all diagnostic logging information, rather than the standard output or the host’s syslog. MongoDB creates the log file at the specified path.
	//
	//The Linux package init scripts do not expect systemLog.path to change from the defaults. If you use the Linux packages and change systemLog.path, you will have to use your own init scripts and disable the built-in scripts.
	Path string

	//Default: false
	//
	//When true, mongos or mongod appends new entries to the end of the existing log file when the mongos or mongod instance restarts. Without this option, mongod will back up the existing log and create a new file.
	LogAppend bool

	//Default: rename
	//
	//The behavior for the logRotate command. Specify either rename or reopen:
	//
	//rename renames the log file.
	//
	//reopen closes and reopens the log file following the typical Linux/Unix log rotate behavior. Use reopen when using the Linux/Unix logrotate utility to avoid log loss.
	//
	//If you specify reopen, you must also set systemLog.logAppend to true.
	LogRotate string

	//The destination to which MongoDB sends all log output. Specify either file or syslog. If you specify file, you must also specify systemLog.path.
	//
	//If you do not specify systemLog.destination, MongoDB sends all log output to standard output.
	//
	//WARNING: The syslog daemon generates timestamps when it logs a message, not when MongoDB issues the message. This can lead to misleading timestamps for log entries, especially when the system is under heavy load. We recommend using the file option for production systems to ensure accurate timestamps.
	Destination string

	//Default: iso8601-local
	//
	//The time format for timestamps in log messages. Specify one of the following values:
	//
	//Value	Description
	//ctime		Displays timestamps as Wed Dec 31 18:17:54.811.
	//iso8601-utc	Displays timestamps in Coordinated Universal Time (UTC) in the ISO-8601 format. For example, for New York at the start of the Epoch: 1970-01-01T00:00:00.000Z
	//iso8601-local	Displays timestamps in local time in the ISO-8601 format. For example, for New York at the start of the Epoch: 1969-12-31T19:00:00.000-0500
	TimeStampFormat string

	//NOTE: Starting in version 4.2, MongoDB includes the Debug verbosity level (1-5) in the log messages. For example, if the verbosity level is 2, MongoDB logs D2. In previous versions, MongoDB log messages only specified D for Debug level.
	Component         Component
	ProcessManagement ProcessManagement
	Cloud             Cloud

	//Changed in version 4.2: MongoDB 4.2 deprecates ssl options in favor of tls options with identical functionality.
	Net Net
}

type Net struct {
	//Default:
	//
	//27017 for mongod (if not a shard member or a config server member) or mongos instance
	//27018 if mongod is a shard member
	//27019 if mongod is a config server member
	//The TCP port on which the MongoDB instance listens for client connections.
	Port int

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
	BindIp []string

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
	BindIpAll bool

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
	MaxIncomingConnections int

	//Default: true
	//
	//When true, the mongod or mongos instance validates all requests from clients upon receipt to prevent clients from inserting malformed or invalid BSON into a MongoDB database.
	//
	//For objects with a high degree of sub-document nesting, net.wireObjectCheck can have a small impact on performance.
	WireObjectCheck bool

	//Default: false
	//
	//Set net.ipv6 to true to enable IPv6 support. mongos/mongod disables IPv6 support by default.
	//
	//Setting net.ipv6 does not direct the mongos/mongod to listen on any local IPv6 addresses or interfaces. To configure the mongos/mongod to listen on an IPv6 interface, you must either:
	//
	//Configure net.bindIp with one or more IPv6 addresses or hostnames that resolve to IPv6 addresses, or
	//Set net.bindIpAll to true.
	Ipv6 bool

	UnixDomainSocket UnixDomainSocket
}

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
	Enabled bool

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
	PathPrefix bool

	//Default: 0700
	//
	//Sets the permission for the UNIX domain socket file.
	//
	//net.unixDomainSocket.filePermissions applies only to Unix-based systems.
	FilePermissions int

	Tls Tls
}

type Tls struct {
}

type Free struct {
	//New in version 4.0: Available for MongoDB Community Edition.
	//
	//Enables or disables free MongoDB Cloud monitoring. cloud.monitoring.free.state accepts the following values:
	//
	//runtime
	//Default. You can enable or disable free monitoring during runtime.
	//
	//To enable or disable free monitoring during runtime, see db.enableFreeMonitoring() and db.disableFreeMonitoring().
	//
	//To enable or disable free monitoring during runtime when running with access control, users must have required privileges. See db.enableFreeMonitoring() and db.disableFreeMonitoring() for details.
	//
	//on	Enables free monitoring at startup; i.e. registers for free monitoring. When enabled at startup, you cannot disable free monitoring during runtime.
	//off	Disables free monitoring at startup, regardless of whether you have previously registered for free monitoring. When disabled at startup, you cannot enable free monitoring during runtime.
	//Once enabled, the free monitoring state remains enabled until explicitly disabled. That is, you do not need to re-enable each time you start the server.
	//
	//For the corresponding command-line option, see --enableFreeMonitoring.
	State string

	//New in version 4.0: Available for MongoDB Community Edition.
	//
	//Optional tag to describe environment context. The tag can be sent as part of the free MongoDB Cloud monitoring registration at start up.
	//
	//For the corresponding command-line option, see --freeMonitoringTag.
	Tags string
}

type Monitoring struct {
	Free Free
}

type Cloud struct {
	Monitoring Monitoring
}

func main() {
	fmt.Println("Hello, playground")
}

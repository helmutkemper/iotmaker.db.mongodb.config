package iotmaker_db_mongodb_config

type ProcessManagement struct {
	//Default: false
	//
	//Enable a daemon mode that runs the mongos or mongod process in the background. By default mongos or mongod does not run as a daemon: typically you will run mongos or mongod as a daemon, either by using processManagement.fork or by using a controlling process that handles the daemonization process (e.g. as with upstart and systemd).
	//
	//The processManagement.fork option is not supported on Windows.
	//
	//The Linux package init scripts do not expect processManagement.fork to change from the defaults. If you use the Linux packages and change processManagement.fork, you will have to use your own init scripts and disable the built-in scripts.
	Fork LogicBoolean `yaml:"fork"`

	//Specifies a file location to store the process ID (PID) of the mongos or mongod process . The user running the the mongod or mongos process must be able to write to this path. If the processManagement.pidFilePath option is not specified, the process does not create a PID file. This option is generally only useful in combination with the processManagement.fork setting.
	//    LINUX: On Linux, PID file management is generally the responsibility of your distro’s init system:
	//      usually a service file in the /etc/init.d directory, or a systemd unit file registered with systemctl. Only use the processManagement.pidFilePath option if you are not using one of these init systems. For more information, please see the respective Installation Guide for your operating system.
	//    MACOS: On macOS, PID file management is generally handled by brew. Only use the processManagement.pidFilePath option if you are not using brew on your macOS system. For more information, please see the respective Installation Guide for your operating system.
	PidFilePath string `yaml:"pidFilePath"`

	//The full path from which to load the time zone database. If this option is not provided, then MongoDB will use its built-in time zone database.
	//
	//The configuration file included with Linux and macOS packages sets the time zone database path to /usr/share/zoneinfo by default.
	//
	//The built-in time zone database is a copy of the Olson/IANA time zone database. It is updated along with MongoDB releases, but the release cycle of the time zone database differs from the release cycle of MongoDB. A copy of the most recent release of the time zone database can be downloaded from https://downloads.mongodb.org/olson_tz_db/timezonedb-latest.zip.
	TimeZoneInfo string `yaml:"timeZoneInfo"`
}

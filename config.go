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
	
}

type Component struct {
	AccessControl AccessControl
	Command Command
	Control Control
	Ftdc ftdc
	Geo Geo
	Index Index
	Network Network
	
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
	component Component
}

func main() {
	fmt.Println("Hello, playground")
}

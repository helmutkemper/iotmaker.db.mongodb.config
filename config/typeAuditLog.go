package iotmaker_db_mongodb_config

type AuditLog struct {
	//When set, auditLog.destination enables auditing and specifies where mongos or mongod
	//sends all audit events.
	//
	//auditLog.destination can have one of the following values:
	//
	//    Value   |  Description
	//    ---------------------------------------------------------------------------------
	//    syslog  |  Output the audit events to syslog in JSON format. Not available on
	//            |  Windows. Audit messages have a syslog severity level of info and a
	//            |  facility level of user.
	//            |  The syslog message limit can result in the truncation of audit
	//            |  messages. The auditing system will neither detect the truncation nor
	//            |  error upon its occurrence.
	//    ---------------------------------------------------------------------------------
	//    console |  Output the audit events to stdout in JSON format.
	//    ---------------------------------------------------------------------------------
	//    file    |  Output the audit events to the file specified in auditLog.path in the
	//            |  format specified in auditLog.format.
	//
	//NOTE: Available only in MongoDB Enterprise and MongoDB Atlas.
	Destination Destination `yaml:"destination"`

	//The format of the output file for auditing if destination is file.
	//The auditLog.format option can have one of the following values:
	//
	//    Value  |  Description
	//    ---------------------------------------------------------------------------------
	//    JSON   |  Output the audit events in JSON format to the file specified in
	//           |  auditLog.path.
	//    ---------------------------------------------------------------------------------
	//    BSON   |  Output the audit events in BSON binary format to the file specified in
	//           |  auditLog.path.
	//
	//Printing audit events to a file in JSON format degrades server performance more than
	//printing to a file in BSON format.
	//
	//NOTE: Available only in MongoDB Enterprise and MongoDB Atlas.
	Format string `yaml:"format"`

	//The output file for auditing if destination has value of file. The auditLog.path
	//option can take either a full path name or a relative path name.
	//
	//NOTE: Available only in MongoDB Enterprise and MongoDB Atlas.
	Path string `yaml:"path"`

	//Type: string representation of a document
	//
	//The filter to limit the types of operations the audit system records. The option
	//takes a string representation of a query document of the form:
	//
	//{ <field1>: <expression1>, ... }
	//The <field> can be any field in the audit message, including fields returned in the
	//param document.
	//The <expression> is a query condition expression.
	//
	//To specify an audit filter, enclose the filter document in single quotes to pass the
	//document as a string.
	//
	//To specify the audit filter in a configuration file, you must use the YAML format of
	//the configuration file.
	//
	//NOTE:
	//    Available only in MongoDB Enterprise and MongoDB Atlas.
	Filter string `yaml:"filter"`
}

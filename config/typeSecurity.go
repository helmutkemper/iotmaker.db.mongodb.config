package iotmaker_db_mongodb_config

type Security struct {
	//The path to a key file that stores the shared secret that MongoDB instances use to authenticate to each other in a sharded cluster or replica set. keyFile implies security.authorization. See Internal/Membership Authentication for more information.
	//
	//Starting in MongoDB 4.2, keyfiles for internal membership authentication use YAML format to allow for multiple keys in a keyfile. The YAML format accepts content of:
	//
	//a single key string (same as in earlier versions),
	//multiple key strings (each string must be enclosed in quotes), or
	//sequence of key strings.
	//The YAML format is compatible with the existing single-key keyfiles that use the text file format.
	KeyFile string `yaml:"keyFile"`

	//Default: keyFile
	//
	//The authentication mode used for cluster authentication. If you use internal x.509 authentication, specify so here. This option can have one of the following values:
	//
	//Value       Description
	//keyFile	    Use a keyfile for authentication. Accept only keyfiles.
	//sendKeyFile	For rolling upgrade purposes. Send a keyfile for authentication but can accept both keyfiles and x.509 certificates.
	//sendX509	  For rolling upgrade purposes. Send the x.509 certificate for authentication but can accept both keyfiles and x.509 certificates.
	//x509	      Recommended. Send the x.509 certificate for authentication and accept only x.509 certificates.
	//If --tlsCAFile or tls.CAFile is not specified and you are not using x.509 authentication, the system-wide CA certificate store will be used when connecting to an TLS-enabled server.
	//
	//If using x.509 authentication, --tlsCAFile or tls.CAFile must be specified unless using --tlsCertificateSelector.
	//
	//For more information about TLS and MongoDB, see Configure mongod and mongos for TLS/SSL and TLS/SSL Configuration for Clients .
	ClusterAuthMode ClusterAuthMode `yaml:"clusterAuthMode"`

	//Default: disabled
	//
	//Enable or disable Role-Based Access Control (RBAC) to govern each user’s access to database resources and operations.
	//
	//Set this option to one of the following:
	//
	//Value	Description
	//enabled	A user can access only the database resources and actions for which they have been granted privileges.
	//disabled	A user can access any database and perform any action.
	//See Role-Based Access Control for more information.
	//
	//The security.authorization setting is available only for mongod.
	Authorization Authorization `yaml:"authorization"`

	//Default: false
	//
	//New in version 3.4: Allows the mongod or mongos to accept and create authenticated and non-authenticated connections to and from other mongod and mongos instances in the deployment. Used for performing rolling transition of replica sets or sharded clusters from a no-auth configuration to internal authentication. Requires specifying a internal authentication mechanism such as security.keyFile.
	//
	//For example, if using keyfiles for internal authentication, the mongod or mongos creates an authenticated connection with any mongod or mongos in the deployment using a matching keyfile. If the security mechanisms do not match, the mongod or mongos utilizes a non-authenticated connection instead.
	//
	//A mongod or mongos running with security.transitionToAuth does not enforce user access controls. Users may connect to your deployment without any access control checks and perform read, write, and administrative operations.
	//
	//NOTE
	//
	//A mongod or mongos running with internal authentication and without security.transitionToAuth requires clients to connect using user access controls. Update clients to connect to the mongod or mongos using the appropriate user prior to restarting mongod or mongos without security.transitionToAuth.
	TransitionToAuth LogicBoolean `yaml:"transitionToAuth"`

	//Default: true
	//
	//Enables or disables the server-side JavaScript execution. When disabled, you cannot use operations that perform server-side execution of JavaScript code, such as the $where query operator, mapReduce command and the db.collection.mapReduce() method.
	JavascriptEnabled LogicBoolean `yaml:"javascriptEnabled"`

	//New in version 3.4: Available in MongoDB Enterprise only.
	//
	//A mongod or mongos running with security.redactClientLogData redacts any message accompanying a given log event before logging. This prevents the mongod or mongos from writing potentially sensitive data stored on the database to the diagnostic log. Metadata such as error or operation codes, line numbers, and source file names are still visible in the logs.
	//
	//Use security.redactClientLogData in conjunction with Encryption at Rest and TLS/SSL (Transport Encryption) to assist compliance with regulatory requirements.
	//
	//For example, a MongoDB deployment might store Personally Identifiable Information (PII) in one or more collections. The mongod or mongos logs events such as those related to CRUD operations, sharding metadata, etc. It is possible that the mongod or mongos may expose PII as a part of these logging operations. A mongod or mongos running with security.redactClientLogData removes any message accompanying these events before being output to the log, effectively removing the PII.
	//
	//Diagnostics on a mongod or mongos running with security.redactClientLogData may be more difficult due to the lack of data related to a log event. See the process logging manual page for an example of the effect of security.redactClientLogData on log output.
	//
	//On a running mongod or mongos, use setParameter with the redactClientLogData parameter to configure this setting.
	RedactClientLogData LogicBoolean `yaml:"redactClientLogData"`

	//New in version 3.6.
	//
	//A list of IP addresses/CIDR (Classless Inter-Domain Routing) ranges against which the mongod validates authentication requests from other members of the replica set and, if part of a sharded cluster, the mongos instances. The mongod verifies that the originating IP is either explicitly in the list or belongs to a CIDR range in the list. If the IP address is not present, the server does not authenticate the mongod or mongos.
	//
	//security.clusterIpSourceWhitelist has no effect on a mongod started without authentication.
	//
	//security.clusterIpSourceWhitelist requires specifying each IPv4/6 address or Classless Inter-Domain Routing (CIDR) range as a YAML list:
	//
	//security:
	//  clusterIpSourceWhitelist:
	//    - 192.0.2.0/24
	//    - 127.0.0.1
	//    - ::1
	//
	//IMPORTANT: Ensure security.clusterIpSourceWhitelist includes the IP address or CIDR ranges that include the IP address of each replica set member or mongos in the deployment to ensure healthy communication between cluster components.
	ClusterIpSourceWhitelist []string `yaml:"clusterIpSourceWhitelist"`

	Sasl Sasl `yaml:"sasl"`

	//Default: false
	//
	//New in version 3.2: Enables encryption for the WiredTiger storage engine. You must set to true to pass in encryption keys and configurations.
	//
	//ENTERPRISE FEATURE
	//
	//Available in MongoDB Enterprise only.
	EnableEncryption LogicBoolean `yaml:"enableEncryption"`

	//Default: AES256-CBC
	//
	//New in version 3.2.
	//
	//The cipher mode to use for encryption at rest:
	//
	//Mode	Description
	//AES256-CBC	256-bit Advanced Encryption Standard in Cipher Block Chaining Mode
	//AES256-GCM
	//256-bit Advanced Encryption Standard in Galois/Counter Mode
	//
	//Available only on Linux.
	//
	//Changed in version 4.0: MongoDB Enterprise on Windows no longer supports AES256-GCM. This cipher is now available only on Linux.
	//
	//ENTERPRISE FEATURE
	//
	//Available in MongoDB Enterprise only.
	EncryptionCipherMode EncryptionCipherMode `yaml:"encryptionCipherMode"`

	//New in version 3.2.
	//
	//The path to the local keyfile when managing keys via process other than KMIP. Only set when managing keys via process other than KMIP. If data is already encrypted using KMIP, MongoDB will throw an error.
	//
	//Requires security.enableEncryption to be true.
	//
	//ENTERPRISE FEATURE
	//
	//Available in MongoDB Enterprise only.
	EncryptionKeyFile string `yaml:"encryptionKeyFile"`
	Kmip              Kmip   `yaml:"kmip"`
	Ldap              Ldap   `yaml:"ldap"`
}

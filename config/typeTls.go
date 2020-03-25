package iotmaker_db_mongodb_config

type Tls struct {
	//New in version 4.2.
	//
	//Enables TLS used for all network connections. The argument to the net.tls.mode setting can be one of the following:
	//
	//Value	Description
	//disabled	The server does not use TLS.
	//allowTLS	Connections between servers do not use TLS. For incoming connections, the server accepts both TLS and non-TLS.
	//preferTLS	Connections between servers use TLS. For incoming connections, the server accepts both TLS and non-TLS.
	//requireTLS	The server uses and accepts only TLS encrypted connections.
	//If --tlsCAFile or tls.CAFile is not specified and you are not using x.509 authentication, the system-wide CA certificate store will be used when connecting to an TLS-enabled server.
	//
	//If using x.509 authentication, --tlsCAFile or tls.CAFile must be specified unless using --tlsCertificateSelector.
	//
	//For more information about TLS and MongoDB, see Configure mongod and mongos for TLS/SSL and TLS/SSL Configuration for Clients .
	Mode Mode `yaml:"mode"`

	//New in version 4.2: The .pem file that contains both the TLS certificate and key.
	//
	//Starting with MongoDB 4.0 on macOS or Windows, you can use the net.tls.certificateSelector setting to specify a certificate from the operating system’s secure certificate store instead of a PEM key file. certificateKeyFile and net.tls.certificateSelector are mutually exclusive. You can only specify one.
	//
	//On Linux/BSD, you must specify net.tls.certificateKeyFile when TLS is enabled.
	//
	//On Windows or macOS, you must specify either net.tls.certificateKeyFile or net.tls.certificateSelector when TLS is enabled.
	//
	//IMPORTANT: For Windows only, MongoDB 4.0 and later do not support encrypted PEM files. The mongod fails to start if it encounters an encrypted PEM file. To securely store and access a certificate for use with TLS on Windows, use net.tls.certificateSelector.
	//
	//For more information about TLS and MongoDB, see Configure mongod and mongos for TLS/SSL and TLS/SSL Configuration for Clients .
	CertificateKeyFile string `yaml:"certificateKeyFile"`

	//New in version 4.2: The password to de-crypt the certificate-key file (i.e. certificateKeyFile). Use the net.tls.certificateKeyPassword option only if the certificate-key file is encrypted. In all cases, the mongos or mongod will redact the password from all logging and reporting output.
	//
	//Starting in MongoDB 4.0:
	//
	//On Linux/BSD, if the private key in the PEM file is encrypted and you do not specify the net.tls.certificateKeyFukePassword option, MongoDB will prompt for a passphrase. See TLS/SSL Certificate Passphrase.
	//On macOS, if the private key in the PEM file is encrypted, you must explicitly specify the net.tls.certificateKeyFilePassword option. Alternatively, you can use a certificate from the secure system store (see net.tls.certificateSelector) instead of a PEM key file or use an unencrypted PEM file.
	//On Windows, MongoDB does not support encrypted certificates. The mongod fails if it encounters an encrypted PEM file. Use net.tls.certificateSelector instead.
	//For more information about TLS and MongoDB, see Configure mongod and mongos for TLS/SSL and TLS/SSL Configuration for Clients .
	CertificateKeyFilePassword string `yaml:"certificateKeyFilePassword"`

	//New in version 4.2: Available on Windows and macOS as an alternative to net.tls.certificateKeyFile.
	//
	//Specifies a certificate property in order to select a matching certificate from the operating system’s certificate store to use for TLS/SSL.
	//
	//net.tls.certificateKeyFile and net.tls.certificateSelector options are mutually exclusive. You can only specify one.
	//
	//net.tls.certificateSelector accepts an argument of the format <property>=<value> where the property can be one of the following:
	//
	//Property	Value type	Description
	//subject	ASCII string	Subject name or common name on certificate
	//thumbprint	hex string
	//A sequence of bytes, expressed as hexadecimal, used to identify a public key by its SHA-1 digest.
	//
	//The thumbprint is sometimes referred to as a fingerprint.
	//
	//When using the system SSL certificate store, OCSP (Online Certificate Status Protocol) is used to validate the revocation status of certificates.
	//
	//The mongod searches the operating system’s secure certificate store for the CA certificates required to validate the full certificate chain of the specified TLS certificate. Specifically, the secure certificate store must contain the root CA and any intermediate CA certificates required to build the full certificate chain to the TLS certificate. Do not use net.tls.CAFile or net.tls.clusterFile to specify the root and intermediate CA certificate
	//
	//For example, if the TLS certificate was signed with a single root CA certificate, the secure certificate store must contain that root CA certificate. If the TLS certificate was signed with an intermediate CA certificate, the secure certificate store must contain the intermedia CA certificate and the root CA certificate.
	//todo:melhorar
	CertificateSelector string `yaml:"certificateSelector"`

	//New in version 4.2: Available on Windows and macOS as an alternative to net.tls.clusterFile.
	//
	//Specifies a certificate property to select a matching certificate from the operating system’s secure certificate store to use for internal x.509 membership authentication.
	//
	//net.tls.clusterFile and net.tls.clusterCertificateSelector options are mutually exclusive. You can only specify one.
	//
	//net.tls.clusterCertificateSelector accepts an argument of the format <property>=<value> where the property can be one of the following:
	//
	//Property	Value type	Description
	//subject	ASCII string	Subject name or common name on certificate
	//thumbprint	hex string
	//A sequence of bytes, expressed as hexadecimal, used to identify a public key by its SHA-1 digest.
	//
	//The thumbprint is sometimes referred to as a fingerprint.
	//
	//The mongod searches the operating system’s secure certificate store for the CA certificates required to validate the full certificate chain of the specified cluster certificate. Specifically, the secure certificate store must contain the root CA and any intermediate CA certificates required to build the full certificate chain to the cluster certificate. Do not use net.tls.CAFile or net.tls.clusterFile to specify the root and intermediate CA certificate.
	//
	//For example, if the cluster certificate was signed with a single root CA certificate, the secure certificate store must contain that root CA certificate. If the cluster certificate was signed with an intermediate CA certificate, the secure certificate store must contain the intermedia CA certificate and the root CA certificate.
	//todo:melhorar
	ClusterCertificateSelector string `yaml:"clusterCertificateSelector"`

	//New in version 4.2: The .pem file that contains the x.509 certificate-key file for membership authentication for the cluster or replica set.
	//
	//Starting with MongoDB 4.0 on macOS or Windows, you can use the net.tls.clusterCertificateSelector option to specify a certificate from the operating system’s secure certificate store instead of a PEM key file. net.tls.clusterFile and net.tls.clusterCertificateSelector options are mutually exclusive. You can only specify one.
	//
	//If net.tls.clusterFile does not specify the .pem file for internal cluster authentication or the alternative net.tls.clusterCertificateSelector, the cluster uses the .pem file specified in the certificateKeyFile setting or the certificate returned by the net.tls.certificateSelector.
	//
	//If using x.509 authentication, --tlsCAFile or tls.CAFile must be specified unless using --tlsCertificateSelector.
	//
	//For more information about TLS and MongoDB, see Configure mongod and mongos for TLS/SSL and TLS/SSL Configuration for Clients .
	//
	//IMPORTANT: For Windows only, MongoDB 4.0 and later do not support encrypted PEM files. The mongod fails to start if it encounters an encrypted PEM file. To securely store and access a certificate for use with membership authentication on Windows, use net.tls.clusterCertificateSelector.
	ClusterFile string `yaml:"clusterFile"`

	//New in version 4.2: The password to de-crypt the x.509 certificate-key file specified with --sslClusterFile. Use the net.tls.clusterPassword option only if the certificate-key file is encrypted. In all cases, the mongos or mongod will redact the password from all logging and reporting output.
	//
	//Starting in MongoDB 4.0:
	//
	//On Linux/BSD, if the private key in the x.509 file is encrypted and you do not specify the net.tls.clusterPassword option, MongoDB will prompt for a passphrase. See TLS/SSL Certificate Passphrase.
	//On macOS, if the private key in the x.509 file is encrypted, you must explicitly specify the net.tls.clusterPassword option. Alternatively, you can either use a certificate from the secure system store (see net.tls.clusterCertificateSelector) instead of a cluster PEM file or use an unencrypted PEM file.
	//On Windows, MongoDB does not support encrypted certificates. The mongod fails if it encounters an encrypted PEM file. Use net.tls.clusterCertificateSelector.
	//For more information about TLS and MongoDB, see Configure mongod and mongos for TLS/SSL and TLS/SSL Configuration for Clients .
	ClusterPassword string `yaml:"clusterPassword"`

	//New in version 4.2: The .pem file that contains the root certificate chain from the Certificate Authority. Specify the file name of the .pem file using relative or absolute paths.
	//
	//Windows/macOS Only: If using net.tls.certificateSelector and/or net.tls.clusterCertificateSelector, do not use net.tls.CAFile to specify the root and intermediate CA certificates. Store all CA certificates required to validate the full trust chain of the net.tls.certificateSelector and/or net.tls.clusterCertificateSelector certificates in the secure certificate store.
	//
	//For more information about TLS and MongoDB, see Configure mongod and mongos for TLS/SSL and TLS/SSL Configuration for Clients .
	CAFile string `yaml:"CAFile"`

	//New in version 4.2: The .pem file that contains the root certificate chain from the Certificate Authority used to validate the certificate presented by a client establishing a connection. Specify the file name of the .pem file using relative or absolute paths. net.tls.clusterCAFile requires that net.tls.CAFile is set.
	//
	//If net.tls.clusterCAFile does not specify the .pem file for validating the certificate from a client establishing a connection, the cluster uses the .pem file specified in the net.tls.CAFile option.
	//
	//net.tls.clusterCAFile lets you use separate Certificate Authorities to verify the client to server and server to client portions of the TLS handshake.
	//
	//Starting in 4.0, on macOS or Windows, you can use a certificate from the operating system’s secure store instead of a PEM key file. See net.tls.clusterCertificateSelector. When using the secure store, you do not need to, but can, also specify the net.tls.clusterCAFile.
	//
	//Windows/macOS Only: If using net.tls.certificateSelector and/or net.tls.clusterCertificateSelector, do not use net.tls.clusterCAFile to specify the root and intermediate CA certificates. Store all CA certificates required to validate the full trust chain of the net.tls.certificateSelector and/or net.tls.clusterCertificateSelector certificates in the secure certificate store.
	//
	//For more information about TLS and MongoDB, see Configure mongod and mongos for TLS/SSL and TLS/SSL Configuration for Clients .
	ClusterCAFile string `yaml:"clusterCAFile"`

	//New in version 4.2.
	//
	//The the .pem file that contains the Certificate Revocation List. Specify the file name of the .pem file using relative or absolute paths.
	//
	//NOTE: Starting in MongoDB 4.0, you cannot specify net.tls.CRLFile on macOS. Use net.tls.certificateSelector instead.
	//
	//For more information about TLS and MongoDB, see Configure mongod and mongos for TLS/SSL and TLS/SSL Configuration for Clients .
	CRLFile string `yaml:"CRLFile"`

	//New in version 4.2.
	//
	//For clients that do not present certificates, mongos or mongod bypasses TLS/SSL certificate validation when establishing the connection.
	//
	//For clients that present a certificate, however, mongos or mongod performs certificate validation using the root certificate chain specified by CAFile and reject clients with invalid certificates.
	//
	//Use the net.tls.allowConnectionsWithoutCertificates option if you have a mixed deployment that includes clients that do not or cannot present certificates to the mongos or mongod.
	//
	//For more information about TLS and MongoDB, see Configure mongod and mongos for TLS/SSL and TLS/SSL Configuration for Clients .
	AllowConnectionsWithoutCertificates LogicBoolean `yaml:"allowConnectionsWithoutCertificates"`

	//New in version 4.2.
	//
	//Enable or disable the validation checks for TLS certificates on other servers in the cluster and allows the use of invalid certificates to connect.
	//
	//NOTE
	//
	//If you specify --tlsAllowInvalidCertificates or tls.allowInvalidCertificates: true when using x.509 authentication, an invalid certificate is only sufficient to establish a TLS connection but is insufficient for authentication.
	//
	//When using the net.tls.allowInvalidCertificates setting, MongoDB logs a warning regarding the use of the invalid certificate.
	//
	//For more information about TLS and MongoDB, see Configure mongod and mongos for TLS/SSL and TLS/SSL Configuration for Clients .
	AllowInvalidCertificates LogicBoolean `yaml:"allowInvalidCertificates"`

	//Default: false
	//
	//When net.tls.allowInvalidHostnames is true, MongoDB disables the validation of the hostnames in TLS certificates, allowing mongod to connect to MongoDB instances if the hostname their certificates do not match the specified hostname.
	//
	//For more information about TLS and MongoDB, see Configure mongod and mongos for TLS/SSL and TLS/SSL Configuration for Clients .
	AllowInvalidHostNames LogicBoolean `yaml:"allowInvalidHostnames"`

	//New in version 4.2.
	//
	//Prevents a MongoDB server running with TLS from accepting incoming connections that use a specific protocol or protocols. To specify multiple protocols, use a comma separated list of protocols.
	//
	//net.tls.disabledProtocols recognizes the following protocols: TLS1_0, TLS1_1, TLS1_2, and starting in version 4.0.4 (and 3.6.9), TLS1_3.
	//
	//On macOS, you cannot disable TLS1_1 and leave both TLS1_0 and TLS1_2 enabled. You must disable at least one of the other two, for example, TLS1_0,TLS1_1.
	//To list multiple protocols, specify as a comma separated list of protocols. For example TLS1_0,TLS1_1.
	//Specifying an unrecognized protocol will prevent the server from starting.
	//The specified disabled protocols overrides any default disabled protocols.
	//Starting in version 4.0, MongoDB disables the use of TLS 1.0 if TLS 1.1+ is available on the system. To enable the disabled TLS 1.0, specify none to net.tls.disabledProtocols. See Disable TLS 1.0.
	//
	//Members of replica sets and sharded clusters must speak at least one protocol in common.
	//
	//SEE ALSO: Disallow Protocols https://docs.mongodb.com/manual/tutorial/configure-ssl/#ssl-disallow-protocols
	DisabledProtocols []DisabledProtocol `yaml:"disabledProtocols"`

	//New in version 4.2.
	//
	//Enable or disable the use of the FIPS mode of the TLS library for the mongos or mongod. Your system must have a FIPS compliant library to use the net.tls.FIPSMode option.
	//
	//NOTE: FIPS-compatible TLS/SSL is available only in MongoDB Enterprise. See Configure MongoDB for FIPS for more information.
	FIPSMode LogicBoolean `yaml:"FIPSMode"`
}

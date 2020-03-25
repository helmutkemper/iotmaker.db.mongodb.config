package iotmaker_db_mongodb_config

type Kmip struct {
	//New in version 3.2.
	//
	//Unique KMIP identifier for an existing key within the KMIP server. Include to use the
	//key associated with the identifier as the system key.
	//You can only use the setting the first time you enable encryption for the mongod
	//instance. Requires security.enableEncryption to be true.
	//
	//If unspecified, MongoDB will request that the KMIP server create a new key to utilize
	//as the system key.
	//
	//If the KMIP server cannot locate a key with the specified identifier or the data is
	//already encrypted with a key, MongoDB will throw an error.
	//
	//ENTERPRISE FEATURE
	//
	//Available in MongoDB Enterprise only.
	KeyIdentifier string `yaml:"keyIdentifier"`

	//Default: false
	//
	//New in version 3.2.
	//
	//If true, rotate the master key and re-encrypt the internal keystore.
	//
	//ENTERPRISE FEATURE
	//
	//Available in MongoDB Enterprise only.
	//
	//SEE ALSO: KMIP Master Key Rotation
	//https://docs.mongodb.com/manual/tutorial/rotate-encryption-key/#kmip-master-key-rotation
	RotateMasterKey LogicBoolean `yaml:"rotateMasterKey"`

	//New in version 3.2.
	//
	//Hostname or IP address of key management solution running a KMIP server.
	//Requires security.enableEncryption to be true.
	//
	//ENTERPRISE FEATURE: Available in MongoDB Enterprise only.
	ServerName string `yaml:"serverName"`

	//Default: 5696
	//
	//New in version 3.2.
	//
	//Port number the KMIP server is listening on. Requires that a security.kmip.serverName
	//be provided. Requires security.enableEncryption to be true.
	//
	//ENTERPRISE FEATURE: Available in MongoDB Enterprise only.
	Port int `yaml:"port"`

	//New in version 3.2.
	//
	//String containing the path to the client certificate used for authenticating MongoDB
	//to the KMIP server. Requires that a security.kmip.serverName be provided.
	//
	//NOTE
	//
	//Starting in 4.0, on macOS or Windows, you can use a certificate from the operating
	//system’s secure store instead of a PEM key file.
	//See security.kmip.clientCertificateSelector.
	//
	//ENTERPRISE FEATURE: Available in MongoDB Enterprise only.
	ClientCertificateFile string `yaml:"clientCertificateFile"`

	//New in version 3.2.
	//
	//The password to decrypt the client certificate
	//(i.e. security.kmip.clientCertificateFile), used to authenticate MongoDB to the KMIP
	//server.
	//Use the option only if the certificate is encrypted.
	//
	//ENTERPRISE FEATURE: Available in MongoDB Enterprise only.
	ClientCertificatePassword string `yaml:"clientCertificatePassword"`

	//New in version 4.0: Available on Windows and macOS as an alternative to
	//security.kmip.clientCertificateFile.
	//
	//security.kmip.clientCertificateFile and security.kmip.clientCertificateSelector
	//options are mutually exclusive. You can only specify one.
	//
	//Specifies a certificate property in order to select a matching certificate from the
	//operating system’s certificate store to authenticate MongoDB to the KMIP server.
	//
	//security.kmip.clientCertificateSelector accepts an argument of the format
	//<property>=<value> where the property can be one of the following:
	//
	//Property	Value type	Description
	//subject	ASCII string	Subject name or common name on certificate
	//thumbprint	hex string
	//A sequence of bytes, expressed as hexadecimal, used to identify a public key by its
	//SHA-1 digest.
	//
	//The thumbprint is sometimes referred to as a fingerprint.
	//
	//ENTERPRISE FEATURE: Available in MongoDB Enterprise only.
	ClientCertificateSelector string `yaml:"clientCertificateSelector"`

	//New in version 3.2.
	//
	//Path to CA File. Used for validating secure client connection to KMIP server.
	//
	//NOTE:
	//    Starting in 4.0, on macOS or Windows, you can use a certificate from the
	//    operating system’s secure store instead of a PEM key file.
	//   See security.kmip.clientCertificateSelector. When using the secure store, you do
	//  not need to, but can, also specify the security.kmip.serverCAFile.
	//
	//ENTERPRISE FEATURE: Available in MongoDB Enterprise only.
	ServerCAFile string `yaml:"serverCAFile"`
}

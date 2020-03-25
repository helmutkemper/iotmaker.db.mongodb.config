package iotmaker_db_mongodb_config

type Bind struct {
	//New in version 3.4: Available in MongoDB Enterprise only.
	//
	//The identity with which mongod or mongos binds as, when connecting to or performing
	//queries on an LDAP server.
	//
	//Only required if any of the following are true:
	//
	//Using LDAP authorization.
	//Using an LDAP query for security.ldap.userToDNMapping.
	//The LDAP server disallows anonymous binds
	//You must use queryUser with queryPassword.
	//
	//If unset, mongod or mongos will not attempt to bind to the LDAP server.
	//
	//This setting can be configured on a running mongod or mongos using setParameter.
	//
	//NOTE: Windows MongoDB deployments can use bindWithOSDefaults instead of queryUser and
	//queryPassword. You cannot specify both queryUser and bindWithOSDefaults at the same
	//time.
	QueryUser string `yaml:"queryUser"`

	//New in version 3.4: Available in MongoDB Enterprise only.
	//
	//The password used to bind to an LDAP server when using queryUser. You must use
	//queryPassword with queryUser.
	//
	//If unset, mongod or mongos will not attempt to bind to the LDAP server.
	//
	//This setting can be configured on a running mongod or mongos using setParameter.
	//
	//NOTE: Windows MongoDB deployments can use bindWithOSDefaults instead of queryPassword
	//and queryPassword. You cannot specify both queryPassword and bindWithOSDefaults at
	//the same time.
	QueryPassword string `yaml:"queryPassword"`

	//Default: false
	//
	//New in version 3.4: Available in MongoDB Enterprise for the Windows platform only.
	//
	//Allows mongod or mongos to authenticate, or bind, using your Windows login
	//credentials when connecting to the LDAP server.
	//
	//Only required if:
	//
	//Using LDAP authorization.
	//Using an LDAP query for username transformation.
	//The LDAP server disallows anonymous binds
	//Use useOSDefaults to replace queryUser and queryPassword.
	UseOSDefaults LogicBoolean `yaml:"useOSDefaults"`

	//Default: simple
	//
	//New in version 3.4: Available in MongoDB Enterprise only.
	//
	//The method mongod or mongos uses to authenticate to an LDAP server. Use with
	//queryUser and queryPassword to connect to the LDAP server.
	//
	//method supports the following values:
	//
	//    Value  |  Description
	//    ---------------------------------------------------------------------------------
	//    simple |  mongod or mongos uses simple authentication.
	//    ---------------------------------------------------------------------------------
	//    sasl   |  mongod or mongos uses SASL protocol for authentication
	//If you specify sasl, you can configure the available SASL mechanisms using
	//security.ldap.bind.saslMechanisms. mongod or mongos defaults to using DIGEST-MD5
	//mechanism.
	Method Method `yaml:"method"`

	//Default: DIGEST-MD5
	//
	//New in version 3.4: Available in MongoDB Enterprise only.
	//
	//A comma-separated list of SASL mechanisms mongod or mongos can use when
	//authenticating to the LDAP server. The mongod or mongos and the LDAP server must
	//agree on at least one mechanism. The mongod or mongos dynamically loads any SASL
	//mechanism libraries installed on the host machine at runtime.
	//
	//Install and configure the appropriate libraries for the selected SASL mechanism(s) on
	//both the mongod or mongos host and the remote LDAP server host. Your operating system
	//may include certain SASL libraries by default. Defer to the documentation associated
	//with each SASL mechanism for guidance on installation and configuration.
	//
	//If using the GSSAPI SASL mechanism for use with Kerberos Authentication, verify the
	//following for the mongod or mongos host machine:
	//
	//Linux:
	//    The KRB5_CLIENT_KTNAME environment variable resolves to the name of the client
	//    Linux Keytab Files for the host machine.
	//    For more on Kerberos environment variables, please defer to the Kerberos
	//    documentation.
	//
	//    The client keytab includes a User Principal for the mongod or mongos to use when
	//    connecting to the LDAP server and execute LDAP queries.
	//
	//
	//Windows:
	//    If connecting to an Active Directory server, the Windows Kerberos configuration
	//    automatically generates a Ticket-Granting-Ticket when the user logs onto the
	//    system.
	//    Set useOSDefaults to true to allow mongod or mongos to use the generated
	//    credentials when connecting to the Active Directory server and execute queries.
	//
	//    Set method to sasl to use this option.
	//
	//NOTE:
	//    For a complete list of SASL mechanisms see the IANA listing.
	//    Defer to the documentation for your LDAP or Active Directory service for
	//    identifying the SASL mechanisms compatible with the service.
	//
	//
	//MongoDB is not a source of SASL mechanism libraries, nor is the MongoDB documentation
	//a definitive source for installing or configuring any given SASL mechanism.
	//For documentation and support, defer to the SASL mechanism library vendor or owner.
	//
	//For more information on SASL, defer to the following resources:
	//
	//For Linux, please see the Cyrus SASL documentation.
	//
	//For Windows, please see the Windows SASL documentation.
	SaslMechanisms string `yaml:"saslMechanisms"`
}

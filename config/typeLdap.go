package iotmaker_db_mongodb_config

type Ldap struct {
	//New in version 3.4: Available in MongoDB Enterprise only.
	//
	//The LDAP server against which the mongod or mongos authenticates users or determines
	//what actions a user is authorized to perform on a given database.
	//If the LDAP server specified has any replicated instances, you may specify the host
	//and port of each replicated server in a comma-delimited list.
	//
	//If your LDAP infrastructure partitions the LDAP directory over multiple LDAP servers,
	//specify one LDAP server or any of its replicated instances to security.ldap.servers.
	//MongoDB supports following LDAP referrals as defined in RFC 4511 4.1.10. Do not use
	//security.ldap.servers for listing every LDAP server in your infrastructure.
	//
	//This setting can be configured on a running mongod or mongos using setParameter.
	//
	//If unset, mongod or mongos cannot use LDAP authentication or authorization.
	Servers string `yaml:"servers"`
	Bind    Bind   `yaml:"bind"`

	//Default: tls
	//
	//New in version 3.4: Available in MongoDB Enterprise only.
	//
	//By default, mongod or mongos creates a TLS/SSL secured connection to the LDAP server.
	//
	//For Linux deployments, you must configure the appropriate TLS Options in
	///etc/openldap/ldap.conf file.
	//Your operating system’s package manager creates this file as part of the MongoDB
	//Enterprise installation, via the libldap dependency.
	//See the documentation for TLS Options in the ldap.conf OpenLDAP documentation for
	//more complete instructions.
	//
	//For Windows deployment, you must add the LDAP server CA certificates to the Windows
	//certificate management tool.
	//The exact name and functionality of the tool may vary depending on operating system
	//version.
	//Please see the documentation for your version of Windows for more information on
	//certificate management.
	//
	//Set transportSecurity to none to disable TLS/SSL between mongod or mongos and the
	//LDAP server.
	//
	//WARNING:
	//    Setting transportSecurity to none transmits plaintext information and possibly
	//    credentials between mongod or mongos and the LDAP server.
	TransportSecurity string `yaml:"transportSecurity"`

	//Default: 10000
	//
	//New in version 3.4: Available in MongoDB Enterprise only.
	//
	//The amount of time in milliseconds mongod or mongos should wait for an LDAP server to
	//respond to a request.
	//
	//Increasing the value of timeoutMS may prevent connection failure between the MongoDB
	//server and the LDAP server, if the source of the failure is a connection timeout.
	//Decreasing the value of timeoutMS reduces the time MongoDB waits for a response from
	//the LDAP server.
	//
	//This setting can be configured on a running mongod or mongos using setParameter.
	TimeoutMS int `yaml:"timeoutMS"`

	//New in version 3.4: Available in MongoDB Enterprise only.
	//
	//Maps the username provided to mongod or mongos for authentication to a LDAP
	//Distinguished Name (DN).
	//You may need to use userToDNMapping to transform a username into an LDAP DN in the
	//following scenarios:
	//
	//Performing LDAP authentication with simple LDAP binding, where users authenticate to
	//MongoDB with usernames that are not full LDAP DNs.
	//
	//Using an LDAP authorization query template that requires a DN.
	//Transforming the usernames of clients authenticating to Mongo DB using different
	//authentication mechanisms (e.g. x.509, kerberos) to a full LDAP DN for authorization.
	//
	//userToDNMapping expects a quote-enclosed JSON-string representing an ordered array of
	//documents.
	//Each document contains a regular expression match and either a substitution or
	//ldapQuery template used for transforming the incoming username.
	//
	//Each document in the array has the following form:
	//
	//    {
	//        match: "<regex>"
	//        substitution: "<LDAP DN>" | ldapQuery: "<LDAP Query>"
	//    }
	//
	//Field	Description	Example
	//match	An ECMAScript-formatted regular expression (regex) to match against a provided
	//username.
	//Each parenthesis-enclosed section represents a regex capture group used by
	//substitution or ldapQuery.	"(.+)ENGINEERING" "(.+)DBA"
	//substitution
	//An LDAP distinguished name (DN) formatting template that converts the authentication
	//name matched by the match regex into a LDAP DN. Each curly bracket-enclosed numeric
	//value is replaced by the corresponding regex capture group extracted from the
	//authentication username via the match regex.
	//
	//The result of the substitution must be an RFC4514 escaped string.
	//
	//"cn={0},ou=engineering, dc=example,dc=com"
	//ldapQuery	A LDAP query formatting template that inserts the authentication name
	//matched by the match regex into an LDAP query URI encoded respecting RFC4515 and
	//RFC4516.
	//
	//Each curly bracket-enclosed numeric value is replaced by the corresponding regex
	//capture group extracted from the authentication username via the match expression.
	//mongod or mongos executes the query against the LDAP server to retrieve the LDAP DN
	//for the authenticated user.
	//mongod or mongos requires exactly one returned result for the transformation to be
	//successful, or mongod or mongos skips this transformation.
	//    "ou=engineering,dc=example, dc=com??one?(user={0})"
	//
	//NOTE:
	//    An explanation of RFC4514, RFC4515, RFC4516, or LDAP queries is out of scope for
	//    the MongoDB Documentation.
	//    Please review the RFC directly or use your preferred LDAP resource.
	//
	//For each document in the array, you must use either substitution or ldapQuery.
	//You cannot specify both in the same document.
	//
	//When performing authentication or authorization, mongod or mongos steps through each
	//document in the array in the given order, checking the authentication username
	//against the match filter.
	//If a match is found, mongod or mongos applies the transformation and uses the output
	//for authenticating the user.
	//mongod or mongos does not check the remaining documents in the array.
	//
	//If the given document does not match the provided authentication name, or the
	//transformation described by the document fails, mongod or mongos continues through
	//the list of documents to find additional matches.
	//If no matches are found in any document, mongod or mongos returns an error.
	//
	//EXAMPLE:
	//    The following shows two transformation documents.
	//    The first document matches against any string ending in @ENGINEERING, placing
	//    anything preceeding the suffix into a regex capture group.
	//
	//    The second document matches against any string ending in @DBA, placing anything
	//    preceeding the suffix into a regex capture group.
	//
	//IMPORTANT:
	//    You must pass the array to userToDNMapping as a string.
	//
	//        "[
	//           {
	//              match: "(.+)@ENGINEERING.EXAMPLE.COM",
	//              substitution: "cn={0},ou=engineering,dc=example,dc=com"
	//           },
	//           {
	//              match: "(.+)@DBA.EXAMPLE.COM",
	//              ldapQuery: "ou=dba,dc=example,dc=com??one?(user={0})"
	//           }
	//        ]"
	//
	//
	//A user with username alice@ENGINEERING.EXAMPLE.COM matches the first document.
	//The regex capture group {0} corresponds to the string alice.
	//The resulting output is the DN "cn=alice,ou=engineering,dc=example,dc=com".
	//
	//A user with username bob@DBA.EXAMPLE.COM matches the second document.
	//The regex capture group {0} corresponds to the string bob.
	//The resulting output is the LDAP query "ou=dba,dc=example,dc=com??one?(user=bob)".
	//mongod or mongos executes this query against the LDAP server, returning the result
	//"cn=bob,ou=dba,dc=example,dc=com".
	//
	//If userToDNMapping is unset, mongod or mongos applies no transformations to the
	//username when attempting to authenticate or authorize a user against the LDAP server.
	//
	//This setting can be configured on a running mongod or mongos using the setParameter
	//database command.
	UserToDNMapping string `yaml:"userToDNMapping"`
	Authz           Authz  `yaml:"authz"`

	//Default: true
	//
	//Available in MongoDB Enterprise
	//
	//A flag that determines if the mongod or mongos instance checks the availability of
	//the LDAP server(s) as part of its startup:
	//
	//If true, the mongod or mongos instance performs the availability check and only
	//continues to start up if the LDAP server is available.
	//If false, the mongod or mongos instance skips the availability check; i.e. the
	//instance starts up even if the LDAP server is unavailable.
	ValidateLDAPServerConfig LogicBoolean `yaml:"validateLDAPServerConfig"`
}

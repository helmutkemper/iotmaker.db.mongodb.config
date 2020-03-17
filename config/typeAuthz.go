package iotmaker_db_mongodb_config

type Authz struct {
	//New in version 3.4: Available in MongoDB Enterprise only.
	//
	//A relative LDAP query URL formatted conforming to RFC4515 and RFC4516 that mongod executes to obtain the LDAP groups to which the authenticated user belongs to. The query is relative to the host or hosts specified in security.ldap.servers.
	//
	//In the URL, you can use the following substituion tokens:
	//
	//Substitution Token	Description
	//{USER}	Substitutes the authenticated username, or the transformed username if a userToDNMapping is specified.
	//{PROVIDED_USER}
	//Substitutes the supplied username, i.e. before either authentication or LDAP transformation.
	//
	//New in version 4.2.
	//
	//When constructing the query URL, ensure that the order of LDAP parameters respects RFC4516:
	//
	//[ dn  [ ? [attributes] [ ? [scope] [ ? [filter] [ ? [Extensions] ] ] ] ] ]
	//If your query includes an attribute, mongod assumes that the query retrieves a list of the DNs which this entity is a member of.
	//
	//If your query does not include an attribute, mongod assumes the query retrieves all entities which the user is member of.
	//
	//For each LDAP DN returned by the query, mongod assigns the authorized user a corresponding role on the admin database. If a role on the on the admin database exactly matches the DN, mongod grants the user the roles and privileges assigned to that role. See the db.createRole() method for more information on creating roles.
	//
	//EXAMPLE
	//
	//This LDAP query returns any groups listed in the LDAP user object’s memberOf attribute.
	//
	//"{USER}?memberOf?base"
	//Your LDAP configuration may not include the memberOf attribute as part of the user schema, may possess a different attribute for reporting group membership, or may not track group membership through attributes. Configure your query with respect to your own unique LDAP configuration.
	//
	//If unset, mongod cannot authorize users using LDAP.
	//
	//This setting can be configured on a running mongod using the setParameter database command.
	//
	//NOTE
	//
	//An explanation of RFC4515, RFC4516 or LDAP queries is out of scope for the MongoDB Documentation. Please review the RFC directly or use your preferred LDAP resource.
	QueryTemplate string `yaml:"-"`

	//Default: true
	//
	//Available in MongoDB Enterprise
	//
	//A flag that determines if the mongod or mongos instance checks the availability of the LDAP server(s) as part of its startup:
	//
	//If true, the mongod or mongos instance performs the availability check and only continues to start up if the LDAP server is available.
	//If false, the mongod or mongos instance skips the availability check; i.e. the instance starts up even if the LDAP server is unavailable.
	ValidateLDAPServerConfig LogicBoolean `yaml:"-"`
}

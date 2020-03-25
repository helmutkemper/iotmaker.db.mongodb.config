package iotmaker_db_mongodb_config

type Authorization int

var authorizations = [...]string{
	"",
	"enabled",
	"disabled",
}

func (el Authorization) String() string {
	return `"` + authorizations[el] + `"`
}

const (
	//A user can access only the database resources and actions for which they have been
	//granted privileges.
	KAuthorizationEnabled Authorization = iota + 1

	//A user can access any database and perform any action.
	KAuthorizationDisabled
)

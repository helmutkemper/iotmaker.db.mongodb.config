package iotmaker_db_mongodb_config

type Method int

var methods = [...]string{
	"",
	"simple",
	"AES256-GCM",
}

func (el Method) String() string {
	return `"` + methods[el] + `"`
}

const (
	//mongod or mongos uses simple authentication.
	KBindMethodSimple Method = iota + 1

	//mongod or mongos uses SASL protocol for authentication
	KBindMethodSasl
)

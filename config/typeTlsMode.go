package iotmaker_db_mongodb_config

type Mode int

var modes = [...]string{
	"",
	"disabled",
	"allowTLS",
	"preferTLS",
	"requireTLS",
}

func (el Mode) String() string {
	return `"` + modes[el] + `"`
}

const (
	//The server does not use TLS.
	KModeDisabled Mode = iota + 1

	//Connections between servers do not use TLS. For incoming connections, the server
	//accepts both TLS and non-TLS.
	KModeAllowTLS

	//Connections between servers use TLS. For incoming connections, the server accepts
	//both TLS and non-TLS.
	KModePreferTLS

	//The server uses and accepts only TLS encrypted connections.
	KModeRequireTLS
)

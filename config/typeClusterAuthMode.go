package iotmaker_db_mongodb_config

type ClusterAuthMode int

var clusterAuthModes = [...]string{
	"",
	"keyFile",
	"sendKeyFile",
	"sendX509",
	"x509",
}

func (el ClusterAuthMode) String() string {
	return `"` + clusterAuthModes[el] + `"`
}

const (
	//Use a keyfile for authentication. Accept only keyfiles.
	KClusterAuthModeKeyFile ClusterAuthMode = iota + 1

	//For rolling upgrade purposes. Send a keyfile for authentication but can accept both
	//keyfiles and x.509 certificates.
	KClusterAuthModeSendKeyFile

	//For rolling upgrade purposes. Send the x.509 certificate for authentication but can
	//accept both keyfiles and x.509 certificates.
	KClusterAuthModeSendX509

	//Recommended. Send the x.509 certificate for authentication and accept only x.509
	//certificates.
	KClusterAuthModeX509
)

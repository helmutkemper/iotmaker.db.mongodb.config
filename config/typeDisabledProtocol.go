package iotmaker_db_mongodb_config

type DisabledProtocol int

var disabledProtocols = [...]string{
	"",
	"TLS1_0",
	"TLS1_1",
	"TLS1_2",
	"TLS1_3",
}

func (el DisabledProtocol) String() string {
	return `"` + disabledProtocols[el] + `"`
}

const (
	KDisabledProtocolTLS1v0 DisabledProtocol = iota + 1
	KDisabledProtocolTLS1v1
	KDisabledProtocolTLS1v2
	KDisabledProtocolTLS1v3
)

package iotmaker_db_mongodb_config

type Component struct {
	AccessControl AccessControlLog `yaml:"accessControl"` //ok
	Command       CommandLog       `yaml:"command"`       //ok
	Control       ControlLog       `yaml:"control"`       //ok
	Ftdc          FtdcLog          `yaml:"ftdc"`          //ok
	Geo           GeoLog           `yaml:"geo"`           //ok
	Index         IndexLog         `yaml:"index"`         //ok
	Network       NetworkLog       `yaml:"Network"`       //ok
	Query         QueryLog         `yaml:"query"`         //ok
	Replication   ReplicationLog   `yaml:"replication"`   //ok
	Transaction   TransactionLog   `yaml:"transaction"`   //ok
	Write         WriteLog         `yaml:"write"`         //ok
}

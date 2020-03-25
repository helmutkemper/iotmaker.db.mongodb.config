package iotmaker_db_mongodb_config

type Component struct {
	AccessControl AccessControlLog `yaml:"accessControl"`
	Command       CommandLog       `yaml:"command"`
	Control       ControlLog       `yaml:"control"`
	Ftdc          FtdcLog          `yaml:"ftdc"`
	Geo           GeoLog           `yaml:"geo"`
	Index         IndexLog         `yaml:"index"`
	Network       NetworkLog       `yaml:"Network"`
	Query         QueryLog         `yaml:"query"`
	Replication   ReplicationLog   `yaml:"replication"`
	Transaction   TransactionLog   `yaml:"transaction"`
	Write         WriteLog         `yaml:"write"`
}

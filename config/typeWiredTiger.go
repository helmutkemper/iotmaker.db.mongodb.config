package iotmaker_db_mongodb_config

type WiredTiger struct {
	EngineConfig WiredTigerEngineConfig `yaml:"-"`
	IndexConfig  IndexConfig            `yaml:"-"`
}

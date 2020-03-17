package iotmaker_db_mongodb_config

type WiredTiger struct {
	EngineConfig     WiredTigerEngineConfig `yaml:"engineConfig"`
	IndexConfig      IndexConfig            `yaml:"indexConfig"`
	CollectionConfig CollectionConfig       `yaml:"collectionConfig"`
}

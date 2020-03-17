package iotmaker_db_mongodb_config

type BlockCompressor int

var blockCompressors = [...]string{
	"",
	"none",
	"snappy",
	"zlib",
	"zstd",
}

func (el BlockCompressor) String() string {
	return `"` + blockCompressors[el] + `"`
}

const (
	KBlockCompressorNone BlockCompressor = iota + 1
	KBlockCompressorSnappy
	KBlockCompressorZLib
	KBlockCompressorZStd
)

package iotmaker_db_mongodb_config

type Compressor int

var compressors = [...]string{
	"",
	"snappy",
	"zlib",
	"zstd",
}

func (el Compressor) String() string {
	return `"` + compressors[el] + `"`
}

const (
	KCompressorSnappy Compressor = iota + 1
	KCompressorZLib
	KCompressorZStd
)

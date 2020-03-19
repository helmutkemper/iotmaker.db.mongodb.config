package iotmaker_db_mongodb_config

type EncryptionCipherMode int

var encryptionCipherModes = [...]string{
	"",
	"AES256-CBC",
	"AES256-GCM",
}

func (el EncryptionCipherMode) String() string {
	return `"` + encryptionCipherModes[el] + `"`
}

const (
	//256-bit Advanced Encryption Standard in Cipher Block Chaining Mode
	KEncryptionCipherModeAES256CBC EncryptionCipherMode = iota + 1

	//256-bit Advanced Encryption Standard in Galois/Counter Mode
	//
	//Available only on Linux.
	//
	//Changed in version 4.0: MongoDB Enterprise on Windows no longer supports AES256-GCM.
	//This cipher is now available only on Linux.
	KEncryptionCipherModeAES256GCM
)

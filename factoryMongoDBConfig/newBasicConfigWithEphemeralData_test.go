package factoryMongoDBConfig

import (
	"fmt"
)

func ExampleNewBasicConfig() {
	var err error
	var file []byte
	var c = NewBasicConfigWithEphemeralData()
	err, file = c.ToYaml(0)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", file)

	// Output:
	// processManagement:
	//   fork: false
	// net:
	//   port: 27017
	//   bindIp:
	//     - "127.0.0.1"
	//     - "0.0.0.0"
	// setParameter:
	//   enableLocalhostAuthBypass: true
	// storage:
	//   journal:
	//     enabled: true
	// systemLog:
	//   path: "/var/log/mongodb/mongod.log"
	//   logAppend: true
	//   destination: "file"
}

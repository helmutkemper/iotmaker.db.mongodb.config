package iotmaker_db_mongodb_config

import "fmt"

func ExampleConfiguration_ToYaml() {
	conf := &Configuration{}

	err, yml := conf.ToYaml(0)
	if err != nil {
		fmt.Printf("error: %v\n", err.Error())
		return
	}

	fmt.Printf("%s\n", yml)
	//output:
	//
}

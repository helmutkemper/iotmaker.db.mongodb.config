# iotmaker.db.mongodb.config

pt_BR: Este módulo está em versão alpha e tem por finalidade gerar arquivos **YAML** de
configuração do **MongoDB**.

en: This module is in alpha version and writes files **YAML** to configurate MongoDB.

```golang
    var err error
	var file []byte
	var c = factoryMongoDBConfig.NewBasicConfigWithEphemeralData()
	err, file = c.ToYaml(0)
	if err != nil {
		panic(err)
	}
    err = ioutil.WriteFile("./mongod.conf", file, 0644) 
``` 

file: **./mongod.conf**
```yaml
processManagement:
  fork: false
net:
  port: 27017
  bindIp:
    - "127.0.0.1"
    - "0.0.0.0"
setParameter:
  enableLocalhostAuthBypass: true
storage:
  journal:
    enabled: true
systemLog:
  path: "/var/log/mongodb/mongod.log"
  logAppend: true
  destination: "file"
```
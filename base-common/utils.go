package base_common

import (
	"encoding/json"
	"log"
	"os"
)

const configPath = "../etc/config.json"

type AppConfig struct {
	Server      string
	MongoDBHost string
	DBUser      string
	DBPwd       string
	Database    string
}
var AppConf *AppConfig

func initConfig() {
	configLoad()
}

func configLoad() {
	file, err := os.Open(configPath)
	defer file.Close()
	if err != nil {
		log.Fatalf("[loadConfig]: %s\n", err)
	}
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&AppConf)
	if err != nil {
		log.Fatalf("[loadConfig]: %s\n", err)
	}
}

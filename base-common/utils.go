package base_common

import (
	"encoding/json"
	"log"
	"os"
)

type dbConfiguration struct {
	Server, MongoDBHost, DBUser, DBPwd, Database string
}

func initConfig() {
	loadConfig()
}

func loadConfig() {
	file, err := os.Open("../etc/config.json")
	defer file.Close()
	if err != nil {
		log.Fatalf("[loadConfig]: %s\n", err)
	}
	decoder := json.NewDecoder(file)
	dbConfig := dbConfiguration{}
	err = decoder.Decode(&dbConfig)
	if err != nil {
		log.Fatalf("[loadConfig]: %s\n", err)
	}
}

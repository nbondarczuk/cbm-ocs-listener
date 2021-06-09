package common

import (
	"encoding/json"
	"log"
	"os"
)

type (
	Configuration struct {
		RunPath,
		Debug,
		OracleDBUser,
		OracleDBPassword,
		OracleServiceName,
		QueueServerUrl,
		QueueName,
		QueueUser,
		QueuePassword,
		CbmServerUrl,
		SystemResponderUrl,
		Send,
		FileSend,
		Parse,
		FileParse string
	}
)

// AppConfig holds the configuration values from config.json file
var AppConfig Configuration
var TestRun bool = false

// Initialize AppConfig
func InitConfig(config string) {
	loadAppConfig(config)
}

//
// Reads config.json and decode into AppConfig
//
func loadAppConfig(config string) {
	// config file name may be overriden by explicit invocation parameter
	log.Printf("Using config file: %s", config)

	// file must exist
	file, err := os.Open(config)
	defer file.Close()
	if err != nil {
		log.Fatalf("Can not open file: %s\n", err)
	}

	// decode json contents
	decoder := json.NewDecoder(file)
	AppConfig = Configuration{}
	err = decoder.Decode(&AppConfig)
	if err != nil {
		log.Fatalf("Can not decode config file: %s\n", err)
	}
}

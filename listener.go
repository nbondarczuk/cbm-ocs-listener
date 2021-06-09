package main

import (
	"log"
	"os"

	"cbm-ocs-listener/gateway"
	"cbm-ocs-listener/common"
	"cbm-ocs-listener/ocs"
	"cbm-ocs-listener/responder"
)

// Variables set via loader -X flag in main module
var (
	version, build, level string
)

// Used in testing like sending json to other instance or processing had hoc file
func tryTestingMode() {
	if common.AppConfig.Send != "" {
		if err := gateway.SendMessage(common.AppConfig.Send); err != nil {
			log.Printf("Error sending message to CBM OCS listener: %s", err.Error())
			os.Exit(1)
		}
		os.Exit(0)
	}

	if common.AppConfig.Parse != "" {
		ocs.HandleEvent(common.AppConfig.Parse)
		os.Exit(0)
	}
}

// Main module initialization
func init() {
	common.LogInit(false)
	common.FlagsInit()
}

// Gateway listener server polling event queue
func main() {
	// Initialisation of the environment
	common.EnvInit(version, build, level)
	tryTestingMode()

	// Start all components of the server
	log.Printf("CBM OCS listner runing version: %s", common.GetVersion())
	log.Printf("Starting CBM OCS listener as PID: %d in RunPath: %s",
		os.Getpid(),
		common.AppConfig.RunPath)

	// Init environment like Oracle, queues, etc.
	common.StartUp()
	// Start health/alive/version responder 
	go responder.StartAll()

	// Start main evenent handling loop
	if err := gateway.ListenAndControl(); err != nil {
		// Error starting or closing listener:
		log.Printf("CBM OCS listener Error: %s", err)
		os.Exit(1)
	} else {
		log.Printf("CBM OCS listener stopped")
		os.Exit(0)
	}
}

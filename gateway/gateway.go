package gateway

import (
	"fmt"
	"log"

	"cbm-ocs-listener/ocs"
	"cbm-ocs-listener/common"
	. "cbm-ocs-listener/ems"
)

// Set by signal
var shutdown bool = false

// Stop next iteration of the message waiting loop upon signal
func Shutdown() (err error) {
	shutdown = true
	
	return
}

// Start polling for the event on the queue
func ListenAndControl() (err error) {
	log.Printf("EMS client start")
	options := NewClientOptions().
		SetServerUrl(common.AppConfig.QueueServerUrl).
		SetUsername(common.AppConfig.QueueUser).
		SetPassword(common.AppConfig.QueuePassword)
	client := NewClient(options).(*Client)
	if client == nil {
		return fmt.Errorf("Can not make new EMS client: %v", options)
	}
	log.Printf("EMS client created with options: %v", options)
		
	// Connect to the messaging server
	err = client.Connect()
	if err != nil {
		return fmt.Errorf("Can not connect to EMS server: %s", err.Error())
	}
	log.Printf("EMS client connected server: %s as: %s/%s",
		common.AppConfig.QueueServerUrl,
		common.AppConfig.QueueUser,
		common.AppConfig.QueuePassword)

	// Start the message reading loop
	var msg string
	log.Printf("EMS client loop started")
	for {
		// If interrupted by signal
		if shutdown {
			break
		}
		
		log.Printf("EMS client waiting for queue: %s", common.AppConfig.QueueName)
		msg, err = client.Receive(common.AppConfig.QueueName)
        if err != nil {
			fmt.Printf("Can not receive from EMS queue: %s", err.Error())
        } else {
			log.Printf("EMS client got event")
			go ocs.HandleEvent(msg)
		}
	}

	err = client.Disconnect()
	
	return
}

// Send message to the queue
func SendMessage(payload string) (err error) {
	options := NewClientOptions().
		SetServerUrl(common.AppConfig.QueueServerUrl).
		SetUsername(common.AppConfig.QueueUser).
		SetPassword(common.AppConfig.QueuePassword)
	client := NewClient(options).(*Client)
	if client == nil {
		return fmt.Errorf("Can not make new EMS client: %v", options)
	}

	// Connect to the messaging server
	err = client.Connect()
	if err != nil {
		return fmt.Errorf("Can not connect to EMS queue: %s", err.Error())
	}
	log.Printf("EMS server connected: %s as: %s/%s",
		common.AppConfig.QueueServerUrl,
		common.AppConfig.QueueUser,
		common.AppConfig.QueuePassword)

	// Sent the payload to the server
	err = client.Send(common.AppConfig.QueueName, payload, 0, "non_persistent", 10000)
	if err != nil {
		return fmt.Errorf("Can not send to queue: %s", err.Error())
	}
	log.Printf("Message sent to EMS server queue: %s", common.AppConfig.QueueName)
	
	err = client.Disconnect()

	return
}

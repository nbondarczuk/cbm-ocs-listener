package ocs

import (
	"fmt"
	"log"

	"cbm-ocs-listener/cbm"
	"cbm-ocs-listener/model"
	"cbm-ocs-listener/repository"
)

const (
	UNKNOWN_EVENT_ID string = ""
	MAX_EVENT_MSG_LEN int = 4000
)

// 1. Parse input json message from input queue
func eventOcsParse(rl *repository.CbmLogRepository, msg string) *model.OcsEvent {
	if eventOcs, err := ParseOcsEventJson(msg); err != nil {
		msg := fmt.Sprintf("Error parsing message: %s", err.Error())
		rl.Create(UNKNOWN_EVENT_ID, &model.CbmLog{Event: "OCS Event parse failed", InputParam: msg, OutputParam: msg})
		log.Printf(msg)
		panic(msg)
	} else {
		log.Printf("OCS Event json parsed: %s", eventOcs.EventId)
		rl.Create(eventOcs.EventId, &model.CbmLog{Event: "OCS Event parsed", InputParam: msg})
		return eventOcs
	}

	return nil
}

// 2. Decode OCS event struture to CBM event with all needed fields
func eventOcsDecode(rl *repository.CbmLogRepository, event *model.OcsEvent) *model.CbmOcsEvent {
	if eventCbm, err := DecodeOcsEvent(event); err != nil {
		msg := fmt.Sprintf("Error decoding OCS event: %s", err.Error())
		rl.Create(event.EventId, &model.CbmLog{Event: "OCS Event decode failed", OutputParam: msg})
		log.Printf(msg)
		panic(msg)
	} else {
		log.Printf("CBM OCS Event created")
		rl.Create(event.EventId, &model.CbmLog{Event: "OCS Event decoded"})
		return eventCbm
	}

	return nil
}

// 3. Create an record in db backend table
func eventCbmOcsCreate(rl *repository.CbmLogRepository, r *repository.CbmOcsEventRepository, eventOcs *model.OcsEvent, eventCbm *model.CbmOcsEvent) {
	if err := r.Create(eventCbm); err != nil {
		msg := fmt.Sprintf("Error creating CBM OCS event: %s", err.Error())
		rl.Create(eventOcs.EventId, &model.CbmLog{Event: "CBM OCS Event create failed", OutputParam: msg})
		log.Printf(msg)
		panic(msg)
	} else {
		log.Printf("CBM OCS Event stored in db")
		rl.Create(eventOcs.EventId, &model.CbmLog{Event: "CBM OCS Event created"})
	}
}

// 4. Trigger CBM action
func flowCbmStart(rl *repository.CbmLogRepository, msg string,  eventOcs *model.OcsEvent, eventCbm *model.CbmOcsEvent) {
	if err := cbm.FlowStart(msg, eventCbm); err != nil {
		msg := fmt.Sprintf("Error starting CBM flow: %s", err.Error())
		rl.Create(eventOcs.EventId, &model.CbmLog{Event: "CBM flow start failed", OutputParam: msg})
		log.Printf(msg)
		panic(msg)
	} else {
		rl.Create(eventOcs.EventId, &model.CbmLog{Event: "CBM flow started"})
	}
}

// Parse json of input message, decode it to CBM, store it and rigger CBM
func HandleEvent(msg string) {
	log.Printf("Handling message: %s", msg)

	var err error

	// Get db connection to store the event
	var r *repository.CbmOcsEventRepository
	r, err = repository.NewCbmOcsEventRepository()
	if err != nil {
		errmsg := fmt.Sprintf("Error creating CBM OCS event repository: %s", err.Error())
		log.Printf(errmsg)
		panic(errmsg)
	}
	log.Printf("CBM OCS repository opened")

	// Get db connection to store the log
	var rl *repository.CbmLogRepository
	rl, err = repository.NewCbmLogRepository()
	if err != nil {
		errmsg := fmt.Sprintf("Error creating CBM log repository: %s", err.Error())
		log.Printf(errmsg)
		panic(errmsg)
	}
	log.Printf("CBM log repository opened")

	// Processing chain of OCS event, may panic
	eventOcs := eventOcsParse(rl, msg)
	eventCbmOcs := eventOcsDecode(rl, eventOcs)
	eventCbmOcs.EventData = msg[0:MAX_EVENT_MSG_LEN-1]
	eventCbmOcsCreate(rl, r, eventOcs, eventCbmOcs)
	flowCbmStart(rl, msg, eventOcs, eventCbmOcs)

	log.Printf("Finished handling msisdn: %s", eventCbmOcs.Msisdn)

	return
}

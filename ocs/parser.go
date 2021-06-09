package ocs

import (
	"encoding/json"
	"log"

	"cbm-ocs-listener/model"
)

// Parse json from input message
func ParseOcsEventJson(msg string) (rv *model.OcsEvent, err error) {
	var event model.OcsEvent

	// Start parsing using model structure
	log.Printf("Parsing json message: %s", msg)
	if err = json.Unmarshal([]byte(msg), &event); err != nil {
		return
	}

	rv = &event

	return
}

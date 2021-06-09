package cbm

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"log"

	"cbm-ocs-listener/common"
	"cbm-ocs-listener/model"
)

// Start CBM flow for the new OCS event
func FlowStart(msg string, e *model.CbmOcsEvent) (err error) {
	// Skip is not address configured
	if common.AppConfig.CbmServerUrl == "" {
		log.Printf("Skip CBM flow")
		return
	}

	log.Printf("Starting CBM flow for: %+v", *e)

	// Make json to shoot
	if j, err := json.Marshal(*e); err != nil {
		return fmt.Errorf("Error encoding json: %s", err.Error())
	} else {
		// Shoot to kill
		if r, err := http.NewRequest("POST", common.AppConfig.CbmServerUrl, bytes.NewBuffer(j)); err != nil {
			return fmt.Errorf("Error building request: %s", err.Error())
		} else {
			client := &http.Client{}
			// Response not required so omit it
			if _, err := client.Do(r); err != nil {
				return fmt.Errorf("Error doing request: %s", err.Error())
			} else {
				log.Printf("CBM flow started")
			}
		}
	}

	return
}

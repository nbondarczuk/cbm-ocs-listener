package responder

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"cbm-ocs-listener/common"
	"cbm-ocs-listener/gateway"
)

func readPayload(r *http.Request) (body []byte, err error) {
	body, err = ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, fmt.Errorf("Can't read request body")
	}
        
	r.Body = ioutil.NopCloser(bytes.NewBuffer(body))
	
	return
}

func StartAll() {
	http.HandleFunc("/system/alive", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Got request: %s", r.URL.Path)
		w.Header().Set("Content-Type", "application/text; charset=utf-8")
		w.Write([]byte("I am alive"))
	})

	http.HandleFunc("/system/health", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Got request: %s", r.URL.Path)
		w.Header().Set("Content-Type", "application/text; charset=utf-8")
		w.Write([]byte("I am healthy"))
	})

	http.HandleFunc("/system/version", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Got request: %s", r.URL.Path)
		w.Header().Set("Content-Type", "application/text; charset=utf-8")
		w.Write([]byte(common.GetVersion()))
	})

	http.HandleFunc("/event", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Got request: %s", r.URL.Path)
		if payload, err := readPayload(r); err != nil {
			log.Printf("Error reading payload: %s", err.Error())
		} else if err := gateway.SendMessage(string(payload)); err != nil {
			log.Printf("Error sending message to CBM OCS listener: %s", err.Error())
		}
	})
	
	log.Printf("Started system and event responders")
	http.ListenAndServe(common.AppConfig.SystemResponderUrl, nil)
}

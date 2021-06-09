package common

import (
	"fmt"
	"io/ioutil"
	"log"
	"time"
)

var LogTimeFormat string = "2006-01-02 15:04:05.000000"

type logWriter struct{}

func (writer logWriter) Write(bytes []byte) (int, error) {
	return fmt.Print("[CBM-OCS-LISTENER] " + time.Now().Format(LogTimeFormat) + " " + string(bytes))
}

//
// Use custom log fomat
//
func LogInit(silent bool) {
	if !silent {
		log.SetFlags(0)
		log.SetOutput(new(logWriter))
	} else {
		log.SetOutput(ioutil.Discard)
	}
}

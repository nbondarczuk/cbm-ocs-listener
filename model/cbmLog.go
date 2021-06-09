package model

import (
	"time"
)

type (
	CbmLog struct {
		LogId      int       `json:"logId" db:"LOG_ID,primarykey"`
		LogTime    time.Time `json:"logTime" db:"LOG_TIME"`
		OcsEventId int       `json:"ocsEventId" db:"OCS_EVENT_ID"`
		Event      string    `json:"Event" db:"EVENT,size:300"`
		InputParam string    `json:"InputParam" db:"INPUT_PARAM,size:2000"`
		OutputParam string   `json:"OutputParam" db:"OUTPUT_PARAM,size:2000"`
	}
)

package model

import (
	"time"
)

type (
	CbmOcsEvent struct {
		EventId          int       `json:"eventId" db:"EVENT_ID,primarykey"`
		EventType        int       `json:"eventType" db:"EVENT_TYPE"`
		EventDate        time.Time `json:"eventDate" db:"EVENT_DATE"`
		State            string    `json:"state" db:"STATE"`
		Msisdn           string    `json:"msisdn" db:"MSISDN, size:100"`
		ExpireDate       time.Time `json:"expiredDate" db:"EXPIRE_DATE"`
		Value            int       `json:"value" db:"VALUE"`
		InitialAmount    int       `json:"initialAmount" db:"INITIAL_AMOUNT"`
		NextPir          string    `json:"nextPir" db:"NEXT_PIR, size:100"`
		StartDate        time.Time `json:"startDate" db:"START_DATE"`
		AccountId1       string    `json:"accountId1" db:"ACCOUNTID1, size:100"`
		AccountId2       string    `json:"accountId2" db:"ACCOUNTID2, size:100"`
		NotificationCode string    `json:"otificationCode" db:"NOTIFICATION_CODE, size:100"`
		EntryDate        time.Time `json:"entryDate" db:"ENTRY_DATE"`
		EventData        string    `json:"eventData" db:"EVENT_DATA, size:4000"`
	}
)

package ocs

import (
	"errors"
	"strconv"
	"time"

	"cbm-ocs-listener/model"
)

const (
	EVENT_STATE_NEW string = "N"
)

var (
	ErrorElementNotFound = errors.New("Element not found")
	OcsTimeFormat string = time.RFC3339Nano
	OcsPropertyTimeFormat string = "02/01/2006 15:04:05"
)

// Find a value on the list of mappings
func find(ps []model.Parameter, name string) (value string, err error) {
	for _, it := range ps {
		if it.Name == name {
			value = it.Value
			return
		}
	}

	err = ErrorElementNotFound

	return
}

// Find value which must be time.Time format
func findTime(mask string, ps []model.Parameter, name string) (ts time.Time, err error) {
	var str string
	str, err = find(ps, name)
	if err != nil {
		return
	}

	// Must be time formated string
	ts, err = time.Parse(mask, str)
	if err != nil {
		return
	}

	return
}

// Find value which must be int format
func findInt(ps []model.Parameter, name string) (value int, err error) {
	var str string
	str, err = find(ps, name)
	if err != nil {
		return
	}

	// Must be and integer
	value, err = strconv.Atoi(str)
	if err != nil {
		return
	}

	return
}

// Find all account values
func findMax2(ps []model.Parameter, name string) (v1, v2 string) {
	var n int
	for _, it := range ps {
		if it.Name == name {
			switch n {
			case 0:
				v1 = it.Value
			case 1:
				v2 = it.Value
			}
			n++
		}
	}

	return
}

// Map record + property list format to CBM specific record
func DecodeOcsEvent(e *model.OcsEvent) (rv *model.CbmOcsEvent, err error) {
	var event model.CbmOcsEvent

	// fixed values
	event.State = EVENT_STATE_NEW;

	// mandatory values
	if event.EventId, err = strconv.Atoi(e.EventId); err != nil {return}
	if event.EventType, err = strconv.Atoi(e.EventType); err != nil {return}
	if event.EventDate, err = time.Parse(OcsTimeFormat, e.SourceDate); err != nil {return}
	if event.Msisdn, err = find(e.Parameters.Parameter, "header1"); err != nil {return}
	if event.ExpireDate, err = findTime(OcsPropertyTimeFormat, e.Parameters.Parameter, "expireDate"); err != nil {return}
	if event.Value, err = findInt(e.Parameters.Parameter, "value"); err != nil {return}
	if event.InitialAmount, err = findInt(e.Parameters.Parameter, "initialAmount"); err != nil {return}
	if event.NextPir, err = find(e.Parameters.Parameter, "nextPIR"); err != nil {return}
	if event.StartDate, err = findTime(OcsPropertyTimeFormat, e.Parameters.Parameter, "startDate"); err != nil {return}
	if event.NotificationCode, err = find(e.Parameters.Parameter, "notificationcode"); err != nil {return}

	// optional values
	event.AccountId1, event.AccountId2 = findMax2(e.Parameters.Parameter, "accountInstanceIDList");

	// system values
	event.EntryDate = time.Now()

	rv = &event

	return
}

package repository

import (
	"fmt"
	"log"
	
	_ "gopkg.in/goracle.v2"

	"cbm-ocs-listener/common"
	"cbm-ocs-listener/model"
)

type CbmOcsEventRepository struct {
	Repository
}

func NewCbmOcsEventRepository() (r *CbmOcsEventRepository, err error) {
	if db, err := common.GetDbSession(); err != nil {
		return nil, err
	} else {
		dbmap := initRepository(db)
		dbmap.AddTableWithName(model.CbmOcsEvent{}, "OCS_EVENTS").
			SetKeys(false, "EVENT_ID")
		r = &CbmOcsEventRepository{
			Repository{
				Db:    db,
				Dbmap: dbmap,
			},
		}
	}

	return
}

func (r *CbmOcsEventRepository) Create(e *model.CbmOcsEvent) (err error) {
	log.Printf("Inserting to OCS_EVENTS: %#v", *e)

	if err = r.Dbmap.Insert(e); err != nil {
		return fmt.Errorf("Error in insert to OCS_EVENTS: %v", err.Error())
	}

	log.Printf("Inserted to OCS_EVENTS: %#v", *e)

	return
}

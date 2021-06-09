package repository

import (
	"fmt"
	"log"
	"strconv"
	"time"
	
	_ "gopkg.in/goracle.v2"

	"cbm-ocs-listener/common"
	"cbm-ocs-listener/model"
)

type CbmLogRepository struct {
	Repository
}

func NewCbmLogRepository() (r *CbmLogRepository, err error) {
	if db, err := common.GetDbSession(); err != nil {
		return nil, err
	} else {
		dbmap := initRepository(db)
		dbmap.AddTableWithName(model.CbmLog{}, "CBM_LOG").
			SetKeys(false, "LOG_ID")
		r = &CbmLogRepository{
			Repository{
				Db:    db,
				Dbmap: dbmap,
			},
		}
	}

	return
}

func (r *CbmLogRepository) Create(eventId string, l *model.CbmLog) {
	log.Printf("Inserting to CBM_LOG: %#v", *l)

	var err error
	
	// Prepare values
	l.LogId, err = r.NextId()
	if err != nil {
		panic(fmt.Sprintf("Sequence error: %s", err.Error()))
	}	
	l.LogTime = time.Now()
	l.OcsEventId, err = strconv.Atoi(eventId)
	if err != nil {
		panic(fmt.Sprintf("Conversion error: %s", err.Error()))
	}

	// Do insert
	err = r.Dbmap.Insert(l);
	if err != nil {
		panic(fmt.Sprintf("Error in insert to CBM_LOG: %s", err.Error()))
	}

	log.Printf("Inserted to CBM_LOG: %#v", *l)

	return
}

func (r *CbmLogRepository) NextId() (id int, err error) {
	rows, err := r.Db.Query("SELECT SEQ_CBM_LOG.NEXTVAL FROM DUAL")
	if err != nil {
		return 0, fmt.Errorf("Error getting nextval from SEQ_CBM_LOG: %v", err.Error())
	}

	for rows.Next() {
		if err = rows.Scan(&id); err != nil {
			return 0, fmt.Errorf("Error getting nextval from SEQ_CBM_LOG: %v", err.Error())
		}
	}

	log.Printf("SEQ_CBM_LOG.Nextvalue: %d", id)
	
	return
}


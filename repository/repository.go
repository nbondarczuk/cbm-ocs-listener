
package repository

import (
	"log"
	"os"
	
	"database/sql"
	_ "gopkg.in/goracle.v2"
	"gopkg.in/gorp.v2"
	
	"cbm-ocs-listener/common"
)

// Pepository being handled by request
type Repository struct {
	Db    *sql.DB
	Dbmap *gorp.DbMap
}

// Create GORP context and associate structures with table name
func initRepository(db *sql.DB) *gorp.DbMap {
	log.Printf("Initializing repository")

	dbmap := &gorp.DbMap{
		Db:      db,
		Dialect: &gorp.OracleDialect{},
	}

	if !common.TestRun {
		dbmap.TraceOn("GORP", log.New(os.Stdout, "[CBM-OCS-LISTENER] ", log.Lmicroseconds))
	}

	log.Printf("Initialized repository")

	return dbmap
}

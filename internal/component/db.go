package component

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql" // mysql driver
	"github.com/muhammad-fakhri/archetype-be/internal/config"
)

// initialization
type DB struct {
	Master *sql.DB
	Slave  *sql.DB
}

func InitDatabase() *DB {
	conf := config.Get()

	if conf.DBMaster == nil || conf.DBSlave == nil {
		log.Fatalf("failed to get DB config")
	}

	var (
		db  = &DB{}
		err error
	)

	db.Master, err = sql.Open("mysql", conf.DBMaster.ConnectionString)
	if err != nil {
		log.Fatalf("failed to open DB master connection. %+v", err)
	}
	db.Master.SetMaxIdleConns(conf.DBMaster.MaxIdle)
	db.Master.SetMaxOpenConns(conf.DBMaster.MaxOpen)
	err = db.Master.Ping()
	if err != nil {
		log.Fatalf("failed to ping DB master. %+v", err)
	}

	db.Slave, err = sql.Open("mysql", conf.DBSlave.ConnectionString)
	if err != nil {
		log.Fatalf("failed to open DB slave connection. %+v", err)
	}
	db.Slave.SetMaxIdleConns(conf.DBSlave.MaxIdle)
	db.Slave.SetMaxOpenConns(conf.DBSlave.MaxOpen)
	err = db.Slave.Ping()
	if err != nil {
		log.Fatalf("failed to ping DB slave. %+v", err)
	}

	return db
}

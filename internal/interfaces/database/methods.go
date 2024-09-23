package database

import (
	"database/sql"
	"github.com/caarlos0/env/v6"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"os"
	"path/filepath"
)

func New() (*Database, error) {

	var cfg Cfg
	e := env.Parse(&cfg)
	if e != nil {
		log.Fatal(e)
	}

	var dbFile string
	if cfg.DBFile != "" {
		dbFile = cfg.DBFile
	} else {
		appPath, err := os.Executable()
		if err != nil {
			log.Fatal(err)
		}
		dbFile = filepath.Join(filepath.Dir(appPath), "scheduler.db")
	}
	_, err := os.Create(dbFile)
	if err != nil {
		log.Fatal(err)
	}
	_, err = os.Stat(dbFile)

	var install bool
	if err == nil {
		install = true
	}

	database := new(Database)

	if install == true {
		db, err := sql.Open("sqlite3", dbFile)
		if err != nil {
			log.Fatal(err)
			return database, err
		}
		database.conn = db

		return database, nil
	}
	return database, err
}

func (db *Database) CreateTable() {

	statement := "CREATE TABLE IF NOT EXISTS scheduler (id INTEGER PRIMARY KEY AUTOINCREMENT, date VARCHAR(8) NULL, title VARCHAR(255) NOT NULL, comment VARCHAR(255) NOT NULL, repeat VARCHAR(128) NULL); CREATE INDEX date_INDEX ON scheduler (date ASC);"
	_, err := db.conn.Exec(statement)
	if err != nil {
		log.Fatal(err)
	}
}

func (db *Database) Close() {
	db.conn.Close()
}

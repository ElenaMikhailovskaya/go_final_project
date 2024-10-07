package database

import (
	"github.com/caarlos0/env/v6"
	"github.com/jmoiron/sqlx"
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

	// определяемся с путем до файла БД
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

	// проверка на существование файла БД, если нет, то создаем
	_, err := os.Stat(dbFile)
	if err != nil {
		_, err = os.Create(dbFile)
		if err != nil {
			log.Fatal(err)
		}
	}

	var install bool
	if err == nil {
		install = true
	}

	// подключение к БД
	database := new(Database)
	if install == true {
		db, err := sqlx.Connect("sqlite3", dbFile)
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

	statement := "CREATE TABLE IF NOT EXISTS scheduler (id INTEGER PRIMARY KEY AUTOINCREMENT, date VARCHAR(8) NULL, title VARCHAR(255) NOT NULL, comment VARCHAR(255) NOT NULL, repeat VARCHAR(128) NULL); CREATE INDEX IF NOT EXISTS date_INDEX ON scheduler (date ASC);"
	_, err := db.conn.Exec(statement)
	if err != nil {
		log.Fatal(err)
	}
}

func (db *Database) Close() {
	db.conn.Close()
}

package database

import "database/sql"

type Cfg struct {
	DBFile string `env:"TODO_DBFILE" envDefault:"./internal/interfaces/database/source/scheduler.db"`
}

type Database struct {
	conn *sql.DB
}

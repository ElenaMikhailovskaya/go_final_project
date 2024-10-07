package database

import (
	"github.com/jmoiron/sqlx"
)

type Cfg struct {
	DBFile string `env:"TODO_DBFILE" envDefault:"./internal/interfaces/database/source/scheduler.db"`
}

type Database struct {
	conn *sqlx.DB
}

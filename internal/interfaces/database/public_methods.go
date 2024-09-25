package database

import (
	"database/sql"
	"github.com/ElenaMikhailovskaya/go_final_project/internal/models"
)

func (db *Database) Add(task models.TaskRequest, nextDate string) (int, error) {
	res, err := db.conn.Exec("INSERT INTO scheduler (date, title, comment, repeat) VALUES (:date, :title, :comment, :repeat)",
		sql.Named("date", nextDate),
		sql.Named("title", task.Title),
		sql.Named("comment", task.Comment),
		sql.Named("repeat", task.Repeat),
	)
	if err != nil {
		return 0, err
	}
	id, _ := res.LastInsertId()
	return int(id), nil
}

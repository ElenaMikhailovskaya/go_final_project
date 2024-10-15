package database

import (
	"database/sql"
	"fmt"
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

func (db *Database) GetList() (models.TasksGetResponse, error) {
	tasks := models.TasksGetResponse{}
	rows, err := db.conn.Query("SELECT * FROM scheduler ORDER BY date ASC")
	if err != nil {
		return tasks, err
	}
	defer rows.Close()

	for rows.Next() {
		task := models.Task{}
		err := rows.Scan(&task.Id, &task.Date, &task.Title, &task.Comment, &task.Repeat)
		if err != nil {
			return tasks, err
		}

		tasks.Tasks = append(tasks.Tasks, task)
	}
	if err := rows.Err(); err != nil {
		return tasks, err
	}
	return tasks, nil
}

func (db *Database) GetById(id string) (models.Task, error) {
	task := models.Task{}
	rows, err := db.conn.Query("SELECT * FROM scheduler WHERE id = :id", sql.Named("id", id))

	if err != nil {
		return task, err
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&task.Id, &task.Date, &task.Title, &task.Comment, &task.Repeat)
		if err != nil {
			return task, err
		}
	}

	if err = rows.Err(); err != nil {
		return task, err
	}
	return task, nil
}

func (db *Database) Update(task models.TaskUpdateRequest) error {
	_, err := db.conn.Exec("UPDATE scheduler SET date = :date, title = :title, comment = :comment, repeat = :repeat WHERE id = :id",
		sql.Named("id", task.Id),
		sql.Named("date", task.Date),
		sql.Named("title", task.Title),
		sql.Named("comment", task.Comment),
		sql.Named("repeat", task.Repeat),
	)
	if err != nil {
		return err
	}
	return nil
}

func (db *Database) Delete(id string) error {
	_, err := db.conn.Exec("DELETE FROM scheduler WHERE id = :id", sql.Named("id", id))
	if err != nil {
		return err
	}
	return nil
}

func (db *Database) Search(search models.Search) (models.TasksGetResponse, error) {
	var rows *sql.Rows
	var err error
	tasks := models.TasksGetResponse{}
	if search.Date != "" {
		rows, err = db.conn.Query("SELECT * FROM scheduler where date = :query ORDER BY date ASC", sql.Named("id", search.Date))
	} else if search.Query != "" {
		rows, err = db.conn.Query("SELECT * FROM scheduler where title = :query ORDER BY date ASC", sql.Named("id", search.Query))
	} else {
		err = fmt.Errorf("Empty query")
	}
	if err != nil {
		return tasks, err
	}
	defer rows.Close()

	for rows.Next() {
		task := models.Task{}
		err := rows.Scan(&task.Id, &task.Date, &task.Title, &task.Comment, &task.Repeat)
		if err != nil {
			return tasks, err
		}

		tasks.Tasks = append(tasks.Tasks, task)
	}
	if err := rows.Err(); err != nil {
		return tasks, err
	}

	return tasks, nil
}

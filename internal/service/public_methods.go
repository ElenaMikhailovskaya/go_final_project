package service

import (
	"errors"
	"fmt"
	"github.com/ElenaMikhailovskaya/go_final_project/internal/models"
	"time"
)

func (s *Server) AddTask(task models.TaskRequest) (int, error) {

	var date string
	var nextDate string
	var newTaskId int
	var err error

	// проверка даты
	if task.Date == "" {
		date = time.Now().Format(models.DateFormat)
	} else {
		_, err = time.Parse(models.DateFormat, task.Date)
		if err != nil {
			return 0, errors.New("Invalid date format")
		}
		date = task.Date
	}

	// проверка параметра повторения
	if task.Repeat == "" {
		nextDate = time.Now().Format(models.DateFormat)
	} else {
		now := time.Now().Format(models.DateFormat)
		nextDate, err = s.NextDate(now, date, task.Repeat)
		if err != nil {
			return 0, err
		}
	}
	newTaskId, err = s.DBase.Add(task, nextDate)

	return newTaskId, nil
}

func (s *Server) UpdateTask(req models.TaskUpdateRequest) error {
	var nextDate string

	_, err := time.Parse(models.DateFormat, req.Date)
	if err != nil {
		return errors.New("Invalid date format")
	}

	task, err := s.GetById(req.Id)
	if err != nil || task.Id == "" {
		return fmt.Errorf("Task with id=%s not found", req.Id)
	}

	// проверка параметра повторения
	if req.Repeat == "" {
		nextDate = time.Now().Format(models.DateFormat)
	} else {
		now := time.Now().Format(models.DateFormat)
		nextDate, err = s.NextDate(now, req.Date, req.Repeat)
		if err != nil {
			return err
		}
	}
	req.Date = nextDate
	err = s.DBase.Update(req)

	return nil
}

func (s *Server) DeleteTask(id string) error {
	task, err := s.GetById(id)
	if err != nil || task.Id == "" {
		return fmt.Errorf("Task with id=%s not found", id)
	}

	err = s.DBase.Delete(id)

	return nil
}

func (s *Server) DoneTask(id string) error {
	task, err := s.GetById(id)
	if err != nil || task.Id == "" {
		return fmt.Errorf("Task with id=%s not found", id)
	}
	if task.Repeat == "" {
		err = s.DeleteTask(id)
		if err != nil {
			return err
		}
	} else {
		now := time.Now().Format(models.DateFormat)
		taskUpdate := models.TaskUpdateRequest{
			Id:      id,
			Date:    now,
			Title:   task.Title,
			Comment: task.Comment,
			Repeat:  task.Repeat,
		}
		err = s.UpdateTask(taskUpdate)
		if err != nil {
			return err
		}
	}
	return nil
}

func (s *Server) GetTaskList() (models.TasksGetResponse, error) {
	tasks := models.TasksGetResponse{}
	tasks, err := s.DBase.GetList()
	if err != nil {
		return tasks, err
	}
	if len(tasks.Tasks) == 0 {
		tasks.Tasks = []models.Task{}
	}

	return tasks, nil
}

func (s *Server) Search(query string) (models.TasksGetResponse, error) {
	tasks := models.TasksGetResponse{}
	search := models.Search{}
	timeQuery, err := time.Parse(models.DateFormatSearch, query)
	if err == nil {
		search.Date = timeQuery.Format(models.DateFormat)
	} else {
		search.Query = query
	}
	tasks, err = s.DBase.Search(search)
	if err != nil {
		return tasks, err
	}
	return tasks, nil
}

func (s *Server) GetById(id string) (models.Task, error) {
	task := models.Task{}
	task, err := s.DBase.GetById(id)
	if err != nil {
		return task, err
	}
	if task.Id == "" {
		return task, fmt.Errorf("Task with id=%s not found", id)
	}

	return task, nil
}

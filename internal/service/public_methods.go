package service

import (
	"github.com/ElenaMikhailovskaya/go_final_project/internal/models"
	"time"
)

func (s *Server) AddTask(task models.TaskRequest) ([]int, error) {

	var date string
	var nextDate []string
	var tasks []int

	if task.Date == "" {
		date = time.Now().Format(models.DateFormat)
	} else {
		date = task.Date
	}
	if task.Repeat == "" {
		nextDate = append(nextDate, time.Now().Format(models.DateFormat))
	} else {
		nextDate, err := NextDate(time.Now(), date, task.Repeat)
		if err != nil {
			return nil, err
		}

		for _, d := range nextDate {
			newTaskId, err := s.DBase.Add(task, d)
			if err != nil {
				return nil, err
			}
			tasks = append(tasks, newTaskId)
		}
	}

	return tasks, nil
}

func (s *Server) UpdateTask() {

}

func (s *Server) DeleteTask() {

}

func (s *Server) GetTask() {

}

func (s *Server) GetTaskList() {

}

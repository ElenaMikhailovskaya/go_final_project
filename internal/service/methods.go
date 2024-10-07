package service

import (
	"errors"
	"fmt"
	"github.com/ElenaMikhailovskaya/go_final_project/internal/models"
	"strconv"
	"strings"
	"time"
)

func (s *Server) NextDate(now string, date string, repeat string) (string, error) {
	var nextDate string
	var nextTime time.Time

	nowTime, err := time.Parse(models.DateFormat, now)
	if err != nil {
		return nextDate, errors.New("Wrong now field")
	}

	dateTime, err := time.Parse(models.DateFormat, date)
	if err != nil {
		return nextDate, errors.New("Wrong date field")
	}

	/*if dateTime.Before(nowTime) {
		return nextDate, errors.New("Date must be equal or after today")
	}*/

	var rulesDays map[string]map[int]struct{}
	rulesDays = make(map[string]map[int]struct{})

	rulesDays["week"] = make(map[int]struct{}, 7)
	for i := 1; i <= 7; i++ {
		rulesDays["week"][i] = struct{}{}
	}

	rulesDays["months"] = make(map[int]struct{}, 12)
	for i := 1; i <= 12; i++ {
		rulesDays["months"][i] = struct{}{}
	}

	rulesDays["days"] = make(map[int]struct{}, 31)
	for i := 1; i <= 31; i++ {
		rulesDays["days"][i] = struct{}{}
	}

	if repeat == "" {
		return nextDate, errors.New("Repeat field is empty")
	}

	repeatSlice := strings.Split(repeat, " ")

	if repeatSlice[0] == "" {
		return now, nil
	}

	switch repeatSlice[0] {
	case "y":
		nextTime = dateTime.AddDate(1, 0, 0)

		for nextTime.Before(nowTime) {
			nextTime = nextTime.AddDate(1, 0, 0)
		}

		nextDate = nextTime.Format(models.DateFormat)
		return nextDate, nil
	case "d":
		if repeatSlice[1] == "" {
			return nextDate, errors.New("Days is empty")
		} else {
			days, err := strconv.Atoi(repeatSlice[1])
			if dateTime.Format(models.DateFormat) == time.Now().Format(models.DateFormat) {
				nextTime = dateTime
			} else {
				if err != nil {
					return nextDate, fmt.Errorf("Days format error %s ", repeatSlice[1])
				}
				if days > MaxMoveDays {
					return nextDate, fmt.Errorf("Days more than %d ", MaxMoveDays)
				}

				nextTime = dateTime.AddDate(0, 0, days)

				for nextTime.Before(nowTime) {
					nextTime = nextTime.AddDate(0, 0, days)
				}
			}
		}
		nextDate = nextTime.Format(models.DateFormat)
		return nextDate, nil
	case "w":
		return nextDate, fmt.Errorf("Unprocessable operation")
		/*if len(repeatSlice) == 1 {
			return nextDate, errors.New("Days is empty")
		}
		days := strings.Split(repeatSlice[1], ",")
		if len(days) > 0 {
			for _, d := range days {
				weekDay, err := strconv.Atoi(d)
				if err != nil {
					return nextDate, fmt.Errorf("Weekday format error %s ", d)
				}

				_, ok := rulesDays["week"][weekDay]
				if !ok {
					return nextDate, fmt.Errorf("Weekday format error %s ", d)
				}

				curWD := int(dateTime.Weekday())
				var dd int
				if curWD > weekDay {
					dd = 7 - curWD + weekDay
				} else {
					dd = weekDay - curWD
				}

				var next time.Time
				for i := 1; i <= dd; i++ {
					next = dateTime.AddDate(0, 0, i)
					if next.After(nowTime) {
						nextDate = next.Format(models.DateFormat)
						continue
					}
				}
			}
		}*/
	case "m":
		return nextDate, fmt.Errorf("Unprocessable operation")
	default:
		return nextDate, fmt.Errorf("Unprocessable symbol %s ", repeatSlice[0])
	}

	return nextDate, nil
}

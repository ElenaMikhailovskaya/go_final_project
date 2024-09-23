package service

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"
)

func calculateNextDate(now time.Time, date string, repeat string) (string, error) {
	var nextDate string
	var nextTime time.Time

	var rulesDays map[string]map[int]struct{}
	rulesDays = make(map[string]map[int]struct{})

	for i := 1; i <= 7; i++ {
		//rulesDays["week"] = make(map[int]struct{})
		rulesDays["week"][i] = struct{}{}
	}
	for i := 1; i <= 12; i++ {
		//rulesDays["months"] = make(map[int]struct{})
		rulesDays["months"][i] = struct{}{}
	}

	if repeat == "" {
		return nextDate, errors.New("Repeat field is empty")
	}

	dateTime, err := time.Parse("20060102", date)
	if err != nil {
		return nextDate, errors.New("Repeat field is empty")
	}

	repeatSlice := strings.Split(repeat, " ")

	switch repeatSlice[0] {
	case "d":
		if repeatSlice[1] == "" {
			return nextDate, errors.New("Days is empty")
		}
		if repeatSlice[1] == "y" {
			nextTime = dateTime.AddDate(1, 0, 0)
		} else {
			days, err := strconv.Atoi(repeatSlice[1])
			if err != nil {
				return nextDate, fmt.Errorf("Days format error %s ", repeatSlice[1])
			}
			if days > MaxMoveDays {
				return nextDate, fmt.Errorf("Days more than %s ", MaxMoveDays)
			}
			nextTime = dateTime.AddDate(1, 0, days)
		}
		nextDate = nextTime.String()
		return nextDate, nil
	case "w":
		if repeatSlice[1] == "" {
			return nextDate, errors.New("Days is empty")
		}
		weekDay, err := strconv.Atoi(repeatSlice[1])
		if err != nil {
			return nextDate, fmt.Errorf("Weekday format error %s ", repeatSlice[1])
		}
		wd, ok := rulesDays["week"][weekDay]
		if !ok {
			return nextDate, fmt.Errorf("Weekday format error %s ", repeatSlice[1])
		}

	case "m":
	default:
		return nextDate, fmt.Errorf("Unprocessable symbol %s ", repeatSlice[0])
	}

	return nextDate, nil
}

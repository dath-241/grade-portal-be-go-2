package helper

import (
	"fmt"
	"time"
)

type Semester struct {
	CURRENT string
	NEXT    string
	PREV    string
}

func SetSemester() Semester {
	t := time.Now()
	year, month := t.Year(), t.Month()
	semesterYear := year - 2000

	var current, next, prev string
	switch {
	case month >= 9 && month <= 12:
		current = fmt.Sprintf("HK%d1", semesterYear)
		next = fmt.Sprintf("HK%d2", semesterYear)
		prev = fmt.Sprintf("HK%d3", semesterYear-1)
	case month >= 1 && month <= 4:
		current = fmt.Sprintf("HK%d2", semesterYear-1)
		next = fmt.Sprintf("HK%d3", semesterYear-1)
		prev = fmt.Sprintf("HK%d1", semesterYear-1)
	case month >= 5 && month <= 8:
		current = fmt.Sprintf("HK%d3", semesterYear-1)
		next = fmt.Sprintf("HK%d1", semesterYear)
		prev = fmt.Sprintf("HK%d2", semesterYear-1)
	}
	return Semester{CURRENT: current, NEXT: next, PREV: prev}
}

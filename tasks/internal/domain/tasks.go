package domain

import (
	"errors"
	"strings"

	"gorm.io/gorm"
)

type Task struct {
	UserID  string
	Summary string
	gorm.Model
}

func New(uid string, summary string) (Task, error) {

	trimSummary := strings.TrimSpace(summary)

	var task Task

	var err error

	if len(trimSummary) > 2500 {
		err = errors.New("task summary exceeds 2500 characters")
	} else {
		task = Task{
			UserID:  uid,
			Summary: trimSummary,
		}
	}

	return task, err
}

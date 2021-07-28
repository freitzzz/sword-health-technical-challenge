package domain

import (
	"errors"
	"strings"

	"gorm.io/gorm"
)

type Task struct {
	UserID   string
	Summary  string
	Disabled bool
	gorm.Model
}

func New(uid string, summary string) (Task, error) {

	trimSummary, err := trimAndCheckIfSummaryStringExceeds2500Characters(summary)

	var task Task

	if err == nil {
		task = Task{
			UserID:   uid,
			Summary:  trimSummary,
			Disabled: false,
		}
	}

	return task, err
}

func Disable(task *Task) {
	task.Disabled = true
}

func UpdateSummary(task *Task, summary string) error {

	trimSummary, err := trimAndCheckIfSummaryStringExceeds2500Characters(summary)

	if err == nil {
		task.Summary = trimSummary
	}

	return err
}

func trimAndCheckIfSummaryStringExceeds2500Characters(summary string) (string, error) {
	trimSummary := strings.TrimSpace(summary)

	if len(trimSummary) > 2500 {
		return trimSummary, errors.New("task summary exceeds 2500 characters")
	} else {
		return trimSummary, nil
	}
}

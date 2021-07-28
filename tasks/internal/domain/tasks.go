package domain

import (
	"crypto/cipher"
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

func New(uid string, summary string, block cipher.Block) (Task, error) {

	trimSummary, err := trimAndCheckIfSummaryStringExceeds2500Characters(summary)

	var task Task

	if err == nil {
		task = Task{
			UserID:   uid,
			Summary:  encryptSummary(block, trimSummary),
			Disabled: false,
		}
	}

	return task, err
}

func Disable(task *Task) {
	task.Disabled = true
}

func UpdateSummary(task *Task, summary string, block cipher.Block) error {

	trimSummary, err := trimAndCheckIfSummaryStringExceeds2500Characters(summary)

	if err == nil {
		task.Summary = encryptSummary(block, trimSummary)
	}

	return err
}

func Summary(task Task, block cipher.Block) string {
	return decryptSummary(block, task.Summary)
}

func trimAndCheckIfSummaryStringExceeds2500Characters(summary string) (string, error) {
	trimSummary := strings.TrimSpace(summary)

	if len(trimSummary) > 2500 {
		return trimSummary, errors.New("task summary exceeds 2500 characters")
	} else {
		return trimSummary, nil
	}
}

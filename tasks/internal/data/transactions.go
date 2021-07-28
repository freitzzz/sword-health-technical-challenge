package data

import (
	"gorm.io/gorm"

	"github.com/freitzzz/sword-health-technical-challenge/tasks/internal/domain"
)

const (
	queryResultsLimit = 20
)

var (
	disabledQueryMap = map[string]interface{}{"disabled": false}
)

func QueryUserTasks(db *gorm.DB, uid string, pidx int) []*domain.Task {

	var tasks []*domain.Task

	offset := PaginationIndexToOffset(pidx)

	db.Limit(queryResultsLimit).Offset(offset).Where(&domain.Task{UserID: uid}, "disabled").Find(&tasks)

	return tasks

}

func QueryTasks(db *gorm.DB, pidx int) []*domain.Task {

	var tasks []*domain.Task

	offset := PaginationIndexToOffset(pidx)

	db.Limit(queryResultsLimit).Offset(offset).Where(disabledQueryMap).Find(&tasks)

	return tasks

}

func InsertTask(db *gorm.DB, task domain.Task) (*domain.Task, error) {

	result := db.Create(&task)

	return &task, result.Error

}

// todo: remove "external" identification, use internal only. Justify no time for external

func QueryTaskById(db *gorm.DB, tid int) (*domain.Task, error) {

	var task domain.Task

	result := db.Where(disabledQueryMap).First(&task, tid, "disabled = ?", "false")

	return &task, result.Error

}

func QueryUserTaskById(db *gorm.DB, uid string, tid int) (*domain.Task, error) {

	var task domain.Task

	result := db.Where(&domain.Task{UserID: uid}, "disabled").First(&task, tid)

	return &task, result.Error

}

func UpdateTask(db *gorm.DB, task domain.Task) (*domain.Task, error) {

	result := db.Save(&task)

	return &task, result.Error

}

func PaginationIndexToOffset(pidx int) int {
	return queryResultsLimit * pidx
}

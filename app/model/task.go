package model

import (
	"gorm.io/gorm"
)

type taskStatus int

const (
	INCOMPLETE taskStatus = 0
	Complete   taskStatus = 1
)

type Task struct {
	ID     int    `json:"id" gorm:"primary_key;autoIncrement"`
	Name   string `json:"name"`
	Status int    `json:"status"`
}

func (t *Task) BeforeSave(tx *gorm.DB) (err error) {
	if !validateStatus(t.Status) {
		err = tx.AddError(gorm.ErrInvalidData)
	}
	return err
}

func validateStatus(status int) bool {
	return taskStatus(status) == INCOMPLETE || taskStatus(status) == Complete
}

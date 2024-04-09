package model

import (
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"testing"
)

func TestTaskModelBeforeSave(t *testing.T) {
	mockDb, _, _ := sqlmock.New()
	dialector := postgres.New(postgres.Config{
		Conn:       mockDb,
		DriverName: "postgres",
	})
	db, _ := gorm.Open(dialector, &gorm.Config{})

	task := Task{
		Name:   "Task 1",
		Status: 9,
	}

	task2 := Task{
		Name:   "Task 2",
		Status: 1,
	}

	err := task.BeforeSave(db)
	assert.Error(t, err)

	err = task2.BeforeSave(db)
	assert.NoError(t, err)
}

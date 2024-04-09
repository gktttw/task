package repository

import (
	"errors"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"task/app/model"
	"task/db/migration"
)

type TaskRepository interface {
	GetAllTask() ([]model.Task, error)
	UpdateTask(ID int, name string, status int) (*model.Task, error)
	CreateTask(name string, status int) (model.Task, error)
	DeleteTask(id int) error
}

type TaskRepositoryImpl struct {
	db *gorm.DB
}

func (t TaskRepositoryImpl) GetAllTask() ([]model.Task, error) {
	var tasks []model.Task

	result := t.db.Order("id asc").Find(&tasks)

	if result.Error != nil {
		log.Error("Got an error finding all tasks. Error: ", result.Error)
		return nil, result.Error
	}

	return tasks, nil
}

func (t TaskRepositoryImpl) UpdateTask(ID int, name string, status int) (*model.Task, error) {
	task := &model.Task{}
	result := t.db.Where("id = ?", ID).First(&task)
	if result.Error != nil {
		log.Error("Got an error when finding a task. Error: ", result.Error)
		return nil, result.Error
	}
	if result.RowsAffected == 0 {
		log.Error("Got an error when finding a task. Error: ", result.Error)
		return nil, errors.New("Task not found")
	}

	task.Name = name
	task.Status = status

	result = t.db.Save(&task)

	if result.Error != nil {
		log.Error("Got an error when updating a task. Error: ", result.Error)
		return &model.Task{}, result.Error
	}
	return task, nil
}

func (t TaskRepositoryImpl) CreateTask(name string, status int) (model.Task, error) {
	task := model.Task{
		Name:   name,
		Status: status,
	}

	err := t.db.Create(&task).Error
	if err != nil {
		log.Error("Got and error when creating tasks. Error: ", err)
		return model.Task{}, err
	}
	return task, nil
}

func (t TaskRepositoryImpl) DeleteTask(id int) error {
	err := t.db.Delete(model.Task{}, id).Error
	if err != nil {
		log.Error("Got an error when deleting task by id. Error: ", err)
		return err
	}
	return nil
}

func TaskRepositoryInit(db *gorm.DB) *TaskRepositoryImpl {
	migration.Migrate(db)

	return &TaskRepositoryImpl{
		db: db,
	}
}

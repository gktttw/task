//go:generate mockery --name=TaskService --output=../../mocks
package service

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"net/http"
	"strconv"
	"task/app/constant"
	"task/app/dto"
	"task/app/model"
	"task/app/pkg"
	"task/app/repository"
)

type TaskService interface {
	GetAllTask(c *gin.Context)
	CreateTask(c *gin.Context)
	UpdateTask(c *gin.Context)
	DeleteTask(c *gin.Context)
}

type TaskServiceImpl struct {
	taskRepository repository.TaskRepository
}

func (t TaskServiceImpl) GetAllTask(c *gin.Context) {
	defer pkg.PanicHandler(c)

	data, err := t.taskRepository.GetAllTask()
	if err != nil {
		log.Error("Happened error when get data from database. Error", err)
		pkg.PanicException(constant.UnknownError)
	}

	c.JSON(http.StatusOK, data)
}

func (t TaskServiceImpl) CreateTask(c *gin.Context) {
	defer pkg.PanicHandler(c)

	var request dto.TaskRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		log.Error("Invalid create params", err)
	}

	taskName := request.Name
	taskStatus := request.Status

	data, err := t.taskRepository.CreateTask(taskName, taskStatus)
	if err != nil {
		log.Error("Happened error when get data from database. Error", err)

		var errorConst constant.ResponseStatus
		if err.Error() == gorm.ErrInvalidData.Error() {
			errorConst = constant.InvalidRequest
		} else {
			errorConst = constant.UnknownError
		}
		pkg.PanicException(errorConst)
	}

	c.JSON(http.StatusOK, data)
}

func (t TaskServiceImpl) UpdateTask(c *gin.Context) {
	defer pkg.PanicHandler(c)
	taskID, _ := strconv.Atoi(c.Param("taskID"))

	var request dto.TaskRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		log.Error("Invalid create params", err)
		pkg.PanicException(constant.InvalidRequest)
	}

	taskName := request.Name
	taskStatus := request.Status

	task := &model.Task{
		ID: taskID,
	}

	task, err := t.taskRepository.UpdateTask(taskID, taskName, taskStatus)
	if err != nil {
		log.Error("Happened error when get data from database. Error", err)
		pkg.PanicException(constant.InvalidRequest)
	}

	c.JSON(http.StatusOK, task)
}

func (t TaskServiceImpl) DeleteTask(c *gin.Context) {
	defer pkg.PanicHandler(c)
	taskID, _ := strconv.Atoi(c.Param("taskID"))

	err := t.taskRepository.DeleteTask(taskID)
	if err != nil {
		log.Error("Happened error when get data from database. Error", err)
		pkg.PanicException(constant.UnknownError)
	}

	c.JSON(http.StatusOK, gin.H{"message": "Task deleted successfully"})
}

func TaskServiceInit(taskRepository repository.TaskRepository) *TaskServiceImpl {
	return &TaskServiceImpl{
		taskRepository: taskRepository,
	}
}

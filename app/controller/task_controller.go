package controller

import (
	"github.com/gin-gonic/gin"
	"task/app/service"
)

type TaskController interface {
	GetAllTaskData(c *gin.Context)
	CreateTask(c *gin.Context)
	UpdateTask(c *gin.Context)
	DeleteTask(c *gin.Context)
}

type TaskControllerImpl struct {
	svc service.TaskService
}

func (u TaskControllerImpl) GetAllTaskData(c *gin.Context) {
	u.svc.GetAllTask(c)
}

func (u TaskControllerImpl) CreateTask(c *gin.Context) {
	u.svc.CreateTask(c)
}

func (u TaskControllerImpl) UpdateTask(c *gin.Context) {
	u.svc.UpdateTask(c)
}

func (u TaskControllerImpl) DeleteTask(c *gin.Context) {
	u.svc.DeleteTask(c)
}

func TaskControllerInit(userService service.TaskService) *TaskControllerImpl {
	return &TaskControllerImpl{
		svc: userService,
	}
}

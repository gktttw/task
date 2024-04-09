package config

import (
	"task/app/controller"
	"task/app/repository"
	"task/app/service"
)

type Initialization struct {
	TaskRepo repository.TaskRepository
	TaskSvc  service.TaskService
	TaskCtrl controller.TaskController
}

func NewInitialization(taskRepo repository.TaskRepository,
	taskService service.TaskService,
	taskCtrl controller.TaskController,
) *Initialization {
	return &Initialization{
		TaskRepo: taskRepo,
		TaskSvc:  taskService,
		TaskCtrl: taskCtrl,
	}
}

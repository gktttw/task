package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"task/config"
)

func Init(init *config.Initialization) *gin.Engine {

	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": "hello world"})
	})

	api := router.Group("/tasks")
	api.GET("", init.TaskCtrl.GetAllTaskData)
	api.POST("", init.TaskCtrl.CreateTask)
	api.PUT("/:taskID", init.TaskCtrl.UpdateTask)
	api.DELETE("/:taskID", init.TaskCtrl.DeleteTask)

	return router
}

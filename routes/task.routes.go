package routes

import (
	"net/http"

	"github.com/GuilleFB/go-party/handlers"
	"github.com/gin-gonic/gin"
)

func RoutesTask(rgin *gin.Engine) {
	rgin.GET("/gin/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})
	
	rgin.GET("/gin/tasks", handlers.GetTasksHandler)
	rgin.GET("/gin/task/:id", handlers.GetTaskHandler)
	rgin.POST("/gin/task/create", handlers.CreateTaskHandler)
	rgin.PATCH("/gin/task/edit/:id", handlers.EditTaskHandler)
	// rgin.DELETE("/gin/task/delete/:id", handlers.DeleteTaskHandler)
}
package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RoutesTask(rgin *gin.Engine) {
	rgin.GET("/gin/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})
	
	// rgin.GET("/tasks", handlers.GetTasksHandler)
	// rgin.GET("/task/{id}", handlers.GetTaskHandler)
	// rgin.POST("/task/create", handlers.PostTaskHandler)
	// rgin.PATCH("/task/edit/{id}", handlers.EditTaskHandler)
	// rgin.DELETE("/task/delete/{id}", handlers.DeleteTaskHandler)
}
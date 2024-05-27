package handlers

import (
	"github.com/GuilleFB/go-party/db"
	"github.com/GuilleFB/go-party/models"
	"github.com/gin-gonic/gin"
)

func GetTasksHandler(c *gin.Context) {
	var tasks []models.Task

	result := db.DB.Find(&tasks)

	if result.Error != nil {
		c.Status(400)
		return
	}

	c.JSON(200, gin.H{ // In Map format
		"tasks": tasks,
	})

}

func GetTaskHandler(c *gin.Context) {
	var task models.Task

	params := c.Param("id")

	db.DB.First(&task, params)

	if task.ID == 0 {
		message := "Task not found"
		c.String(404, message)
		return
	}

	c.JSON(200, task)
}

func CreateTaskHandler(c *gin.Context) {
	var task models.Task

	c.Bind(&task)

	createdTask := db.DB.Create(&task)

	err := createdTask.Error

	if err != nil {
		c.Status(400)
		return
	}

	c.JSON(200, task)

}

func EditTaskHandler(c *gin.Context) {
	var task models.Task

	params := c.Param("id")

	db.DB.First(&task, params)

	if task.ID == 0 {
		message := "Task not found"
		c.String(404, message)
		return
	}

	// var updates map[string]interface{}
	// if err := c.BindJSON(&updates); err != nil {
	//     c.String(400, "Invalid request body")
	//     return
	// }

	c.BindJSON(&task)

	db.DB.Model(&task).Updates(task)

	// db.DB.Save(&user) // Save is a combination function. If save value does not contain primary key, it will execute Create, otherwise it will execute Update (with all fields).
	// Actualizar solo los campos proporcionados en la solicitud
	// if err := db.DB.Model(&task).Updates(task).Error; err != nil { // https://gorm.io/docs/update.html#Updates-multiple-columns
	//     c.String(500, "Failed to update task")
	//     return
	// }

	c.JSON(200, task)
}

func DeleteTaskHandler(c *gin.Context) {
	var task models.Task

	params := c.Param("id")

	db.DB.First(&task, params)

	if task.ID == 0 {
		message := "Task not found"
		c.String(404, message)
		return
	}

	db.DB.Unscoped().Delete(&task) // Hard delete https://gorm.io/docs/delete.html#Delete-permanently
	// db.DB.Delete(&user) // Soft delete https://gorm.io/docs/delete.html#Soft-Delete

	c.String(200, "Taks Deleted")
}

package tests

import (
	"github.com/gin-gonic/gin"
	"github.com/holanda-lucas/gotasks-back/database"
	"github.com/holanda-lucas/gotasks-back/models"
)


var UserId uint
var TaskId uint

func GetTestRouter() *gin.Engine{
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	// Conectando com o banco de dados
	database.ConnectWithDatabase()

	return r
}

func RegisterMockUser() {
	user := models.User{Email: "test@test.test", Name: "test", Password: "test"}
	UserId = database.RegisterUser(&user)
}

func DeleteMockUser() {
	database.DeleteUser(UserId)
}

func RegisterMockTask() {
	task := models.Task{User_id: UserId, Title: "test", Description: "test", Tag: "white"}
	TaskId = database.RegisterTask(&task)
}

func DeleteMockTask() {
	database.DeleteTask(TaskId)
}
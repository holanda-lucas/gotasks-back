package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/holanda-lucas/gotasks-back/database"
	"github.com/holanda-lucas/gotasks-back/models"
)

func GetTask(c *gin.Context) {
	id := c.Param("id")

	intId, err := strconv.Atoi(id)
	uintId := uint(intId)

	// Verificando se o id é inteiro
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"erro":"insira um id inteiro.",
		})
	}

	task := database.GetTask(uintId)

	if task.Title == "" {
		c.JSON(http.StatusBadRequest, gin.H {
			"erro": "Tarefa não encontrada.",
		})	
		return
	}

	c.JSON(http.StatusOK, task)
}

func DeleteTask(c *gin.Context) {
	id := c.Param("id")

	intId, err := strconv.Atoi(id)
	uintId := uint(intId)

	// Verificando se o id é inteiro
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"erro":"insira um id inteiro.",
		})
	}

	database.DeleteTask(uintId)

	c.JSON(http.StatusOK, gin.H{
		"sucesso":"tarefa deletada",
	})
}

func CreateTask(c *gin.Context) {
	var task models.Task

	err := c.ShouldBindJSON(&task)

	// Erro nos dados
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	user := database.GetUser(task.User_id)
	if user.Name == "" {
		c.JSON(http.StatusBadRequest, gin.H {
			"erro": "Não é possível criar uma tarefa para um usuário inexistente.",
		})
		return
	}
	
	if task.Tag == "" {
		task.Tag = "white"
	}

	database.RegisterTask(&task)
	c.JSON(http.StatusAccepted, task)
}

func EditTask(c *gin.Context) {
	id := c.Param("id")
	intId, err := strconv.Atoi(id)
	uintId := uint(intId)
	
	user := database.GetUser(uintId)

	if user.Email == "" {
		c.JSON(http.StatusBadRequest, gin.H {
			"erro": "Usuário não encontrado",
		})
		return
	}

	// Erro nos dados
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	var newTask models.Task

	err = c.ShouldBindJSON(&newTask)

	// Erro nos dados
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// Setando o ID para saber qual registro alterar
	newTask.ID = uint(intId)

	database.EditTask(&newTask)

	c.JSON(http.StatusOK, newTask)
}
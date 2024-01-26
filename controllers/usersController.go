package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/holanda-lucas/gotasks-back/database"
	"github.com/holanda-lucas/gotasks-back/models"
	"github.com/holanda-lucas/gotasks-back/util"
)

func GetUser(c *gin.Context) {
	id := c.Param("id")

	intId, err := strconv.Atoi(id)

	// Verificando se o id é inteiro
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"erro":"insira um número inteiro como id.",
		})
	}

	user := database.GetUser(intId)

	if user.Email == "" {
		c.JSON(http.StatusBadRequest, gin.H {
			"erro": "Usuário não encontrado",
		})
		return
	}

	c.JSON(http.StatusOK, user)
}

func DeleteUser(c *gin.Context) {
	id := c.Param("id")

	intId, err := strconv.Atoi(id)
	uintId := uint(intId)

	// Verificando se o id é inteiro
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"erro":"insira um id inteiro.",
		})
		return
	}

	database.DeleteUser(uintId)

	// Deletando as tasks associadas ao user deletado
	userTasks := database.GetTasksFromUser(uintId)
	for _, task := range userTasks {
		database.DeleteTask(task.ID)
	}

	c.JSON(http.StatusOK, gin.H{
		"sucesso":"usuário deletado",
	})
}

func CreateUser(c *gin.Context) {
	var user models.User

	err := c.ShouldBindJSON(&user)

	// Erro no bind
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H {
			"erro": "formato de informações errado",
		})
	}

	hash, err := util.HashPassword(user.Password)

	// Erro no bcrypt
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H {
			"erro": "não foi possível criptografar sua senha",
		})
	}

	user.Password = hash

	// Erro nos dados
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	database.RegisterUser(&user)

	jsonUser := database.GetUser(user.ID)


	c.JSON(http.StatusAccepted, jsonUser)
}

func EditUser(c *gin.Context) {
	id := c.Param("id")
	intId, err := strconv.Atoi(id)
	uintId := uint(intId)

	// Erro nos dados
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	var newUser models.User

	err = c.ShouldBindJSON(&newUser)

	// Erro nos dados
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// Setando o ID para saber qual registro alterar
	newUser.ID = uint(intId)

	user := database.GetUser(uintId)

	if user.Email == "" {
		c.JSON(http.StatusBadRequest, gin.H {
			"erro": "Usuário não encontrado",
		})
		return
	}
	database.EditUser(&newUser)

	newUserJSON := database.GetUser(newUser.ID)
	c.JSON(http.StatusOK, newUserJSON)
}

func GetUserTasks(c *gin.Context) {
	id := c.Param("id")

	intId, err := strconv.Atoi(id)
	uintId := uint(intId)

	// Verificando se o id é inteiro
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"erro":"insira um número inteiro como id.",
		})
	}

	user := database.GetUser(uintId)

	if user.Email == "" {
		c.JSON(http.StatusBadRequest, gin.H {
			"erro": "Usuário não encontrado",
		})
		return
	}

	tasks := database.GetTasksFromUser(uintId)

	c.JSON(http.StatusAccepted, tasks)
}
package tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/holanda-lucas/gotasks-back/controllers"
	"github.com/holanda-lucas/gotasks-back/database"
	"github.com/holanda-lucas/gotasks-back/models"
	"github.com/stretchr/testify/assert"
)

func TestGetUserById(t *testing.T) {
	// Preparando o router
	r := GetTestRouter()
	r.GET("/users/:id", controllers.GetUser)

	// Criando usuário mock
	RegisterMockUser()
	defer DeleteMockUser()

	strId := strconv.Itoa(int(UserId))

	// Preparando requisição
	req, _ := http.NewRequest("GET", "/users/" + strId, nil)
	res := httptest.NewRecorder()

	// Executando requisição
	r.ServeHTTP(res, req)
	
	// Puxando dados da resposta
	var respondedUser models.User
	err := json.Unmarshal(res.Body.Bytes(), &respondedUser)
	assert.NoError(t, err)
	
	// Fazendo as verificações
	assert.Equal(t, http.StatusOK, res.Code)
	assert.Equal(t, "test@test.test", respondedUser.Email)
	assert.Equal(t, "test", respondedUser.Name)
}

func TestDeleteUser(t *testing.T) {
	// Preparando o router
	r := GetTestRouter()
	r.DELETE("/users/:id", controllers.DeleteUser)

	// Registrando mock
	RegisterMockUser()

	strId := strconv.Itoa(int(UserId))

	// Preparando requisição
	req, _ := http.NewRequest("DELETE", "/users/" + strId, nil)
	res := httptest.NewRecorder()

	// Executando requisição
	r.ServeHTTP(res, req)

	// Fazendo verificação
	assert.Equal(t, http.StatusNoContent, res.Code)
}

func TestCreateUser(t *testing.T) {
	// Preparando o router
	r := GetTestRouter()
	r.POST("/users", controllers.CreateUser)

	// Criando json de usuario mock
	user := models.User{Email: "test@test.test", Name: "test", Password: "test"}
	jsonData, err := json.Marshal(user)

	if err != nil {
		panic(err)
	}

	// Preparando requisição
	req, _ := http.NewRequest("POST", "/users", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")
	res := httptest.NewRecorder()

	// Executando requisição
	r.ServeHTTP(res, req)

	// Puxando dados da resposta
	var respondedUser models.User
	err = json.Unmarshal(res.Body.Bytes(), &respondedUser)
	assert.NoError(t, err)

	// Deletando o usuário criado
	defer database.DeleteUser(respondedUser.ID)

	// Fazendo verificaçao
	assert.Equal(t, http.StatusAccepted, res.Code)
	assert.Equal(t, "test", respondedUser.Name)
	assert.Equal(t, "test@test.test", respondedUser.Email)
}

func TestEditUser(t *testing.T) {
	// Preparando o router
	r := GetTestRouter()
	r.PUT("/users/:id", controllers.EditUser)

	// Registrando mock
	RegisterMockUser()
	strId := strconv.Itoa(int(UserId))

	// Criando json de atualizações do mock
	user := models.User{Email: "test2@test.test", Name: "test2", Password: "test2"}
	jsonData, err := json.Marshal(user)

	if err != nil {
		panic(err)
	}

	// Preparando requisição
	req, _ := http.NewRequest("PUT", "/users/" + strId, bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")
	res := httptest.NewRecorder()

	// Executando a requisição
	r.ServeHTTP(res, req)

	// Puxando dados da resposta
	var respondedUser models.User
	err = json.Unmarshal(res.Body.Bytes(), &respondedUser)
	assert.NoError(t, err)

	// Deletando o usuário criado
	defer database.DeleteUser(respondedUser.ID)

	// Fazendo verificaçao
	assert.Equal(t, http.StatusOK, res.Code)
	assert.Equal(t, "test2", respondedUser.Name)
	assert.Equal(t, "test2@test.test", respondedUser.Email)
}

func TestAuthenticateUser(t *testing.T) {
	
}
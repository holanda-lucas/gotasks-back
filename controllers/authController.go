package controllers

import (
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/holanda-lucas/gotasks-back/database"
	"github.com/holanda-lucas/gotasks-back/models"
)

func Login (c *gin.Context) {
	var loginData models.UserLoginData

	c.ShouldBindJSON(&loginData)
	id := database.AuthenticateUser(loginData)

	if id != 0 {
		
		// Cria um token com a claim do id
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"sub": id,
			"exp": time.Now().Add(time.Hour * 4).Unix(),
		})
		
		// Assinando o token com a chave secreta
		key := []byte(os.Getenv("SECRET_KEY"))
		tokenString, err := token.SignedString(key)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H {
				"erro": err.Error(),
			})
			return
		}

		// Enviando um cookie com o token
		finalToken := "Bearer " + tokenString
		c.SetSameSite(http.SameSiteLaxMode)
		c.SetCookie("authorization", finalToken, 3600 * 4, "", "", false, true)

		c.JSON(http.StatusAccepted, gin.H {
			"id": id,
			"token": finalToken,
		})
		return
	}

	// Caso o usuário não esteja autenticado
	c.JSON(http.StatusForbidden, gin.H {
		"erro": "Credenciais incorretas",
	})
}

func Logout (c *gin.Context) {
	c.SetSameSite(http.SameSiteLaxMode)
	tokenCookie, err := c.Cookie("authorization")
	
	// Erro na hora de pegar o cookie
	if err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	tokenCookieSplit := strings.Split(tokenCookie, " ")

	// Erro no formato do cookie
	if len(tokenCookieSplit) < 2 {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	tokenString := tokenCookieSplit[1]

	// Registrando token como blacklisted
	// Recusa o pedido se o token já estiver na blacklist
	if database.CheckToken(tokenString) {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	} else {
		database.RegisterToken(tokenString)
	}

	c.JSON(http.StatusOK, gin.H {
		"sucesso": "logout realizado.",
	})
}
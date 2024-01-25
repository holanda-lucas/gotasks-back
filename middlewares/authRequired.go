package middlewares

import (
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/holanda-lucas/gotasks-back/database"
)

func AuthMiddleware(c *gin.Context) {
	authCookie, err := c.Cookie("authorization")

	// Tratando erro ao puxar o cookie
	if err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	authCookieSplit := strings.Split(authCookie, " ")

	// Tratando caso o header não possua o formato esperado
	if len(authCookieSplit) < 2 {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	tokenString := authCookieSplit[1]

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Checando o método de assinatura
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("método de assinatura inesperado: %v", token.Header["alg"])
		}

		return []byte(os.Getenv("SECRET_KEY")), nil
	})

	// Tratando para o caso de token inválido
	if err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	// Checando se o token está na blacklist
	t := database.GetToken(tokenString)
	if t.ID != 0 {
		c.AbortWithStatus(http.StatusForbidden)
	}

	// Checa os claims
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		// Checa se o token ainda não venceu
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			c.AbortWithStatus(http.StatusUnauthorized)
		}

		user := database.GetUser(claims["sub"])

		if user.ID == 0 {
			c.AbortWithStatus(http.StatusUnauthorized)
		}

		// Colocando o id no corpo da requisição
		c.Set("id", user.ID)

		c.Next()
	} else {
		c.AbortWithStatus(http.StatusUnauthorized)
	}
}
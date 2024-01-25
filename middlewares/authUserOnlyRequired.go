package middlewares

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func AuthUserOnlyMiddleware(c *gin.Context) {
	userId, exists := c.Get("id")

	if !exists {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	expectedId, err := strconv.Atoi(c.Param("id"))
	
	if err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	// Verifica se o ID da url é o mesmo id do usuário logado
	if uint(expectedId) != userId {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	c.Next()
}
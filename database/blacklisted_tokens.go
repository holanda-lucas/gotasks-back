package database

import (
	"time"

	"github.com/holanda-lucas/gotasks-back/models"
)

func RegisterToken(token string) {
	var t models.BlacklistedToken
	t.Token = token
	t.ExpirationDate = time.Now().Add(time.Hour * 4).Unix()

	DB.Create(&t)
}

func GetToken(token string) models.BlacklistedToken {
	var t models.BlacklistedToken
	DB.Where("Token = ?", token).First(&t)

	return t
}

func CheckToken(token string) bool {
	t := GetToken(token)

	return t.ID != 0
}

func DisposeExpiredTokens() {
	expiredTokens := getExpiredTokens()

	for _, token := range expiredTokens {
		deleteToken(token.ID)
	}
}

func deleteToken (id uint) {
	DB.Delete(&models.BlacklistedToken{}, id)
}

func getExpiredTokens () []models.BlacklistedToken {
	var t []models.BlacklistedToken
	DB.Where("expiration_date < ?", time.Now().Unix()).Find(&t)

	return t
}

package models

import (
	"gorm.io/gorm"
)

type BlacklistedToken struct {
	gorm.Model
	Token             string
	ExpirationDate    int64
}

// Identificando a tabela desse Model
func (BlacklistedToken) TableName() string {
	return "blacklisted_tokens"
}
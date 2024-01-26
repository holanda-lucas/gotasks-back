package models

type BlacklistedToken struct {
	ID                uint `gorm:"primaryKey"`
	Token             string
	ExpirationDate    int64
}

// Identificando a tabela desse Model
func (BlacklistedToken) TableName() string {
	return "blacklisted_tokens"
}
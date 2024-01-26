package models

import (
	"time"

	"github.com/go-playground/validator/v10"
)

type Task struct {
	ID                uint `gorm:"primaryKey"`
  	CreatedAt         time.Time `json:"created_at"`
  	UpdatedAt 		  time.Time `json:"updated_at"`
	User_id           uint `json:"user_id" validate:"required"`
	Title             string `json:"title" validate:"required"`
	Description       string `json:"description"`
	Tag               string `json:"tag" validate:"required"`
}

func (t *Task) ValidateTask() error {
	validate := validator.New()
	err := validate.Struct(t)

	return err
}

// Método que especifica a qual tabela o modelo pertence
func (Task) TableName() string {
    return "tasks"
}
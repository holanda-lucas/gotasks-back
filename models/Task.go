package models

import (
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type Task struct {
	gorm.Model
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
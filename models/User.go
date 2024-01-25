package models

import (
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email              string `json:"email" validate:"required,email" gorm:"unique"`
	Name               string `json:"name" validate:"required"`
    Password           string `json:"password" validate:"required"`
}

// Diferentes structs para diferentes possibilidades de dados
type UserJSON struct {
	gorm.Model
	Email              string `json:"email" validate:"required,email"`
	Name               string `json:"name" validate:"required"`
}

type UserLoginData struct {
	Email              string `json:"email" validate:"required,email"`
	Password           string `json:"password" validate:"required"`
}

// Métodos do User
func (u *User) CreateTask(title string, description string, tag string) *Task{
	newTask := Task{User_id: u.ID, Title: title, Description: description, Tag: tag}
	return &newTask
}

func (u *User) ValidateUser() error {
	validate := validator.New()
	err := validate.Struct(u)

	return err
}

// Métodos que especificam a qual tabela o modelo pertence
func (User) TableName() string {
    return "users"
}

func (UserJSON) TableName() string {
    return "users"
}

func (UserLoginData) TableName() string {
	return "users"
}
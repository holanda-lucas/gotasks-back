package database

import (
	"log"
	"time"

	"github.com/holanda-lucas/gotasks-back/models"
	"golang.org/x/crypto/bcrypt"
)

func RegisterUser(u *models.User) uint {
	err := u.ValidateUser()

	if err != nil {
		log.Fatal("Cannot register user: " + err.Error())
	}

	DB.Create(&u)

	return u.ID
}

func GetUser(id interface{}) models.UserJSON {
	var user models.UserJSON

	DB.First(&user, id)
	return user
}

func EditUser(u *models.User) {
	err := u.ValidateUser()

	if err != nil {
		log.Fatal("não foi possível editar o usuário: " + err.Error())
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)

	if err != nil {
		log.Fatal("não foi possível editar o usuário: " + err.Error())
	}
	
	u.Password = string(hashedPassword)
	u.UpdatedAt = time.Now()
	
	DB.Table("users").UpdateColumns(u)
}

func DeleteUser(id uint) {
	DB.Unscoped().Delete(&models.User{}, id)
}

func AuthenticateUser(loginData models.UserLoginData) uint {
	var user models.User
	result := DB.Where(&models.User{Email: loginData.Email}).First(&user)

	if result.Error != nil {
		return 0
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginData.Password))

	if err != nil {
		return 0
	}

	return user.ID
}
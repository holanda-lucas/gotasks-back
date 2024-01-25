package database

import (
	"log"

	"github.com/holanda-lucas/gotasks-back/models"
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
		log.Fatal("Cannot edit user: " + err.Error())
	}
	
	DB.Table("users").UpdateColumns(u)
}

func DeleteUser(id uint) {
	DB.Delete(&models.User{}, id)
}

func AuthenticateUser(loginData models.UserLoginData) uint {
	var user models.User
	result := DB.Where(&models.User{Email: loginData.Email, Password: loginData.Password}).First(&user)

	if result.Error != nil {
		return 0
	}

	return user.ID
}
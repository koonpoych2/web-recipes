package services

import (
	"github.com/koonpoych2/web-recipes/back-end/database"
	"github.com/koonpoych2/web-recipes/back-end/internal/models"
)

type UserInput struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func GetAllUsers() ([]models.User, error) {
	var users []models.User
	err := database.DB.Find(&users).Error
	return users, err
}

func CreateUser(input UserInput) (models.User, error) {
	user := models.User{Name: input.Name, Email: input.Email}
	err := database.DB.Create(&user).Error
	return user, err
}

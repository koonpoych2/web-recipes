package services

import (
	"your-project/database"
	"your-project/internal/models"
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

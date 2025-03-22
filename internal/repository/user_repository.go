package repository

import (
	"github.com/AmmarBerkovic/GoBuyExample/db"
	"github.com/AmmarBerkovic/GoBuyExample/internal/models"
)

func GetAllUsers() ([]models.User, error) {
	var users []models.User
	if err := db.DB.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func CreateUser(user models.User) (models.User, error) {
	if err := db.DB.Create(&user).Error; err != nil {
		return models.User{}, err
	}
	return user, nil
}

func UpdateUser(id string, user models.User) (models.User, error) {
	var existingUser models.User
	if err := db.DB.First(&existingUser, id).Error; err != nil {
		return models.User{}, err
	}

	existingUser.Name = user.Name
	existingUser.Email = user.Email
	existingUser.Age = user.Age

	if err := db.DB.Save(&existingUser).Error; err != nil {
		return models.User{}, err
	}
	return existingUser, nil
}

func DeleteUser(id string) error {
	var user models.User
	if err := db.DB.First(&user, id).Error; err != nil {
		return err
	}

	if err := db.DB.Delete(&user).Error; err != nil {
		return err
	}
	return nil
}

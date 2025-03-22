package services

import (
	"github.com/AmmarBerkovic/GoBuyExample/internal/models"
	"github.com/AmmarBerkovic/GoBuyExample/internal/repository"
)

func GetAllUsers() ([]models.User, error) {
	return repository.GetAllUsers()
}
func GetUsersWithPagination(page int, pageSize int) ([]models.User, error) {
	return repository.GetUsersWithPagination(page, pageSize)
}
func CreateUser(user models.User) (models.User, error) {
	return repository.CreateUser(user)
}

func UpdateUser(id string, user models.User) (models.User, error) {
	return repository.UpdateUser(id, user)
}

func DeleteUser(id string) error {
	return repository.DeleteUser(id)
}

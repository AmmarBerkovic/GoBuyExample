package services

import (
	"github.com/AmmarBerkovic/GoBuyExample/internal/pb"
	"github.com/AmmarBerkovic/GoBuyExample/internal/repository"
)

func GetAllUsers() ([]*pb.User, error) {
	return repository.GetAllUsers()
}

func GetUsersWithPagination(page int, pageSize int) ([]*pb.User, error) {
	return repository.GetUsersWithPagination(page, pageSize)
}

func CreateUser(user *pb.User) (*pb.User, error) {
	return repository.CreateUser(user)
}

func UpdateUser(id string, user *pb.User) (*pb.User, error) {
	return repository.UpdateUser(id, user)
}

func DeleteUser(id string) error {
	return repository.DeleteUser(id)
}

package repository

import (
	"github.com/AmmarBerkovic/GoBuyExample/db"
	"github.com/AmmarBerkovic/GoBuyExample/internal/models"
)

// Temporary struct for database operations (due to GORM struct tag requirements)
type UserDB struct {
	ID    uint `gorm:"primaryKey"`
	Name  string
	Email string
	Age   int32
}

// Convert UserDB to pb.User
func toProtoUser(user UserDB) *models.User {
	return &models.User{
		Id:    uint32(user.ID),
		Name:  user.Name,
		Email: user.Email,
		Age:   user.Age,
	}
}

// Convert pb.User to UserDB
func toDBUser(user *models.User) UserDB {
	return UserDB{
		ID:    uint(user.Id),
		Name:  user.Name,
		Email: user.Email,
		Age:   user.Age,
	}
}

func GetAllUsers() ([]*models.User, error) {
	var users []UserDB
	if err := db.DB.Find(&users).Error; err != nil {
		return nil, err
	}

	var protoUsers []*models.User
	for _, user := range users {
		protoUsers = append(protoUsers, toProtoUser(user))
	}
	return protoUsers, nil
}

func GetUsersWithPagination(page int, pageSize int) ([]*models.User, error) {
	var users []UserDB
	offset := (page - 1) * pageSize

	if err := db.DB.Limit(pageSize).Offset(offset).Find(&users).Error; err != nil {
		return nil, err
	}

	var protoUsers []*models.User
	for _, user := range users {
		protoUsers = append(protoUsers, toProtoUser(user))
	}
	return protoUsers, nil
}

func CreateUser(user *models.User) (*models.User, error) {
	dbUser := toDBUser(user)
	if err := db.DB.Create(&dbUser).Error; err != nil {
		return nil, err
	}
	return toProtoUser(dbUser), nil
}

func UpdateUser(id string, user *models.User) (*models.User, error) {
	var existingUser UserDB
	if err := db.DB.First(&existingUser, id).Error; err != nil {
		return nil, err
	}

	existingUser.Name = user.Name
	existingUser.Email = user.Email
	existingUser.Age = user.Age

	if err := db.DB.Save(&existingUser).Error; err != nil {
		return nil, err
	}
	return toProtoUser(existingUser), nil
}

func DeleteUser(id string) error {
	var user UserDB
	if err := db.DB.First(&user, id).Error; err != nil {
		return err
	}
	return db.DB.Delete(&user).Error
}

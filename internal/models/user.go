package models

// User represents the structure of the users table
type User struct {
	ID    uint   `gorm:"primaryKey" json:"id"`
	Name  string `gorm:"type:varchar(255);not null" json:"name"`
	Email string `gorm:"type:varchar(255);uniqueIndex;not null" json:"email"`
	Age   int    `json:"age"`
}

// TableName overrides the default table name
func (User) TableName() string {
	return "users"
}

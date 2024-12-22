package db

import (
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	Id        string         `gorm:"column:id;primaryKey;size:36"`
	Username  string         `gorm:"column:username;unique;not null;size:128"`
	Password  string         `gorm:"column:password;not null;size:255"`
	FullName  string         `gorm:"column:full_name;not null;size:128"`
	CreatedAt time.Time      `gorm:"column:created_at"`
	UpdatedAt time.Time      `gorm:"column:updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at;index"`
}

func (user *User) BeforeCreate(tx *gorm.DB) error {
	// Create UUID for user id
	genUuid, err := uuid.NewV7()
	if err != nil {
		return err
	}
	stringUuid := genUuid.String()
	user.Id = stringUuid

	// Hash user password
	hashedBytes, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	if err != nil {
		return err
	}
	user.Password = string(hashedBytes)
	return nil
}

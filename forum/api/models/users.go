package models

import (
	"html"
	"strings"

	"fmt"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email    string    `json:"email" gorm:"unique"`
	Username string    `json:"username" gorm:"unique"`
	Password string    `json:"password"`
	Post     []Post    `gorm:"foreignKey:UserID" json:"-"`
	Comment  []Comment `gorm:"foreignKey:UserID" json:"-"`
}

func (u *User) Prepare() {
	u.Email = html.EscapeString(strings.TrimSpace(u.Email))
	u.Username = html.EscapeString(strings.TrimSpace(u.Username))
}

func CreateUser(u *User, db *gorm.DB) error {
	u.Prepare()

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), 12)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)

	if err := db.Create(u).Error; err != nil {
		return err
	}

	return nil
}

func Login(db *gorm.DB, username, password string) (*User, error) {
	u, err := GetUser(db, username)
	if err != nil {
		return nil, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password)); err != nil {
		return nil, fmt.Errorf("incorrect password")
	}

	return u, nil
}

func DeleteUser(db *gorm.DB, userID uint, targetUserID uint) error {

	if userID != targetUserID {
		return fmt.Errorf("unauthorized: cannot delete another user")
	}
	result := db.Where("id = ?", targetUserID).Delete(&User{})

	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return fmt.Errorf("user not found")
	}

	return nil
}

func GetUser(db *gorm.DB, username string) (*User, error) {
	var user User

	if err := db.Where("username = ?", username).First(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func GetPostsByUserID(db *gorm.DB, userID uint) ([]Post, error) {
	var posts []Post
	err := db.Where("user_id = ?", userID).Find(&posts).Error
	return posts, err
}

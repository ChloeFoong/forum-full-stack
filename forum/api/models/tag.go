package models

import (
	"fmt"
	"html"
	"strings"

	"gorm.io/gorm"
)

type Tag struct {
	gorm.Model
	Name    string `json:"name"`
	TopicID uint   `json:"topic_id"`
	Topic   *Topic `gorm:"foreignKey:TopicID" json:"-"`
	UserID  uint   `json:"user_id"`
	User    *User  `gorm:"foreignKey:UserID" json:"-"`
	Post    []Post `gorm:"many2many:post_tags" json:"-"`
}

func (t *Tag) Prepare() {
	t.Name = html.EscapeString(strings.TrimSpace(t.Name))
}

func CreateTag(db *gorm.DB, t *Tag) error {
	t.Prepare()
	if err := db.Create(t).Error; err != nil {
		return err
	}
	return nil
}

func DeleteTag(db *gorm.DB, userID uint, t *Tag) error {
	if t.UserID != userID {
		return fmt.Errorf("unauthorized: cannot delete another user's post")
	}
	result := db.Where("id = ? AND user_id = ?", t.ID, userID).Delete(&Tag{})
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return fmt.Errorf("post not found or unauthorized")
	}

	return nil
}

func GetTag(db *gorm.DB, tagID uint) (*Tag, error) {
	var tag Tag
	if err := db.Where("id = ?", tagID).First(&tag).Error; err != nil {
		return nil, err
	}
	return &tag, nil

}

func GetAllTag(db *gorm.DB, postID uint) ([]Tag, error) {
	var tags []Tag
	if err := db.Where("post_id = ?", postID).Find(&tags).Error; err != nil {
		return nil, err
	}

	return tags, nil
}

package models

import (
	"fmt"

	"gorm.io/gorm"
)

type Comment struct {
	gorm.Model
	Content string `json:"content"`
	UserID  uint   `json:"user_id"`
	User    *User  `gorm:"foreignKey:UserID" json:"-"`
	PostID  uint   `json:"post_id"`
	Post    *Post  `gorm:"foreignKey:PostID" json:"-"`
}

func CreateComment(db *gorm.DB, userID uint, postID uint, content string) (*Comment, error) {
	comment := Comment{
		Content: content,
		UserID:  userID,
		PostID:  postID,
	}

	if err := db.Create(&comment).Error; err != nil {
		return nil, err
	}

	return &comment, nil
}

func UpdateComment(db *gorm.DB, userID, commentID uint, newContent string) (*Comment, error) {
	c, err := GetComment(db, commentID)
	if c.UserID != userID {
		return c, fmt.Errorf("unauthorized: cannot update another user's comment")
	}

	result := db.Model(c).Updates(Comment{Content: newContent})
	if err != nil {
		return &Comment{}, err
	}
	if result.Error != nil {
		return nil, result.Error
	}

	if result.RowsAffected == 0 {
		return &Comment{}, fmt.Errorf("post not found or unauthorized")
	}

	return c, nil
}

func DeleteComment(db *gorm.DB, userID uint, c *Comment) error {
	if c.UserID != userID {
		return fmt.Errorf("unauthorized: cannot delete another user's comment")
	}
	result := db.Where("id = ? AND user_id = ?", c.ID, userID).Delete(&Comment{})
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return fmt.Errorf("comment not found or unauthorized")
	}

	return nil
}

func GetComment(db *gorm.DB, commentID uint) (*Comment, error) {
	var comment Comment
	if err := db.Where("id = ?", commentID).First(&comment).Error; err != nil {
		return nil, err
	}
	return &comment, nil
}

func GetAllComment(db *gorm.DB, postID uint) ([]Comment, error) {
	var comments []Comment
	if err := db.Where("post_id = ?", postID).Find(&comments).Error; err != nil {
		return nil, err
	}

	return comments, nil
}

package models

import (
	"fmt"
	"html"
	"strings"

	"gorm.io/gorm"
)

type Post struct {
	gorm.Model
	Heading string    `json:"heading"`
	Content string    `json:"content"`
	UserID  uint      `json:"user_id"`
	User    *User     `gorm:"foreignKey:UserID" json:"-"`
	TopicID uint      `json:"topic_id"`
	Topic   *Topic    `gorm:"foreignKey:TopicID" json:"-"`
	Comment []Comment `gorm:"foreignKey:PostID" json:"-"`
	Tag     []Tag     `json:"tags" gorm:"many2many:post_tags;"`
}

func (p *Post) Prepare() {
	p.Heading = html.EscapeString(strings.TrimSpace(p.Heading))
	p.Content = html.EscapeString(strings.TrimSpace(p.Content))

}

func CreatePost(db *gorm.DB, p *Post) error {
	p.Prepare()
	if err := db.Create(p).Error; err != nil {
		return err
	}
	return nil
}

func UpdatePost(db *gorm.DB, userID uint, newHeading, newContent string, postID uint) (*Post, error) {
	p, err := GetPost(db, postID)
	if p.UserID != userID {
		return p, fmt.Errorf("unauthorized: cannot update another user's post")
	}

	result := db.Model(p).Updates(Post{Heading: newHeading, Content: newContent})
	if err != nil {
		return &Post{}, err
	}
	if result.Error != nil {
		return nil, result.Error
	}

	if result.RowsAffected == 0 {
		return &Post{}, fmt.Errorf("post not found or unauthorized")
	}

	return p, nil
}

func DeletePost(db *gorm.DB, userID uint, p *Post) error {
	if p.UserID != userID {
		return fmt.Errorf("unauthorized: cannot delete another user's post")
	}
	result := db.Where("id = ? AND user_id = ?", p.ID, userID).Delete(&Post{})
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return fmt.Errorf("post not found or unauthorized")
	}

	return nil
}

func GetPost(db *gorm.DB, postID uint) (*Post, error) {
	var post Post
	if err := db.Where("id = ?", postID).First(&post).Error; err != nil {
		return nil, err
	}
	return &post, nil

}

func GetAllPost(db *gorm.DB, topicID uint) ([]Post, error) {
	var posts []Post
	if err := db.Where("topic_id = ?", topicID).Preload("Comment").Find(&posts).Error; err != nil {
		return nil, err
	}

	return posts, nil
}

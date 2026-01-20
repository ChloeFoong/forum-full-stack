package models

import (
	"log"

	"gorm.io/gorm"
)

type Topic struct {
	gorm.Model
	Name string `json:"name" gorm:"unique"`
	Post []Post `gorm:"foreignKey:TopicID" json:"-"`
}

// initialise topics. Should only be done by admin and not reachable by clients

func SeedTopics(db *gorm.DB) {
	var count int64
	db.Model(&Topic{}).Count(&count)
	if count > 0 {
		return
	}

	topics := []Topic{
		{Name: "Technology and Computing"},
		{Name: "Gaming"},
		{Name: "Career and Education"},
		{Name: "Finance and Business"},
		{Name: "Lifestyles and Hobbies"},
		{Name: "Relationship and Personal Life"},
		{Name: "Entertainment"},
		{Name: "News and Society"},
		{Name: "Products and Review"},
	}

	if err := db.Create(&topics).Error; err != nil {
		log.Println("Failed to seed topics:", err)
	} else {
		log.Println("Topics seeded successfully")
	}

}

func GetAllTopics(db *gorm.DB) ([]Topic, error) {
	var topics []Topic

	if err := db.Find(&topics).Error; err != nil {
		return nil, err
	}

	return topics, nil
}

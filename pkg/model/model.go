package model

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"time"
)

type User struct {
	// ID of user
	Id int `gorm:"primary_key" json:"id"`
	// Name of user
	FirstName string `json:"firstName" binding:"required"`
	// Second Name of user
	SecondName string `json:"secondName" binding:"required"`
	// Username of user
	Username string `gorm:"unique" json:"username" binding:"required"`
	// HashPassword of user
	HashPassword string `json:"password" binding:"required"`
	// User creation date
	CreatedAt time.Time `json:"createdAt"`
	// User update date
	UpdatedAt time.Time `json:"updatedAt"`

	Segments []Segment `gorm:"many2many:user_segments;" json:"segments"`
}

// swagger:model Segment
type Segment struct {
	// ID of segment
	Id int `gorm:"primary_key" json:"id"`
	// Slug of segment
	Slug string `gorm:"unique" json:"slug" binding:"required"`
	// User creation date
	CreatedAt time.Time `json:"createdAt"`
	// User update date
	UpdatedAt time.Time `json:"updatedAt"`

	Users []User `gorm:"many2many:user_segments;" json:"users"`
}

type UserSegments struct {
	// List of segment slugs to add
	AddSlugs []string `json:"addSlugs" binding:"required"`
	// List of segment slugs to remove
	RemoveSlugs []string `json:"removeSlugs" binding:"required"`
	// ID of user
	UserId int `json:"userId" binding:"required"`
}

func DBMigrate(db *gorm.DB) *gorm.DB {
	db.AutoMigrate(&User{}, &Segment{})
	return db
}

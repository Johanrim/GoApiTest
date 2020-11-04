package Models

import (
	"time"

	"gorm.io/gorm"
)

type Book struct {
	ID   uint   `gorm:"primaryKey"`
	Name string `gorm:"not null" json:"name"`
	// Category string `gorm:"not null" json:"category"`
	CategoryID uint           `gzorm:"not null" json:"category"`
	Category   Category       `gorm:"ForeignKey:CategoryID;" json:"-"`
	CreatedAt  time.Time      `json:"created_at"`
	UpdatedAt  time.Time      `json:"updated_at"`
	DeletedAt  gorm.DeletedAt `gorm:"index"`
	// Author     Author         `gorm:"ForeignKey:AuthorID;"json:"-"`
	// AuthorID   uint           `gorm:"not null" json:"-"`
}

func (b *Book) TableName() string {
	return "book"
}

type Category struct {
	ID   uint   `gorm:"primaryKey" json:"id"`
	Name string `gorm:"not null" json:"name"`
}

func (c *Category) TableName() string {
	return "category"
}

// type Author struct {
// 	ID    uint   `gorm:"primaryKey" json:"id"`
// 	Name  string `gorm:"not null" json:"name"`
// 	Age   uint   `gorm:"not null" json:"age"`
// 	Phone string `gorm:"not null" json:"phone"`
// }

// func (a *Author) TableName() string {
// 	return "author"
// }

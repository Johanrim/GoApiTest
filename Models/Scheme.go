package Models

type Book struct {
	ID       uint   `gorm:"primaryKey"`
	Name     string `gorm:"not null" json:"name"`
	Category string `gorm:"not null"json:"category"`
	Author   string `gorm:"not null" json:"author"`
	// Author   Author `gorm:"ForeignKey:AuthorID;"json:"-"`
	// AuthorID uint   `gorm:"not null" json:"-"`
}

func (b *Book) TableName() string {
	return "book"
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

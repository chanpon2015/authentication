package model

// User is
type User struct {
	Gorm
	UserID   string `gorm:"unique;not null"`
	Password string `gorm:"not null"`
	Name     string `gorm:"not null"`
}

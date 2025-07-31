package entity

import (
	"time"

	"gorm.io/gorm"
)

type Fine struct {
	gorm.Model

	FineID uint `json:"fine_id" gorm:"primaryKey"`

	FineDate time.Time `json:"fine_date"`

	LastName string `json:"last_name"`

	Email string `json:"email"`

	PhoneNumber string `json:"phonenumber"`
	
	BorrowID uint `json:"borrow_id"`
	BorrowRecord BorrowRecord `gorm:"foreignKey:BorrowID"`
}

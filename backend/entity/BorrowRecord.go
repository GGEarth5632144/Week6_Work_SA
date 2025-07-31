package entity

import (
	"time"

	"gorm.io/gorm"
)

type BorrowRecord struct {
	gorm.Model

	BorrowID uint `json:"borrow_id" gorm:"primaryKey"`

	BorrowDate time.Time `json:"borrow_date"`

	DueDate time.Time `json:"due_date"`

	ReturnDate time.Time `json:"return_date"`

	BookID   uint `json:"book_id"`
	MemberID uint `json:"member_id"`

	Book            Book            `gorm:"foreignKey:BookID"`
	LibrarianMember LibrarianMember `gorm:"foreignKey:MemberID"`
}

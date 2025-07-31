package entity

import (
	"gorm.io/gorm"
)

type LibrarianMember struct {
	
	gorm.Model

	MemberID uint `json:"member_id" gorm:"primaryKey"`

	FirstName string `json:"first_name"`

	LastName string `json:"last_name"`

	Email string `json:"email"`

	PhoneNumber string `json:"phonenumber"`
}

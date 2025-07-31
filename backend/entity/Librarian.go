package entity


import (

   "gorm.io/gorm"

)

type Librarian struct {

   gorm.Model

   LibrarianID uint    `json:"librarian_id" gorm:"primaryKey"`

   FirstName string    `json:"first_name"`

   LastName  string    `json:"last_name"`

   Email     string    `json:"email"`

   PhoneNumber string  `json:"phonenumber"`

}
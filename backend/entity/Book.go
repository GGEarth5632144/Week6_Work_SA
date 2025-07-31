package entity


import (

   "gorm.io/gorm"

)

type Book struct {

   gorm.Model

   BookID uint      `json:"book_id" gorm:"primaryKey"`

   Title string    `json:"title"`

   Author string    `json:"author"`

   Publisher string    `json:"publisher"`

   PublicationYear int    `json:"publication_year"`

   Caregory string    `json:"category"`

   BookStatus string    `json:"book_status"`

}
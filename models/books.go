package models

import "gorm.io/gorm"

type Book struct{
    gorm.Model
    ID      int `gorm:"primaryKey" json:"id"`
    Title    string `json:"title"`
    Author    string `json:"author"`
    Publisher   string `json:"publisher"`
    City   string `json:"city"`
    Year   string `json:"year"`
    Amount   int `json:"amount"`
}
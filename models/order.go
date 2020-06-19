package models

import "github.com/jinzhu/gorm"

type Order struct {
	gorm.Model
	Province 	string   `gorm:"size:255" json:"province"`
	Area     	string   `gorm:"size:255" json:"area"`
	Shop		string   `gorm:"size:255" json:"shop"`
	PhoneType	string 	 `gorm:"size:255" json:"phone_type"`
	Name        string 	 `gorm:"size:255" json:"name"`
	Sex         string 	 `gorm:"size:1" json:"sex"`
	Phone       string   `gorm:"size:20" json:"phone"`
	Birthday    string   `gorm:"size:20" json:"birthday"`
	Nation      string   `gorm:"size:100" json:"nation"`
}


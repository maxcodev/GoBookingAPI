package models

import "gorm.io/gorm"

type Hotel struct {
	gorm.Model
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Outstanding bool    `gorm:"default:false" json:"outstanding"`
	Rooms       []Rooms `gorm:"foreignKey:HotelID" json:"rooms"`
}

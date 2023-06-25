package hotel

import (
	"gorm.io/gorm"
)

type Hotel struct {
	gorm.Model
	Name        string       `json:"name"`
	Description string       `json:"description"`
	Outstanding bool         `gorm:"default:false" json:"outstanding"`
	Rooms       []Rooms      `gorm:"foreignKey:HotelID" json:"rooms"`
	Pictures    []Pictures   `gorm:"foreignKey:HotelID" json:"pictures"`
	Categories  []Categories `gorm:"foreignKey:HotelID" json:"categories"`
	Types       []Type       `gorm:"foreignKey:HotelID" json:"types"`
	Services    []Services   `gorm:"foreignKey:HotelID" json:"services"`
}

package hotel

import "gorm.io/gorm"

type Type struct {
	gorm.Model
	Type    string `json:"type"`
	HotelID uint   `json:"hotel_id"`
}

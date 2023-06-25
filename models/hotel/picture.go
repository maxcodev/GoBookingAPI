package hotel

import "gorm.io/gorm"

type Pictures struct {
	gorm.Model
	Name        string `json:"picture_name"`
	PictureJson string `json:"picture_json"`
	HotelID     uint   `json:"hotel_id"`
}

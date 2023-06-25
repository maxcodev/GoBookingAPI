package hotel

import "gorm.io/gorm"

type Services struct {
	gorm.Model
	ServiceJson string `json:"service_json"`
	HotelID     uint   `json:"hotel_id"`
}

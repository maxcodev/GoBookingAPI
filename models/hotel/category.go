package hotel

import "gorm.io/gorm"

type Categories struct {
	gorm.Model
	Rating  int16 `json:"name"`
	HotelID uint  `json:"hotel_id"`
}

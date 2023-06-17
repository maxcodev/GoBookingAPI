package models

import "gorm.io/gorm"

type Rooms struct {
	gorm.Model
	RoomNumber  uint16 `json:"room_number"`
	Description string `json:"description"`
	HotelID     uint   `json:"hotel_id"`
}

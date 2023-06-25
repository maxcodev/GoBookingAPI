package service

import (
	"github.com/gorilla/mux"
	"github.com/maxcodev/booking_ws_api/models/hotel"
	"net/http"
	"time"

	"github.com/maxcodev/booking_ws_api/repository"
)

func GetHotels(w http.ResponseWriter, r *http.Request) {
	var hotels []hotel.Hotel
	repository.GetHotels(w, &hotels)
}

func GetHotelById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var hotel hotel.Hotel
	repository.GetHotelById(w, r, &hotel, params)
}

func CreateHotel(w http.ResponseWriter, r *http.Request, newHotel *hotel.Hotel) {
	repository.CreateHotel(w, r, newHotel)
}

func UpdateHotelService(w http.ResponseWriter, r *http.Request, updatedHotel *hotel.Hotel) {
	var upHotel hotel.Hotel
	upHotel = repository.GetRelated(updatedHotel.ID)
	upHotel.Name = updatedHotel.Name
	upHotel.Outstanding = updatedHotel.Outstanding
	upHotel.Description = updatedHotel.Description
	upHotel.UpdatedAt = time.Now()
	repository.UpdateHotelRepo(w, r, &upHotel)
}

func DeleteHotel(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var hotel hotel.Hotel
	repository.DeleteHotel(w, r, &hotel, params)
}

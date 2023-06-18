package service

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/maxcodev/booking_ws_api/models"
	"github.com/maxcodev/booking_ws_api/repository"
)

func GetHotels(w http.ResponseWriter, r *http.Request) {
	var hotels []models.Hotel
	repository.GetHotels(w, &hotels)
}

func GetHotelById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var hotel models.Hotel
	repository.GetHotelById(w, r, &hotel, params)
}

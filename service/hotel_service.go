package service

import (
	"encoding/json"
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

func CreateHotel(w http.ResponseWriter, r *http.Request) {
	var newHotel models.Hotel
	json.NewDecoder(r.Body).Decode(&newHotel)
	repository.CreateHotel(w, r, &newHotel)
}

func DeleteHotel(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var hotel models.Hotel
	repository.DeleteHotel(w, r, &hotel, params)
}

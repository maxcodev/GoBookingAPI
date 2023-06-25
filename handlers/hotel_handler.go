package handlers

import (
	"encoding/json"
	"github.com/maxcodev/booking_ws_api/models/hotel"
	"net/http"

	"github.com/maxcodev/booking_ws_api/service"
)

func GetHotel(w http.ResponseWriter, r *http.Request) {
	defer func() {
		if err := recover(); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Internal server error"))
		}
	}()
	service.GetHotels(w, r)
}

func GetHotelById(w http.ResponseWriter, r *http.Request) {
	defer func() {
		if err := recover(); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Internal server error"))
		}
	}()

	service.GetHotelById(w, r)
}

func CreateHotel(w http.ResponseWriter, r *http.Request) {
	var newHotel hotel.Hotel
	err := json.NewDecoder(r.Body).Decode(&newHotel)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Invalid request	payload"))
		return
	}

	if newHotel.Name == "" || newHotel.Description == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Name and description is required"))
		return
	}

	service.CreateHotel(w, r, &newHotel)
}

func UpdateHotel(w http.ResponseWriter, r *http.Request) {
	var updateHotel hotel.Hotel
	err := json.NewDecoder(r.Body).Decode(&updateHotel)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Invalid request	payload"))
		return
	}

	if updateHotel.Name == "" || updateHotel.Description == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Name and description is required"))
		return
	}

	service.UpdateHotelService(w, r, &updateHotel)
}

func DeleteHotelHandler(w http.ResponseWriter, r *http.Request) {
	defer func() {
		if err := recover(); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Internal server error"))
		}
	}()
	service.DeleteHotel(w, r)
}

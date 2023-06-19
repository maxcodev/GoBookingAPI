package handlers

import (
	"encoding/json"
	"net/http"

	md "github.com/maxcodev/booking_ws_api/models"
	"github.com/maxcodev/booking_ws_api/service"
)

func GetHotelHandler(w http.ResponseWriter, r *http.Request) {
	defer func() {
		if err := recover(); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Internal server error"))
		}
	}()
	service.GetHotels(w, r)
}

func GetHotelByIdHandler(w http.ResponseWriter, r *http.Request) {
	defer func() {
		if err := recover(); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Internal server error"))
		}
	}()

	service.GetHotelById(w, r)
}

func CreateHotel(w http.ResponseWriter, r *http.Request) {
	var newHotel md.Hotel

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

	service.CreateHotel(w, r)
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

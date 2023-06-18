package handlers

import (
	"net/http"

	"github.com/maxcodev/booking_ws_api/service"
)

func GetHotelHandler(w http.ResponseWriter, r *http.Request) {
	service.GetHotels(w, r)
}

func GetHotelByIdHandler(w http.ResponseWriter, r *http.Request) {
	service.GetHotelById(w, r)
}

package handlers

import (
	"github.com/maxcodev/booking_ws_api/service"
	"net/http"
)

func GetHotelHandler(w http.ResponseWriter, r *http.Request) {
	service.GetHotelHandler(r.Context(), w, r)
}

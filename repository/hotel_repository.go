package repository

import (
	"encoding/json"
	"net/http"

	"github.com/maxcodev/booking_ws_api/database"
	"github.com/maxcodev/booking_ws_api/models"
)

func GetHotels(w http.ResponseWriter, hotels *[]models.Hotel) {
	//Obtener todos los hoteles de la base de datos
	database.DB.Find(&hotels)

	// Codificar los hoteles en formato JSON y enviar la respuesta al cliente
	json.NewEncoder(w).Encode(&hotels)
}

func GetHotelById(w http.ResponseWriter, r *http.Request, hotel *models.Hotel, params map[string]string) {
	//Obtener todos los hoteles de la base de datos
	id := params["id"]
	database.DB.First(&hotel, id)
	if hotel.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Hotel not found"))
		return
	}
	// Codificar los hoteles en formato JSON y enviar la respuesta al cliente
	json.NewEncoder(w).Encode(&hotel)
}

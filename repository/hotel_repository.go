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

func CreateHotel(w http.ResponseWriter, r *http.Request, hotel *models.Hotel) {
	err := database.DB.Create(hotel).Error //Crear el nuevo hotel en la BD
	if err != nil {                        //Si hay un error
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal server error"))
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("New hotel has been added succefully!"))
}

func DeleteHotel(w http.ResponseWriter, r *http.Request, hotel *models.Hotel, params map[string]string) {
	id := params["id"]
	database.DB.First(&hotel, id)

	if hotel.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Hotel not found"))
		return
	}

	//database.DB.Unscoped().Delete(&hotel) // Unscoped permite borrado fisico
	database.DB.Delete(&hotel) // Borrado lógico
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Hotel deleted"))
}

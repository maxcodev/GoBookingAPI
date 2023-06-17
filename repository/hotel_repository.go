package repository

import (
	"context"
	"encoding/json"
	"github.com/maxcodev/booking_ws_api/database"
	"github.com/maxcodev/booking_ws_api/models"
	"net/http"
)

func GetHotels(ctx context.Context, w http.ResponseWriter) {
	//Obtener todos los hoteles de la base de datos
	var hotels []models.Hotel
	database.DB.Find(&hotels)

	// Codificar los hoteles en formato JSON y enviar la respuesta al cliente
	json.NewEncoder(w).Encode(&hotels)
}

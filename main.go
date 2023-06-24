package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/maxcodev/booking_ws_api/database"
	"github.com/maxcodev/booking_ws_api/handlers"
	"github.com/maxcodev/booking_ws_api/models"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	database.DBConnection()

	database.DB.AutoMigrate(models.Hotel{})
	database.DB.AutoMigrate(models.Rooms{})

	r := mux.NewRouter()

	// Index route
	r.HandleFunc("/", handlers.HomeHandler)

	s := r.PathPrefix("/api").Subrouter()

	// Hotel routes
	s.HandleFunc("/hotels", handlers.GetHotel).Methods(http.MethodGet)
	s.HandleFunc("/hotel/{id}", handlers.GetHotelById).Methods(http.MethodGet)
	s.HandleFunc("/addHotel", handlers.CreateHotel).Methods(http.MethodPost)
	s.HandleFunc("/updateHotel", handlers.UpdateHotel).Methods(http.MethodPut)
	s.HandleFunc("/deleteHotel/{id}", handlers.DeleteHotelHandler).Methods(http.MethodDelete)

	http.ListenAndServe(":4000", r)
}

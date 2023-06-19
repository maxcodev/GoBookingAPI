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
	s.HandleFunc("/hotels", handlers.GetHotelHandler).Methods("GET")
	s.HandleFunc("/hotel/{id}", handlers.GetHotelByIdHandler).Methods("GET")
	s.HandleFunc("/addHotel", handlers.CreateHotel).Methods("POST")
	s.HandleFunc("/deleteHotel/{id}", handlers.DeleteHotelHandler).Methods("DELETE")

	http.ListenAndServe(":4000", r)
}

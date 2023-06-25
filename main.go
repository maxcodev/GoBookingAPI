package main

import (
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/maxcodev/booking_ws_api/database"
	"github.com/maxcodev/booking_ws_api/handlers"
	"github.com/maxcodev/booking_ws_api/models/hotel"
	"log"
	"net/http"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	database.DBConnection()

	database.DB.AutoMigrate(hotel.Hotel{})
	database.DB.AutoMigrate(hotel.Rooms{})
	database.DB.AutoMigrate(hotel.Pictures{})
	database.DB.AutoMigrate(hotel.Categories{})
	database.DB.AutoMigrate(hotel.Type{})
	database.DB.AutoMigrate(hotel.Services{})

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

package service

import (
	"context"
	"github.com/maxcodev/booking_ws_api/repository"
	"net/http"
)

func GetHotelHandler(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	repository.GetHotels(ctx, w)
}

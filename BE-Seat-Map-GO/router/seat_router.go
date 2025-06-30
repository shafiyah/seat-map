package router

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"github.com/shafiyah/seat-map/BE-Seat-Map-Go/featurs/flight"
	"github.com/shafiyah/seat-map/BE-Seat-Map-Go/featurs/seat"
	"github.com/shafiyah/seat-map/BE-Seat-Map-Go/handler"
)

func RegisterSeatRoutes(r *gin.Engine, db *gorm.DB) {
	flightRepo := flight.FlightRepository(db)
	flightService := flight.FlightService(flightRepo)
	seatRepo := seat.SeatRepository(db)
	seatService := seat.SeatService(seatRepo)
	seatHandler := handler.SeatHandler(seatService, flightService)

	r.GET("/flights/:id/seats", seatHandler.GetSeatMap)
}

package router

import (
	"github.com/gin-gonic/gin"
	"github.com/shafiyah/seat-map/BE-Seat-Map-Go/featurs/flight"
	"github.com/shafiyah/seat-map/BE-Seat-Map-Go/handler"
	"gorm.io/gorm"
)

func RegisterFlightRouters(r *gin.Engine, db *gorm.DB) {
	flightRepo := flight.FlightRepository(db)
	flightService := flight.FlightService(flightRepo)
	flightHandler := handler.FlightHandler(flightService)

	r.POST("/flights/import-data", flightHandler.ImportDataFromJson)
}

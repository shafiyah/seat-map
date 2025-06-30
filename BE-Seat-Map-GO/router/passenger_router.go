package router

import (
	"github.com/gin-gonic/gin"
	"github.com/shafiyah/seat-map/BE-Seat-Map-Go/featurs/passenger"
	"github.com/shafiyah/seat-map/BE-Seat-Map-Go/handler"
	"gorm.io/gorm"
)

func RegisterPassengerRouter(r *gin.Engine, db *gorm.DB) {
	passengerRepository := passenger.PassangerRepository(db)
	passengerService := passenger.PassengerService(passengerRepository)
	passengerHandler := handler.PassengerHandler(passengerService)

	r.GET("/passenger/:id", passengerHandler.GetPassangerById)

}

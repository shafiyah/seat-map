package router

import (
	"github.com/gin-gonic/gin"
	"github.com/shafiyah/seat-map/BE-Seat-Map-Go/featurs/seat"
	"github.com/shafiyah/seat-map/BE-Seat-Map-Go/featurs/selection"
	"github.com/shafiyah/seat-map/BE-Seat-Map-Go/handler"
	"gorm.io/gorm"
)

func RegisterSelectionRouter(r *gin.Engine, db *gorm.DB) {
	seatRepo := seat.SeatRepository(db)
	seatService := seat.SeatService(seatRepo)

	selectionRepo := selection.SelectionRepository(db)
	selectionService := selection.SelectionService(selectionRepo, seatService)
	selectionHandler := handler.SelectionHandler(selectionService)

	r.POST("/seat/selection", selectionHandler.SeadSelection)

}

package router

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func NewRouter(db *gorm.DB) *gin.Engine {
	r := gin.Default()
	r.Use(cors.Default())

	RegisterSeatRoutes(r, db)
	RegisterPassengerRouter(r, db)
	RegisterFlightRouters(r, db)
	RegisterSelectionRouter(r, db)

	return r
}

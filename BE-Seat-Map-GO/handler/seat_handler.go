package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/shafiyah/seat-map/BE-Seat-Map-Go/featurs/flight"
	"github.com/shafiyah/seat-map/BE-Seat-Map-Go/featurs/seat"
)

type seatHandler struct {
	seatService   seat.SeatServiceInf
	flightService flight.FlightServiceInf
}

func SeatHandler(service seat.SeatServiceInf, flightService flight.FlightServiceInf) *seatHandler {
	return &seatHandler{
		seatService:   service,
		flightService: flightService}
}

func (h *seatHandler) GetSeatMap(c *gin.Context) {
	flightIDStr := c.Param("id")
	flightID, err := strconv.Atoi(flightIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid flight ID"})
		return
	}

	seatMap, err := h.seatService.GetSeatByFlightID(uint(flightID))

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to retrieve seat map"})
		return
	}

	flightDetail, err := h.flightService.GetFlightById(uint(flightID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to get flight"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"seatMaps":     seatMap,
		"flightDetail": flightDetail,
	})

}

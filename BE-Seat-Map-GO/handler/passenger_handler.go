package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/shafiyah/seat-map/BE-Seat-Map-Go/featurs/passenger"
)

type passengerHandler struct {
	passengerService passenger.PassengerServiceInf
}

func PassengerHandler(passengerService passenger.PassengerServiceInf) *passengerHandler {
	return &passengerHandler{passengerService: passengerService}
}

func (h *passengerHandler) GetPassangerById(c *gin.Context) {
	passengerIdStr := c.Param("id")
	passengerID, err := strconv.Atoi(passengerIdStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid flight ID"})
		return
	}

	passenger, err := h.passengerService.GetPassangerById(uint(passengerID))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error ": "failded to found user"})
	}

	c.JSON(http.StatusOK, passenger)

}

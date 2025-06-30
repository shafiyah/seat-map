package handler

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/shafiyah/seat-map/BE-Seat-Map-Go/featurs/flight"
)

type flightHandler struct {
	flightService flight.FlightServiceInf
}

func FlightHandler(flightService flight.FlightRepositoryInf) *flightHandler {
	return &flightHandler{flightService}
}

func (h *flightHandler) ImportDataFromJson(c *gin.Context) {

	jsonFile, err := os.Open("../utils/SeatMapResponse.json")

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed open file JSON:"})
		fmt.Println(err)
		return
	}
	defer jsonFile.Close()

	byteValue, _ := io.ReadAll(jsonFile)
	var jsonData map[string]interface{}
	if err := json.Unmarshal(byteValue, &jsonData); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed parse JSON:"})
		return
	}

	if err = h.flightService.ImportAndMappingFile(jsonData); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed import data "})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "success import data",
	})

}

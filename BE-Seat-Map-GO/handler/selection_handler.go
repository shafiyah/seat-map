package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/shafiyah/seat-map/BE-Seat-Map-Go/featurs/selection"
)

type selectionHandler struct {
	selectionService selection.SelecetionServiceInfc
}

func SelectionHandler(selectionService selection.SelecetionServiceInfc) *selectionHandler {
	return &selectionHandler{selectionService: selectionService}
}

func (h *selectionHandler) SeadSelection(c *gin.Context) {

	var seatSelection selection.SeatSelection
	if err := c.ShouldBindJSON(&seatSelection); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "invalid request body"})
		return
	}

	err := h.selectionService.SeatSelection(seatSelection)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed seat seat"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "success selected seat",
	})
}

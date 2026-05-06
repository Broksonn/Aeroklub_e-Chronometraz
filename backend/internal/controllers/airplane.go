package controllers

import (
	"net/http"

	"backend/internal/database"
	"backend/internal/models"

	"github.com/gin-gonic/gin"
)

func GetAirplanes(c *gin.Context) {
	var airplanes []models.Airplane
	database.DB.Find(&airplanes)

	var prices []models.Price
	database.DB.Find(&prices)

	// Map prices by model type
	priceMap := make(map[string]float64)
	for _, p := range prices {
		priceMap[p.Type] = p.PricePerMinute
	}

	// Create combined response
	type AirplaneResponse struct {
		ID             uint    `json:"id"`
		Model          string  `json:"model"`
		Registration   string  `json:"registration"`
		Status         string  `json:"status"`
		PricePerMinute float64 `json:"pricePerMinute"`
	}

	var response []AirplaneResponse
	for _, a := range airplanes {
		response = append(response, AirplaneResponse{
			ID:             a.ID,
			Model:          a.Model,
			Registration:   a.Registration,
			Status:         a.Status,
			PricePerMinute: priceMap[a.Model],
		})
	}

	c.JSON(http.StatusOK, response)
}

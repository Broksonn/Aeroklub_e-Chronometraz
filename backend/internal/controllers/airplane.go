package controllers

import (
	"net/http"

	"backend/internal/database"
	"backend/internal/models"

	"github.com/gin-gonic/gin"
)

// GetAirplanes pobiera listę wszystkich samolotów
func GetAirplanes(c *gin.Context) {
	var airplanes []models.Airplane
	database.DB.Find(&airplanes)

	var prices []models.Price
	database.DB.Find(&prices)

	// Mapujemy ceny po modelu
	priceMap := make(map[string]float64)
	for _, p := range prices {
		priceMap[p.Type] = p.PricePerMinute
	}

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

// CreateAirplane dodaje nowy samolot
func CreateAirplane(c *gin.Context) {
	var input struct {
		Model        string `json:"model" binding:"required"`
		Registration string `json:"registration" binding:"required"`
		Status       string `json:"status" binding:"required"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Wszystkie pola są wymagane"})
		return
	}

	airplane := models.Airplane{
		Model:        input.Model,
		Registration: input.Registration,
		Status:       input.Status,
		PriceID:      1, // Domyślna cena
	}

	if err := database.DB.Create(&airplane).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Nie udało się dodać samolotu"})
		return
	}

	c.JSON(http.StatusCreated, airplane)
}

// DeleteAirplane usuwa samolot
func DeleteAirplane(c *gin.Context) {
	id := c.Param("id")

	if err := database.DB.Delete(&models.Airplane{}, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Błąd podczas usuwania samolotu"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Samolot usunięty"})
}

// UpdateAirplaneStatus aktualizuje samo pole statusu samolotu
func UpdateAirplaneStatus(c *gin.Context) {
	id := c.Param("id")
	
	var input struct {
		Status string `json:"status" binding:"required,oneof=available maintenance"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Nieprawidłowy status. Dozwolone: available, maintenance"})
		return
	}

	result := database.DB.Model(&models.Airplane{}).Where("id = ?", id).Update("status", input.Status)
	
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Błąd bazy danych podczas aktualizacji statusu"})
		return
	}

	// Jeśli nie zmieniono żadnego wiersza, oznacza to, że samolotu o takim ID nie ma w bazie
	if result.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Nie znaleziono samolotu o podanym ID"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Status zaktualizowany pomyślnie"})
}
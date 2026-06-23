package controllers

import (
	"fmt"
	"net/http"
	"time"

	"backend/internal/database"
	"backend/internal/models"

	"github.com/gin-gonic/gin"
)

// GetReservations zwraca listę wszystkich rezerwacji z bazy
func GetReservations(c *gin.Context) {
	var reservations []models.Reservation
	database.DB.Preload("User").Find(&reservations)

	c.JSON(http.StatusOK, reservations)
}

// CreateReservation tworzy nową rezerwację, zapobiegając konfliktom czasowym i weryfikując wyprzedzenie
func CreateReservation(c *gin.Context) {
	var input struct {
		AirplaneID uint   `json:"airplane_id" binding:"required"`
		Date       string `json:"date" binding:"required"`
		StartTime  string `json:"start_time" binding:"required"`
		EndTime    string `json:"end_time" binding:"required"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Nieprawidłowe dane rezerwacji"})
		return
	}

	// 1. WALIDACJA CZASU (Min. 30 minut wyprzedzenia i brak rezerwacji w przeszłości)
	layout := "2006-01-02 15:04"
	dateTimeStr := fmt.Sprintf("%s %s", input.Date, input.StartTime)

	// Parsowanie czasu startu w lokalnej strefie czasowej serwera
	reservationStartTime, err := time.ParseInLocation(layout, dateTimeStr, time.Local)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Nieprawidłowy format daty lub czasu"})
		return
	}

	// Obliczamy ile czasu zostało do startu
	timeUntilFlight := time.Until(reservationStartTime)
	if timeUntilFlight < 30*time.Minute {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Rezerwację należy złożyć z co najmniej 30-minutowym wyprzedzeniem. Rezerwacje w przeszłości są zabronione."})
		return
	}

	// 2. BLOKADA NAKŁADANIA SIĘ CZASU (Overlapping)
	var count int64
	database.DB.Model(&models.Reservation{}).
		Where("airplane_id = ? AND date = ? AND start_time < ? AND end_time > ?",
			input.AirplaneID, input.Date, input.EndTime, input.StartTime).
		Count(&count)

	if count > 0 {
		c.JSON(http.StatusConflict, gin.H{"error": "Samolot jest już zajęty w wybranych godzinach!"})
		return
	}

	userIDClaim, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Brak autoryzacji"})
		return
	}

	userID := uint(userIDClaim.(float64))

	reservation := models.Reservation{
		UserID:     userID,
		AirplaneID: input.AirplaneID,
		Date:       input.Date,
		StartTime:  input.StartTime,
		EndTime:    input.EndTime,
	}

	if err := database.DB.Create(&reservation).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Błąd podczas zapisywania rezerwacji"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Rezerwacja zakończona sukcesem!"})
}

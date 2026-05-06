package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Zaślepki dla kalendarza (później dopiszesz logikę rezerwacji)
func GetReservations(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Тут будет список бронирований для календаря"})
}

func CreateReservation(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Тут будет создание новой брони"})
}

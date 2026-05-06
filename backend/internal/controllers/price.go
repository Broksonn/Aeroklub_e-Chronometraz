package controllers

import (
	"net/http"

	"backend/internal/database"
	"backend/internal/models"

	"github.com/gin-gonic/gin"
)

func GetPrices(c *gin.Context) {
	var prices []models.Price
	database.DB.Find(&prices)
	c.JSON(http.StatusOK, prices)
}

func UpdatePrices(c *gin.Context) {
	var inputPrices []models.Price
	if err := c.BindJSON(&inputPrices); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный формат данных"})
		return
	}

	// Usuwamy stare ceny i zapisujemy nowe (proste rozwiązanie dla cennika)
	database.DB.Exec("DELETE FROM prices")
	database.DB.Create(&inputPrices)

	c.JSON(http.StatusOK, gin.H{"message": "Цены успешно обновлены", "pricing": inputPrices})
}

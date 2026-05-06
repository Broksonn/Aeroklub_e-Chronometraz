package models

type Price struct {
	ID             uint    `json:"id" gorm:"primaryKey"`
	Type           string  `json:"type"`
	PricePerMinute float64 `json:"pricePerMinute"`
}

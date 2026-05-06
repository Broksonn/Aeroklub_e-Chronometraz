package models

type Airplane struct {
	ID           uint   `json:"id" gorm:"primaryKey"`
	Model        string `json:"model"`
	Registration string `json:"registration"`
	Status       string `json:"status"` // 'available', 'maintenance'
}

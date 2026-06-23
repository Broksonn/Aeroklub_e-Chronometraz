package models

type Reservation struct {
	ID         uint     `json:"id" gorm:"primaryKey"`
	UserID     uint     `json:"user_id"`
	AirplaneID uint     `json:"airplane_id"`
	Date       string   `json:"date"`       // Format: YYYY-MM-DD
	StartTime  string   `json:"start_time"` // Format: HH:mm (np. 08:30)
	EndTime    string   `json:"end_time"`   // Format: HH:mm (np. 10:00)
	User       User     `json:"user" gorm:"foreignKey:UserID"`
	Airplane   Airplane `json:"airplane" gorm:"foreignKey:AirplaneID"`
}

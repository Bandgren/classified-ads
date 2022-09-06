package entities

import "time"

type Ad struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Email     string    `json:"email"`
	Subject   string    `json:"subject"`
	Price     float64   `json:"price"`
	Body      string    `json:"body"`
	CreatedAt time.Time `json:"created_at"`
}

type AdPayload struct {
	Email   string  `json:"email"`
	Subject string  `json:"subject"`
	Price   float64 `json:"price"`
	Body    string  `json:"body"`
}

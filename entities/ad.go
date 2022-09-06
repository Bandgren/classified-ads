package entities

type Ad struct {
	ID      uint    `json:"id"`
	Email   string  `json:"email"`
	Subject string  `json:"subject"`
	Price   float64 `json:"price"`
	Body    string  `json:"body"`
}

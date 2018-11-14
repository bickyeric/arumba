package model

// Page ...
type Page struct {
	ID   int    `json:"id"`
	Link string `json:"link"`
}

// Episode merepresentasikan objek episode
type Episode struct {
	Page []Page `json:"page"`
}

package models

type Album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"Artist"`
	Price  float64 `json:"price"`
}

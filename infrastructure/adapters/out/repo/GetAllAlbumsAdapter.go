package repo

type Album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"Artist"`
	Price  float64 `json:"price"`
}

func InMemoryDB() []Album {
	return []Album{
		{ID: "1", Title: "Nevermind", Artist: "Nirvana", Price: 19.99},
		{ID: "2", Title: "Bleach", Artist: "Nirvana", Price: 9.99},
		{ID: "3", Title: "In Bloom", Artist: "Nirvana", Price: 1.99},
	}
}

func (a Album) GetAllAlbumsData() []Album {
	albums := InMemoryDB()
	return albums
}

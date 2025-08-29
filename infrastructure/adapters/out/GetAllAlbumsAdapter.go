package out

type AllAlbums struct {
	Albums []Album	
}

func InMemoryDB() AllAlbums {
	return AllAlbums{
		Albums: []Album{
			{ID: "1", Title: "Nevermind", Artist: "Nirvana", Price: 19.99},
			{ID: "2", Title: "Bleach", Artist: "Nirvana", Price: 9.99},
			{ID: "3", Title: "In Bloom", Artist: "Nirvana", Price: 1.99}
		}
	}
} 

func (a AllAlbums) GetAllAlbumsData() []Album  {
	albums := a.Albums
	return albums;
}
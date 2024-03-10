package entities

type GenreSimple struct {
	ID   int    `json:"id" db:"id"`
	Name string `json:"name" db:"name"`
}
type Genre struct {
	GenreSimple
	Description string `json:"description" db:"description"`
}

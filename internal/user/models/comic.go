package models

type SimpleComic struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	OtherName string `json:"other_name"`
	Image     string `json:"image"`
}

package models

type ComicStatus int
type ComicSortBy int

const (
	StatusOnGoing  ComicStatus = 0
	StatusFinished ComicStatus = 1
)

const (
	SortByNewComic    ComicSortBy = 0
	SortByNewChapter  ComicSortBy = 1
	SortByMostView    ComicSortBy = 2
	SortByMostChapter ComicSortBy = 3
)

const (
	SortDescending = 0
	SortAscending  = -1
)

type ComicFilter struct {
	Genres     []int       `json:"genres" form:"genres" db:"genre_ids"`
	Author     int         `json:"author" form:"author" db:"author_id"`
	Status     ComicStatus `json:"status" form:"status" db:"status"`
	MinChapter int         `json:"min_chapter" form:"min_chapter" db:"min_chapter"`
	SortBy     ComicSortBy `json:"sort_by" form:"sort_by" db:"sort_by"`
	Sort       int         `json:"sort" form:"sort" db:"sort"`
}

type ComicSearch struct {
	Keyword string `json:"-" form:"keyword" binding:"required"`
}

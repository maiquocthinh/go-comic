package common

import "math"

type Paging struct {
	Page     int `json:"page" form:"page"`
	PageSize int `json:"page_size" form:"page_size"`
	//
	Total     int  `json:"total"`
	TotalPage int  `json:"total_page"`
	HasMore   bool `json:"has_more"`
}

func (p *Paging) Fulfill() {
	if p.Page <= 0 {
		p.Page = 1
	}

	if p.PageSize <= 0 {
		p.PageSize = 30
	}
}
func (p *Paging) Sync() {
	p.TotalPage = int(math.Ceil(float64(p.Total) / float64(p.PageSize)))
	p.HasMore = p.Page < p.TotalPage
}

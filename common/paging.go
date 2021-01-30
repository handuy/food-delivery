package common

type Paging struct {
	Page int     `json:"page"`
	Limit int    `json:"limit"`
	Total int    `json:"total"`
}

func(p *Paging) FillPage() {
	if p.Page <= 0 {
		p.Page = 0
	}

	if p.Limit <= 0 {
		p.Limit = 10
	}
}
package dto

type ListQuery struct {
	PageNum  int
	PageSize int
}

type List[T any] struct {
	Items    []T   `json:"items"`
	PageNum  int   `json:"pageNum"`
	PageSize int   `json:"pageSize"`
	Total    int64 `json:"total"`
}

type DeleteIds struct {
	Ids []*int `json:"ids"`
}

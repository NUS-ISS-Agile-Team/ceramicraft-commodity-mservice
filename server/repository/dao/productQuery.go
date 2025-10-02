package dao

type ListProductQuery struct {
	Keyword string
	Category string
	Offset int
	Limit int
	IsCustomer bool
	OrderBy int // 0-updateTime desc, 1-updateTime inc
}
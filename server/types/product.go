package types

type ProductInfo struct {
	Name             string `json:"name"`
	Category         string `json:"category"`
	Price            int64  `json:"price"`
	Desc             string `json:"desc"`
	Stock            int64  `json:"stock"`
	PicInfo          string `json:"pic_info"`
	Dimensions       string `json:"dimensions"`
	Material         string `json:"material"`
	Weight           string `json:"weight"`
	Capacity         string `json:"capacity"`
	CareInstructions string `json:"care_instructions"`
	Status           int32  `json:"status"` // 0: 未上架, 1: 已上架
}

type ProductSimplifiedInfo struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Category string `json:"category"`
	Price    int64  `json:"price"`
	Desc     string `json:"desc"`
	Stock    int64  `json:"stock"`
	PicInfo  string `json:"pic_info"`
	Status   int32  `json:"status"` // 0: 未上架, 1: 已上架
}

type UpdateProductStatusRequest struct {
	Status int `json:"status"` // 0-新的状态是下架，1-新的状态是上架
}

type UpdateProductStockRequest struct {
	Stock int `json:"stock"`
}

type GetProductListQuery struct {
	Keyword    string `json:"keyword"`
	Category   string `json:"category"`
	Offset     int    `json:"offset"`
	Limit      int    `json:"limit"`
	IsCustomer bool   `json:"is_customer"`
	OrderBy    int    `json:"order_by"` // 0-updateTime desc, 1-updateTime inc
}

type GetProductListRequest struct {
	Keyword  string `json:"keyword"`
	Category string `json:"category"`
	Offset   int    `json:"offset"`
	OrderBy  int    `json:"order_by"` // 0-updateTime desc, 1-updateTime inc
}

type UpdateProductInfoRequest struct {
	ID               int    `json:"id"`
	Name             string `json:"name"`
	Category         string `json:"category"`
	Price            int64  `json:"price"`
	Desc             string `json:"desc"`
	PicInfo          string `json:"pic_info"`
	Dimensions       string `json:"dimensions"`
	Material         string `json:"material"`
	Weight           string `json:"weight"`
	Capacity         string `json:"capacity"`
	CareInstructions string `json:"care_instructions"`
}

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

type UpdateProductStatusRequest struct {
	ID     int `json:"id"`
}

type UpdateProductStockRequest struct {
	ID    int   `json:"id"`
	Stock int `json:"stock"`
}
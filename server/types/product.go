package types

type ProductInfo struct {
	Name      string `json:"name"`
	Price     int64  `json:"price"`
	Desc      string `json:"desc"`
	Stock     int64  `json:"stock"`
	PicInfo   string `json:"pic_info"`
	Status    int32  `json:"status"` // 0: 未上架, 1: 已上架
}
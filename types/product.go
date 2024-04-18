package types

type ProductSearchReq struct {
	Id            uint   `form:"id" json:"id"`
	Name          string `form:"name" json:"name"`
	CategoryId    int    `form:"category_id" json:"category_id"`
	Title         string `form:"title" json:"title" `
	Info          string `form:"info" json:"info" `
	Price         string `form:"price" json:"price"`
	DiscountPrice string `form:"discount_price" json:"discount_price"`
	OnSale        bool   `form:"on_sale" json:"on_sale"`
	BasePage
}

type ProductCreateReq struct {
	Id            uint   `form:"id" json:"id"`
	Name          string `form:"name" json:"name"`
	CategoryId    uint   `form:"category_id" json:"category_id"` //分类
	Title         string `form:"title" json:"title" `
	Info          string `form:"info" json:"info" `
	ImgPath       string `form:"img_path" json:"img_path"` //展示图
	Price         string `form:"price" json:"price"`
	DiscountPrice string `form:"discount_price" json:"discount_price"` //折扣价
	OnSale        bool   `form:"on_sale" json:"on_sale"`               //是否在售
	Num           int    `form:"num" json:"num"`
	BasePage
}

type ProductListReq struct {
	CategoryId uint `form:"category_id" json:"category_id"`
	BasePage
}

type ProductDeleteReq struct {
	Id uint `form:"id" json:"id"`
	BasePage
}

type ProductShowReq struct {
	Id uint `form:"id" json:"id"`
}

type ProductUpdateReq struct {
	Id            uint   `form:"id" json:"id"`
	Name          string `form:"name" json:"name"`
	CategoryID    uint   `form:"category_id" json:"category_id"`
	Title         string `form:"title" json:"title" `
	Info          string `form:"info" json:"info" `
	ImgPath       string `form:"img_path" json:"img_path"`
	Price         string `form:"price" json:"price"`
	DiscountPrice string `form:"discount_price" json:"discount_price"`
	OnSale        bool   `form:"on_sale" json:"on_sale"`
	Num           int    `form:"num" json:"num"`
	BasePage
}

type ListProductImgReq struct {
	Id uint `json:"id" form:"id"`
}

type ProductResp struct {
	Id            uint   `json:"id"`
	Name          string `json:"name"`
	CategoryId    uint   `json:"category_id"`
	Title         string `json:"title"`
	Info          string `json:"info"`
	ImgPath       string `json:"img_path"`
	Price         string `json:"price"`
	DiscountPrice string `json:"discount_price"`
	View          uint64 `json:"view"`
	CreatedAt     int64  `json:"created_at"`
	Num           int    `json:"num"`
	OnSale        bool   `json:"on_sale"`
	BossId        uint   `json:"boss_id"`
	BossName      string `json:"boss_name"`
	BossAvatar    string `json:"boss_avatar"`
	BasePage
}

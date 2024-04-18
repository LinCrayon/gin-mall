package types

type OrderServiceReq struct {
	OrderId   uint `form:"order_id" json:"order_id"`
	ProductId uint `form:"product_id" json:"product_id"`
	Num       uint `form:"num" json:"num"`
	AddressId uint `form:"address_id" json:"address_id"`
	Money     int  `form:"money" json:"money"`
	BossId    uint `form:"boss_id" json:"boss_id"`
	UserId    uint `form:"user_id" json:"user_id"`
	OrderNum  uint `form:"order_num" json:"order_num"`
	Type      int  `form:"type" json:"type"`
	*BasePage
}

type OrderCreateReq struct {
	OrderId   uint `form:"order_id" json:"order_id"`
	ProductId uint `form:"product_id" json:"product_id"`
	Num       uint `form:"num" json:"num"`
	AddressId uint `form:"address_id" json:"address_id"`
	Money     int  `form:"money" json:"money"`
	BossId    uint `form:"boss_id" json:"boss_id"`
	UserId    uint `form:"user_id" json:"user_id"`
	OrderNum  uint `form:"order_num" json:"order_num"`
	Type      int  `form:"type" json:"type"`
}

type OrderListReq struct {
	Type int `form:"type" json:"type"`
	BasePage
}

type OrderShowReq struct {
	OrderId uint `json:"order_id" form:"order_id"`
}

type OrderDeleteReq struct {
	OrderId uint `json:"order_id" form:"order_id"`
}

type OrderListResp struct {
	Id            uint   `json:"id"`
	OrderNum      uint64 `json:"order_num"`
	CreatedAt     int64  `json:"created_at"`
	UpdatedAt     int64  `json:"updated_at"`
	UserId        uint   `json:"user_id"`
	ProductId     uint   `json:"product_id"`
	BossId        uint   `json:"boss_id"`
	Num           uint   `json:"num"`
	AddressName   string `json:"address_name"`
	AddressPhone  string `json:"address_phone"`
	Address       string `json:"address"`
	Type          uint   `json:"type"`
	Name          string `json:"name"`
	ImgPath       string `json:"img_path"`
	DiscountPrice string `json:"discount_price"`
}

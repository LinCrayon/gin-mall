package types

type ShowMoneyReq struct {
	Key string `json:"key" form:"key"`
}

type ShowMoneyResp struct {
	UserId    uint   `json:"user_id" form:"user_id"`
	UserName  string `json:"user_name" form:"user_name"`
	UserMoney string `json:"user_money" form:"user_money"`
}

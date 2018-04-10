package controllers

import "github.com/astaxie/beego/orm"

const (
	ApiIndex     = "/"
	ApiLogin     = "/api/login"
	ApiRegister = "/api/register"
	ApiImage = "/api/image"
	ApiWechatAuth ="/api/wechat"
	ApiProductList="/api/product"
	APiOrderList="/api/order"
)

const (
	ErrInvalid = "无效的用户ID"
	ErrDelete = "删除失败"
)
const OkDelete  = "用户注销成功"

type OrderList struct {
	Id int64 `json:"id"`
	Date int64 `json:"date"`
	TotalCount float64`json:"total_count"`
	ProductId int64 `json:"product_id"`
	//Products []*Product `json:"products" orm:"reverse(many)"`

}

type Product struct {
	Id int64 `json:"id"`
	Name string `json:"name"`
	Descript string `json:"descriptor"`
	Price float64 `json:"price"`
	SalePrice float64 `json:"sale_price"`
	Category int64 `json:"category"`
	ImgUrl string `json:"img_url"`
	order OrderList `orm:"ref(fk)"`
}


func init()  {
	orm.RegisterModel(new (Product),new(OrderList))
}
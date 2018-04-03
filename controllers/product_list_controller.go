package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"strconv"
	"fmt"
)

type ProductListController struct {
	beego.Controller
}

type Product struct {
	Id int64 `json:"id"`
	Name string `json:"name"`
	Descript string `json:"descriptor"`
	Price float64 `json:"price"`
	SalePrice float64 `json:"sale_price"`
	Category int64 `json:"category"`
	ImgUrl string `json:"img_url"`
}


func (this * ProductListController)Get()  {
	
	this.Ctx.WriteString("welcome to seek.com")
}


func (this * ProductListController)Post()  {
	product_name:=this.Ctx.Request.Form.Get("name")
	product_price := this.Ctx.Request.Form.Get("price")
	product_sale_price:= this.Ctx.Request.Form.Get("sale_price")
	product_img_url:= this.Ctx.Request.Form.Get("img_url")
	product_desc:= this.Ctx.Request.Form.Get("descript")
	product_category:= this.Ctx.Request.Form.Get("category")

	price,_:= strconv.ParseFloat(product_price,64)
	sale_price,_:= strconv.ParseFloat(product_sale_price,64)
	category,_:= strconv.ParseInt(product_category,10,64)
	product:=new (Product)
	product.Name = product_name
	product.Descript=product_desc
	product.Category = category
	product.ImgUrl = product_img_url
	product.Price = price
	product.SalePrice = sale_price
	o:= orm.NewOrm()

	fmt.Println(o.Insert(product))
	this.Ctx.WriteString("OK")

}

func init()  {
	orm.RegisterModel(new (Product))
}

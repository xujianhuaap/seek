package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"strconv"
	"github.com/xujianhuaap/seek/net"
	"encoding/json"
)

type ProductListController struct {
	beego.Controller
}




func (this * ProductListController)Get()  {
	var products []*Product
	o:= orm.NewOrm()
	_,err:=o.QueryTable(new(Product)).All(&products)
	var resp net.Response
	if err != nil {
		resp=net.Response{1,err.Error(),net.DataNil{}}
	}else {
		resp=net.Response{0,"",products}
	}
	json,err:= json.Marshal(resp)
	defer this.Ctx.WriteString(string(json))
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
	index,err:=o.Insert(product)
	var resp net.Response
	if err!= nil{
		resp=net.Response{1,err.Error(),"create failrue"}
	}else {
		resp=net.Response{0,"",index}
	}
	json,_:=json.Marshal(resp)
	this.Ctx.WriteString(string(json))
}



package main

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/xujianhuaap/seek/controllers"
	_"github.com/go-sql-driver/mysql"

	"github.com/astaxie/beego/orm"
)



func main() {
	fmt.Println("hello,seek")
	beego.Run("127.0.0.1:8098")
}

func init()  {
	beego.Router(controllers.ApiIndex,&controllers.IndexController{})
	beego.Router(controllers.ApiLogin,&controllers.LoginController{})
	beego.Router(controllers.ApiRegister,&controllers.RegisterController{})
	beego.Router(controllers.ApiWechatAuth,&controllers.WechatAuthController{})
	beego.Router(controllers.ApiProductList,&controllers.ProductListController{})
	beego.Router(controllers.APiOrderList,&controllers.OrderListController{})
	beego.SetStaticPath(controllers.ApiImage,"images")

	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", "root:123456@/seek?charset=utf8")
}
package main

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/xujianhuaap/seek/controllers"
	_"github.com/go-sql-driver/mysql"

)

func main() {
	fmt.Println("hello,seek")
	beego.Run("127.0.0.1:8098")
}

func init()  {
	beego.Router(controllers.ApiIndex,&controllers.IndexController{})
	beego.Router(controllers.ApiLogin,&controllers.LoginController{})
	beego.Router(controllers.ApiRegister,&controllers.RegisterController{})
	beego.SetStaticPath(controllers.ApiImage,"images")
}
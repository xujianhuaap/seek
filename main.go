package main

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/xujianhuaap/seek/controllers"
	_"github.com/go-sql-driver/mysql"

)

func main() {
	fmt.Println("hello,seek")
	beego.Run("127.0.0.1:8099")
}

func init()  {
	beego.Router(controllers.Api_Index,&controllers.IndexController{})
	beego.Router(controllers.Api_Login,&controllers.LoginController{})
}
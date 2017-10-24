package controllers

import (
	"github.com/astaxie/beego"
	"github.com/xujianhuaap/seek/net"
	"encoding/json"
	"github.com/xujianhuaap/seek/data/bean"
	"fmt"
)

type LoginController struct {
	beego.Controller
}

func (this * LoginController)Get()  {
	this.Ctx.WriteString("seek never stop")
	
}

func (this * LoginController)Post()  {
	//fmt.Printf("Request Header \t%v\n",this.Ctx.Request.Header)
	fmt.Printf("Request Form \t%v\n",this.Ctx.Request.Form)
	//buffer:= make([]byte,this.Ctx.Request.ContentLength)
	//this.Ctx.Request.Body.Read(buffer)
	//fmt.Printf("Request Body \t%v\n",string(buffer))
	//fmt.Printf("Request URI \t%v\n",this.Ctx.Request.RequestURI)
	//fmt.Printf("Request PostForm \t%v\n",this.Ctx.Request.PostForm)
	fmt.Printf("Request MultipartForm \t%v\n",this.Ctx.Request.MultipartForm)

	isExits := bean.LoadUserFromDB(this.Ctx.Request.Form.Get("mobile"),this.Ctx.Request.Form.Get("password"))

	var resp interface{}
	if isExits {
		resp= net.Response{0,"login success",net.DataNil{}}

	}else {
		resp= net.Response{1,"user not register",net.DataNil{}}
	}

	byte,err:= json.Marshal(resp)
	if err == nil {
		this.Ctx.WriteString(string(byte))
	}else {
		this.Ctx.WriteString("{}")
	}



}

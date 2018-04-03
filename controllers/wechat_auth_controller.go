package controllers

import "github.com/astaxie/beego"

type WechatAuthController struct {
	beego.Controller
}

func (this * WechatAuthController)Get()  {
	this.Ctx.WriteString("Wechat\n")
}


func (this * WechatAuthController)Post()  {
	this.Ctx.WriteString("Wechat\n")
}

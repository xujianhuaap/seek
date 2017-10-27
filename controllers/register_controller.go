package controllers

import (
	"github.com/astaxie/beego"
	"github.com/xujianhuaap/seek/data/bean"
	"github.com/xujianhuaap/seek/net"
	"encoding/json"
	"strings"
	"strconv"
)

type RegisterController struct {
	beego.Controller
}

func (this *RegisterController) Post()  {
	var mobile, password string
	mobile = this.Ctx.Request.Form.Get("mobile")
	password = this.Ctx.Request.Form.Get("password")
	result := bean.WriteUserToDB(mobile,password)
	var resp interface{}
	if result {
		resp = net.Response{0,"register success",net.DataNil{}}
	}else {
		resp = net.Response{1,"register failure",net.DataNil{}}
	}
	json,err := json.Marshal(resp)
	if err != nil {
		this.Ctx.WriteString("{}")
	}else {
		this.Ctx.WriteString(string(json))
	}
}

func (this *RegisterController) Delete()  {
	var msg string
	var id string
	var result bool
	id = strings.TrimSpace(this.Ctx.Request.Form.Get("id"))

	if len(id) <= 0 {
		msg = ErrInvalid
	}else {
		userId,err := strconv.ParseInt(id,10,64)
		if err == nil {
			result = bean.DeleteUserFromDB(userId)
			if result {
				msg = OkDelete
			}else{
				msg = ErrDelete
			}
		}else {
			msg = ErrInvalid
		}

	}

	var resp interface{}
	if result {
		resp = net.Response{0,msg,net.DataNil{}}
	}else {
		resp = net.Response{1,msg,net.DataNil{}}
	}
	bytes,err := json.Marshal(resp)
	if err == nil {
		this.Ctx.WriteString(string(bytes))
	}else {
		this.Ctx.WriteString("{}")
	}

}

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
type User struct {
	UserId string `json:"user_id"`
	Mobile string `json:"mobile"`
}

func (this * LoginController)Get()  {
	this.Ctx.WriteString("seek never stop")
	
}

func (this * LoginController)Post()  {
	//resp,_:=httplib.Post("http://127.0.0.1:8098/api/wechat").Response()
	//data,_:=ioutil.ReadAll(resp.Body)
	//fmt.Printf("resp %v \n",string(data))


	fmt.Printf("Request Header \t%v\n",this.Ctx.Request.Header)
	fmt.Printf("Request Form \t%v\n",this.Ctx.Request.Form)
	//buffer:= make([]byte,this.Ctx.Request.ContentLength)
	//this.Ctx.Request.Body.Read(buffer)
	//fmt.Printf("Request Body \t%v\n",string(buffer))
	//fmt.Printf("Request URI \t%v\n",this.Ctx.Request.RequestURI)
	//fmt.Printf("Request PostForm \t%v\n",this.Ctx.Request.PostForm)
	fmt.Printf("Request MultipartForm \t%v\n",this.Ctx.Request.MultipartForm)
	mobile:= this.Ctx.Request.Form.Get("mobile");
	password := this.Ctx.Request.Form.Get("password");
	isExits,tip := bean.LoadUserFromDB(mobile,password);

	var resp interface{}
	if isExits {
		var data User = User{UserId:tip,Mobile:mobile}
		resp= net.Response{0,"login success",&data}

	}else {
		resp= net.Response{1,tip,net.DataNil{}}
	}

	byte,err:= json.Marshal(&resp)
	fmt.Printf("result\t%v\n",string(byte))
	if err == nil {
		this.Ctx.WriteString(string(byte))
	}else {
		this.Ctx.WriteString("{}")
	}




}

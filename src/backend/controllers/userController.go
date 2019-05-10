package controllers

import (
	"Friday/models"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type UserController struct {
	beego.Controller
}

type JSONStruct struct {
	Rtn int
	Msg string
}

func (c* UserController) Login() {
	user_name := c.GetString("username")
	password := c.GetString("password")
	// check input
	if user_name == "" || password == ""{
		c.Data["json"] = &JSONStruct{-1, "登陆失败"}
		c.ServeJSON()
		return
	}
	// check password
	o := orm.NewOrm()
	user := models.User{}
	err := o.QueryTable("user").Filter("user_name", user_name).One(&user)
	if err == nil{
		// query success
		if password == user.Password{
			// password correct
			c.SetSession("user_id", user.Id)
			c.Data["json"] = &JSONStruct{0, "ok"}
			c.ServeJSON()
			return
		}
	}
	// failed
	c.Data["json"] = &JSONStruct{-1, "登陆失败"}
	c.ServeJSON()
	return
}

func (c* UserController) Logout() {
	c.DelSession("user_id")
	c.Data["json"] = &JSONStruct{0, "ok"}
	c.ServeJSON()
	return
}

func (c* UserController) Register() {
	user_name := c.GetString("username")
	password := c.GetString("password")
	if user_name == "" || password == ""{
		c.Data["json"] = &JSONStruct{-1, "用户名和密码不能为空"}
		c.ServeJSON()
		return
	}
	// check password
	o := orm.NewOrm()
	user := models.User{}
	// check if user name exist
	err := o.QueryTable("user").Filter("user_name", user_name).One(&user)
	if err == orm.ErrNoRows {
		beego.Info(user.Password)
		// create new user
		user.UserName = user_name
		user.Password = password
		_, err = o.Insert(&user)
		if err != nil{
			c.Data["json"] = &JSONStruct{-1, "注册失败"}
			beego.Info("insert into database failed")
			c.ServeJSON()
			return
		} else {
			beego.Info("insert success")
			c.Data["json"] = &JSONStruct{0, "ok"}
			c.ServeJSON()
			return
		}
	}else {
		c.Data["json"] = &JSONStruct{1, "用户名已存在，请更换用户名注册"}
		c.ServeJSON()
		return
	}
}
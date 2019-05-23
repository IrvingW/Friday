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

type IfLoginStruct struct {
	Rtn int
	Msg string
	Username string
}

func (c* UserController) IfLogin() {
	user_name := c.GetSession("user_name")
	if user_name == nil {
		// not login
		c.Data["json"] = &IfLoginStruct{0, "not log in", ""}
	}else{
		// logged in
		c.Data["json"] = &IfLoginStruct{0, "logged in", user_name.(string)}
	}
	c.ServeJSON()
	return
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
	user, err := models.UserModel.GetUserByUsername(user_name)
	if err == nil{
		// query success
		if password == user.Password{
			// password correct
			c.SetSession("user_id", user.Id)
			c.SetSession("user_name", user.UserName)
			c.Data["json"] = &JSONStruct{0, "ok"}
			c.ServeJSON()
			return
		}
	}
	beego.Error(err)
	// failed
	c.Data["json"] = &JSONStruct{-1, "登陆失败"}
	c.ServeJSON()
	return
}

func (c* UserController) Logout() {
	beego.Debug(c.GetSession("user_id"))
	c.DelSession("user_id")
	c.DelSession("user_name")
	c.Data["json"] = &JSONStruct{0, "ok"}
	c.ServeJSON()
	return
}

func (c* UserController) Register() {
	user_name := c.GetString("username")
	password := c.GetString("password")
	token := c.GetString("token")
	if user_name == "" || password == "" || token == ""{
		c.Data["json"] = &JSONStruct{-1, "用户名和密码和令牌不能为空"}
		c.ServeJSON()
		return
	}
	// check user_name
	_, err := models.UserModel.GetUserByUsername(user_name)
	if err == orm.ErrNoRows {
		// create new user
		u := models.User{UserName:user_name, Password:password}
		_, err = models.UserModel.AddUser(&u, token)
		if err != nil{
			c.Data["json"] = &JSONStruct{-1, "注册失败"}
			beego.Error(err)
			c.ServeJSON()
			return
		} else {
			c.Data["json"] = &JSONStruct{0, "ok"}
			c.ServeJSON()
			return
		}
	}else {
		c.Data["json"] = &JSONStruct{-1, "用户名已存在，请更换用户名注册"}
		c.ServeJSON()
		return
	}
}

type creatUserRtn struct {
	Rtn int
	Msg string
	Token string
}

func (c* UserController) CreateUser() {
	// check user role
	user_id := c.GetSession("user_id")
	if user_id == nil {
		beego.Error("get user id from session error")
		c.Data["json"] = &creatUserRtn{-1, "身份验证失败", ""}
		c.ServeJSON()
		return
	}
	admin, err := models.UserModel.IsAdmin(user_id.(int))
	if err != nil {
		beego.Error(err)
		c.Data["json"] = &creatUserRtn{-1, "身份验证失败", ""}
		c.ServeJSON()
		return
	}
	if !admin {
		beego.Error("not admin")
		c.Data["json"] = &creatUserRtn{-1, "无权限创建用户", ""}
		c.ServeJSON()
		return
	}
	auth := c.GetString("auth")
	token, err := models.UserModel.GenerateToken(auth)
	if err != nil {
		beego.Error(err)
		c.Data["json"] = &creatUserRtn{-1, "令牌生成失败", ""}
		c.ServeJSON()
		return
	}
	c.Data["json"] = &creatUserRtn{0, "success", token}
	c.ServeJSON()
}

type ifAdminRtn struct {
	Rtn int
	Msg string
	Admin bool
}

func (c *UserController) IfAdmin() {
	// check user role
	user_id := c.GetSession("user_id")
	if user_id == nil {
		beego.Error("get user id from session error")
		c.Data["json"] = &ifAdminRtn{-1, "身份验证失败", false}
		c.ServeJSON()
		return
	}
	admin, err := models.UserModel.IsAdmin(user_id.(int))
	if err != nil {
		beego.Error(err)
		c.Data["json"] = &ifAdminRtn{-1, "身份验证失败", false}
		c.ServeJSON()
		return
	}
	c.Data["json"] = &ifAdminRtn{0, "success", admin}
	c.ServeJSON()
}
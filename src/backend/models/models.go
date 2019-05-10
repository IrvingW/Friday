package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	)


type User struct {
	Id int `orm:"pk;auto"`
	UserName string
	Password string
}

func init(){
	orm.RegisterDataBase("default", "mysql", "irving:mysql2016@tcp(127.0.0.1:3306)/dclab?charset=utf8")
	orm.RegisterModel(new (User))
	orm.RunSyncdb("default", false, true)
}

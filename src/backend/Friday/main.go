package main

import (
	_ "Friday/models"
	_ "Friday/routers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/plugins/cors"
)

func main() {
	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Authorization", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
		AllowCredentials: true,

	}))

	// cron tasks
	//helloTask := toolbox.NewTask("hello", "10 * * * * *", cron_task.Hello)
    //toolbox.AddTask("hello", helloTask)
	//toolbox.StartTask()
	//defer toolbox.StopTask()


	beego.Run()
}


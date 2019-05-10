package routers

import (
	"Friday/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})

	beego.Router("/api/user/login", &controllers.UserController{}, "post:Login")
	beego.Router("/api/user/logout", &controllers.UserController{}, "get:Logout")
	beego.Router("/api/user/register", &controllers.UserController{}, "post:Register")

	// cluster controller
	beego.Router("/api/cluster/list", &controllers.ClusterController{}, "get:ClusterList")
	beego.Router("/api/cluster/add", &controllers.ClusterController{}, "post:AddCluster")
	beego.Router("/api/cluster/detail", &controllers.ClusterController{}, "get:ClusterDetail")
	beego.Router("/api/cluster/remove", &controllers.ClusterController{}, "get:RemoveCluster")

	// task controller
	beego.Router("/api/task/create", &controllers.TaskController{}, "post:CreateTask")
	beego.Router("/api/task/start", &controllers.TaskController{}, "get:StartTask")
	beego.Router("/api/task/stop", &controllers.TaskController{}, "get:StopTask")
	beego.Router("/api/task/show", &controllers.TaskController{}, "get:ShowTask")
	beego.Router("/api/task/remove", &controllers.TaskController{}, "get:RemoveTask")
	beego.Router("/api/task/update", &controllers.TaskController{}, "post:UpdateTask")
}

package routers

import (
	"Friday/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/api/user/login", &controllers.UserController{}, "post:Login")
	beego.Router("/api/user/logout", &controllers.UserController{}, "get:Logout")
	beego.Router("/api/user/register", &controllers.UserController{}, "post:Register")
	beego.Router("/api/user/if_login", &controllers.UserController{}, "get:IfLogin")
	beego.Router("/api/user/create", &controllers.UserController{}, "get:CreateUser")
	beego.Router("/api/user/if_admin", &controllers.UserController{}, "get:IfAdmin")

	// cluster controller
	beego.Router("/api/cluster/search", &controllers.ClusterController{}, "get:SearchCluster")
	beego.Router("/api/cluster/add", &controllers.ClusterController{}, "post:AddCluster")
	beego.Router("/api/cluster/detail", &controllers.ClusterController{}, "get:ClusterDetail")
	beego.Router("/api/cluster/remove", &controllers.ClusterController{}, "get:RemoveCluster")
	beego.Router("/api/cluster/machine/add", &controllers.ClusterController{}, "post:AddMachine")

	// task controller
	beego.Router("/api/task/create", &controllers.TaskController{}, "post:CreateTask")
	beego.Router("/api/task/start", &controllers.TaskController{}, "get:StartTask")
	beego.Router("/api/task/stop", &controllers.TaskController{}, "get:StopTask")
	beego.Router("/api/task/show", &controllers.TaskController{}, "get:ShowTask")
	beego.Router("/api/task/remove", &controllers.TaskController{}, "get:RemoveTask")
	beego.Router("/api/task/update", &controllers.TaskController{}, "post:UpdateTask")

	// image controller
	beego.Router("/api/image/package/upload", &controllers.ImageController{}, "post:UploadPackage")
	beego.Router("/api/image/dockerfile/upload", &controllers.ImageController{}, "post:UploadDockerfile")
	beego.Router("/api/image/add/package", &controllers.ImageController{}, "post:AddImageFromPackage")
	beego.Router("/api/image/add/dockerfile", &controllers.ImageController{}, "post:AddImageFromDockerfile")
	beego.Router("/api/image/repo/create", &controllers.ImageController{}, "post:CreateRepo")
	beego.Router("/api/image/repo/search", &controllers.ImageController{}, "get:SearchRepo")
	beego.Router("/api/image/repo/info", &controllers.ImageController{}, "get:ShowRepoInfo")
	beego.Router("/api/image/label/list", &controllers.ImageController{}, "get:GetAllLabels")

	// picture controller
	beego.Router("/api/picture/show", &controllers.PictureController{}, "get:GetPicture")
	beego.Router("/api/picture/save", &controllers.PictureController{}, "post:SavePicture")
}

package models

import (
	"github.com/astaxie/beego/orm"
	"sync"
)

var (
	globalOrm orm.Ormer
	once      sync.Once

	// models
	UserModel		*userModel
	ClusterModel	*clusterModel
	TaskModel		*taskModel
	KubeHelper		*kubeHelper
	KubeClient      *kubeClient
	RegistryModel 	*registryModel
	ImageModel    	*imageModel
	RedisModel 		*redisModel
)

const DBDriver = "mysql"
const DBAlias = "default"

func init() {
	// connect database
	orm.RegisterDataBase(DBAlias, DBDriver, "irving:mysql2016@tcp(127.0.0.1:3306)/dclab?charset=utf8")

	// init orm tables
	orm.RegisterModel(
		new(User),
		new(Cluster),
		new(Machine),
		new(MachineStatus),
		new(Deployment),
		new(Task),
		new(Label),
		new(Image),
		new(Repository))

	orm.RunSyncdb(DBAlias, false, true)

	// init public models
	UserModel = &userModel{}
	ClusterModel = &clusterModel{}
	TaskModel = &taskModel{}
	KubeHelper = &kubeHelper{}
	KubeClient = &kubeClient{}
	RegistryModel = &registryModel{}
    ImageModel = &imageModel{}
    RedisModel = &redisModel{}

    // insert admin user
    admin := &User{UserName: "admin", Password: "admin", Admin:true, }
    orm.NewOrm().Insert(admin)
}


// singleton init ormer ,only use for normal db operation
// if you begin transaction，please use orm.NewOrm()
func Ormer() orm.Ormer {
	once.Do(func() {
		globalOrm = orm.NewOrm()
	})
	return globalOrm
}
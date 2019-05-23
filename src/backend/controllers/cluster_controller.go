package controllers

import (
	"Friday/models"
	"github.com/astaxie/beego"
)

type ClusterController struct {
	beego.Controller
}

type searchClusterRtn struct {
	Rtn int
	Msg string
	Clusters []*models.Cluster
}

func (c* ClusterController) SearchCluster(){
	key := c.GetString("key")
	clusterList, err := models.ClusterModel.SearchCluster(key)
	if err != nil {
		beego.Error(err)
		c.Data["json"] = &searchClusterRtn{-1, "获取集群列表失败", nil}
		c.ServeJSON()
		return
	}
	c.Data["json"] = &searchClusterRtn{0, "success", clusterList}
	c.ServeJSON()
	return
}

func (c* ClusterController) ClusterDetail(){

}

func (c* ClusterController) AddCluster(){
	type addClusterRtn struct {
		Rtn int
		Msg string
	}
	masterId, err := c.GetInt("master_id")
	if err != nil {
		beego.Error(err)
		c.Data["json"] = &addClusterRtn{-1, "参数获取失败"}
		c.ServeJSON()
		return
	}
	clusterName := c.GetString("name")
	description := c.GetString("description")
	kubePath := c.GetString("path")
	port, err := c.GetInt("port")
	if err != nil || clusterName == "" || kubePath == ""{
		beego.Error(err)
		c.Data["json"] = &addClusterRtn{-1, "参数获取失败"}
		c.ServeJSON()
	}
	master, err:= models.ClusterModel.GetMachineById(masterId)
	if err != nil {
		beego.Error(err)
		c.Data["json"] = &addClusterRtn{-1, "获取master节点失败"}
		c.ServeJSON()
	}
	key, err := models.ClusterModel.ExtractKey(master)
	if err != nil {
		beego.Error(err)
		c.Data["json"] = &addClusterRtn{-1, "获取master节点失败"}
		c.ServeJSON()
		return
	}
	valid, err := models.KubeHelper.CheckMaster(*key, kubePath, port)
	if err != nil || !valid{
		beego.Error(err)
		c.Data["json"] = &addClusterRtn{-1, "验证Kubernetes服务失败"}
		c.ServeJSON()
		return
	}

	// add cluster
	userId := c.GetSession("user_id").(int)
	_, err = models.ClusterModel.AddCluster(masterId, clusterName, description, kubePath, port, userId, nil)
	if err != nil {
		beego.Error(err)
		c.Data["json"] = &addClusterRtn{-1, "添加集群失败"}
		c.ServeJSON()
	}
	c.Data["json"] = &addClusterRtn{0, "success"}
	c.ServeJSON()
}

func (c* ClusterController) RemoveCluster(){

}

func (c* ClusterController) UpdateCluster() {

}

type addMachineRtn struct {
	Rtn int
	Msg string
	MasterId int
}

func (c* ClusterController) AddMachine() {
	ip := c.GetString("ip")
	userName := c.GetString("user_name")
	password := c.GetString("password")
	port, err := c.GetInt("port")
	if err != nil {
		beego.Error(err)
		c.Data["json"] = &addMachineRtn{-1, "参数中获取端口信息失败", -1}
		c.ServeJSON()
		return
	}
	machineKey := models.MachineKey{Host: ip, Password: password, Port: port, UserName: userName}
	masterId, err := models.KubeHelper.AddMachine(machineKey, nil)
	if err != nil {
		beego.Error(err)
		c.Data["json"] = &addMachineRtn{-1, "添加机器节点失败", -1}
		c.ServeJSON()
		return
	}
	c.Data["json"] = &addMachineRtn{0, "success", masterId}
	c.ServeJSON()
	return
}

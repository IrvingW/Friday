package controllers

import (
	"Friday/models"
	"fmt"
	"github.com/astaxie/beego"
	"os/exec"
)

type TaskController struct {
	beego.Controller
}


type createTaskRtn struct {
	Rtn int
	Msg string
}

func submitTask(imageName string, className string, taskId int, taskName string, jarPath string){
	command := fmt.Sprintf("/home/irving/friday/spark/bin/spark-submit" +
		" --master k8s://https://192.168.1.116:6443" +
		" --deploy-mode cluster" +
		" --name %s" +
		" --class %s" +
		" --conf spark.executor.instances=1" +
		" --conf spark.kubernetes.container.image=%s" +
		" --conf spark.kubernetes.authenticate.driver.serviceAccountName=spark" +
		" local://%s", taskName, className, imageName, jarPath)

	beego.Info(command)
	cmd := exec.Command("/bin/bash", "-c", command)
	out, err := cmd.Output()
	if err != nil {
		beego.Error(err)
	}
	beego.Info(string(out))

	// insert result into database
	task := &models.Task{Id:taskId}
	models.Ormer().Read(task)
	task.Output = string(out)
	task.Status = "done"
	_, err = models.Ormer().Update(task)
	if err != nil {
		beego.Error(err)
	}
}


func (c* TaskController) CreateTask(){
	imageId, err := c.GetInt("image_id")
	if err != nil {
		beego.Error(err)
		c.Data["json"] = &createTaskRtn{-1, "获取镜像信息失败"}
		c.ServeJSON()
		return
	}
	className := c.GetString("class_name")
	taskName := c.GetString("name")
	description := c.GetString("description")
	jarPath := c.GetString("jar_path")

	userId := c.GetSession("user_id").(int)

	// insert into database
	task := &models.Task{Name: taskName, Description: description}
	taskId, err := models.TaskModel.CreateTask(task, userId)

	// use image submit spark task
	image := &models.Image{Id:imageId}
	models.Ormer().Read(image)
	models.Ormer().LoadRelated(image, "Repository")
	imageName := "dclab.registry.docker.com:5000/" + image.Repository.Name + ":" + image.Tag

	// update download times
	image.Download = image.Download + 1
	models.Ormer().Update(image)

	go submitTask(imageName, className, taskId, taskName, jarPath)

	c.Data["json"] = &createTaskRtn{0, "success"}
	c.ServeJSON()
	return
}

type getResultRtn struct {
	Rtn int
	Msg string
	Result string
}

func (c* TaskController) GetResult(){
	taskId, _ := c.GetInt("task_id")
	result, err := models.TaskModel.GetResult(taskId)
	if err != nil {
		beego.Error(err)
		c.Data["json"] = &getResultRtn{-1, "获取任务运行结果失败", ""}
		c.ServeJSON()
		return
	}
	c.Data["json"] = &getResultRtn{0, "success", result}
	c.ServeJSON()
	return
}

func (c* TaskController) StartTask(){

}

func (c* TaskController) StopTask(){

}

func (c* TaskController) RemoveTask(){

}

func (c* TaskController) UpdateTask(){

}

func (c* TaskController) ShowTask() {

}


type searchTaskRtn struct {
	Rtn int
	Msg string
	Tasks []*models.Task
}

func (c* TaskController) SearchTask() {
	key := c.GetString("key")
	taskList, err := models.TaskModel.SearchTask(key)
	if err != nil {
		beego.Error(err)
		c.Data["json"] = &searchTaskRtn{-1, "获取任务列表失败", nil}
		c.ServeJSON()
		return
	}
	c.Data["json"] = &searchTaskRtn{0, "success", taskList}
	c.ServeJSON()
	return
}
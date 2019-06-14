package models

import (
	"errors"
	"github.com/astaxie/beego"
	"time"
)

type taskModel struct {}

type Deployment struct {
	Id int 						`orm:"pk;auto" json:"id"`
	Name string 				`orm:"unique" json:"name"`
	FilePath string 			`json:"file_path"`
	FileContent string 			`orm:"null;type(text)" json:"file_content"`
	Task *Task 					`orm:"rel(fk)" json:"task_id"`
}

type Task struct {
	Id int 						`orm:"pk;auto" json:"id"`
	Name string		 			`orm:"unique;" json:"name"`
	Description string 			`orm:"null;type(text)" json:"description"`
	CreateTime time.Time 		`orm:"auto_now_add;type(datetime)" json:"create_time"`
	StartTime time.Time `orm:"auto_now;type(datetime)" json:"start_time"`
	Cluster *Cluster			`orm:"rel(fk)"  json:"cluster"`
	Creator *User				`orm:"rel(fk)" json:"creator"`
	Status string 				`json:"status"`
	Output string               `orm:"type(text)" json:"output"`
}

func (* taskModel) SearchTask(key string) (taskList []*Task, err error) {
	var list []*Task
	_, err = Ormer().QueryTable("task").Filter("name__icontains", key).All(&list)
	if err != nil {
		beego.Error(err)
		return nil , errors.New("query cluster by name error")
	}
	for _, t := range list {
		Ormer().LoadRelated(t, "Creator")
		Ormer().LoadRelated(t, "Cluster")
	}
	return list, nil
}

func (* taskModel) CreateTask(task *Task, userId int, ) (rtn int, err error) {
	creator := &User{Id: userId}
	task.Creator = creator
	cluster := &Cluster{Id: 1}
	task.Cluster = cluster
	task.Status = "running"
	taskId, err := Ormer().Insert(task)
	if err != nil {
		beego.Error(err)
		return -1, errors.New("insert task into database error")
	}
	return int(taskId), nil
}

func (* taskModel) GetResult(taskId int) (result string, err error) {
	task := &Task{Id:taskId}
	Ormer().Read(task)
	return task.Output, nil
}
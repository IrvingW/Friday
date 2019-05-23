package models

import "time"

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
	Cluster *Cluster			`orm:"rel(fk)"  json:"cluster_id"`
	Author *User				`orm:"rel(fk)" json:"user_id"`
	Status *string 				`json:"status"`
	Deployments []*Deployment	`orm:"reverse(many)" json:"-"`
}

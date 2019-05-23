package models

import (
	"errors"
	"github.com/astaxie/beego"
	"time"
)

type clusterModel struct {
}

type MachineStatus struct {
	Id int 						`orm:"pk;auto" json:"id"`
	Machine *Machine			`orm:"reverse(one)" json:"-"`
	DiskTotal float32			`orm:"default(0)" json:"disk_total"`
	DiskUsage float32			`orm:"default(0)" json:"disk_usage"`
	DiskUsedPercent float32 	`orm:"default(0)" json:"disk_used_percent"`
	MemoryTotal float32			`orm:"default(0)" json:"memory_total"`
	MemoryUsage float32 		`orm:"default(0)" json:"memory_usage"`
	MemoryUsedPercent float32 	`orm:"default(0)" json:"memory_used_percent"`
	CpuUsedPercent    float32 	`orm:"default(0)" json:"cpu_used_percent"`
	CpuCnt	int		 			`orm:"default(0)" json:"cpu_cnt"`
}

type Machine struct {
	Id int						`orm:"pk;auto" json:"id"`
	Ip string					`orm:"unique" json:"ip"`
	UserName string 			`json:"user_name"`
	Password string 			`json:"-"`
	Port int					`json:"port"`
	Cluster *Cluster 			`orm:"rel(fk);null" json:"-"`
	Status *MachineStatus 		`orm:"rel(one)" json:"status"`
}

type Cluster struct {
	Id int 					`orm:"pk;auto" json:"id"`
	Name string 			`orm:"unique" json:"name"`
	Description string 		`orm:"null;type(text)" json:"description"`
	Master *Machine			`orm:"rel(fk)" json:"master"`
	Machines []*Machine 	`orm:"reverse(many)" json:"machines"`
	KubePath string 		`json:"-"`
	KubePort int			`json:"-"`
	CreateTime time.Time 	`orm:"auto_now_add;type(datetime)" json:"create_time"`
	UserAccess []*User	 	`orm:"reverse(many)" json:"-"`
	Status string 			`orm:"default(normal)" json:"status"`
	TaskCnt int 			`orm:"default(0)" json:"task_cnt"`
}

func (*clusterModel) AddCluster(master_id int, cluster_name string, desc string, kubePath string, kubePort int, user_id int, slaves []*Machine) (id int, err error){
	// get master
	master := &Machine{Id: master_id}
	creator := &User{Id:user_id}
	machines := append(slaves, master)
	cluster := &Cluster{Master:master, Name:cluster_name, Description:desc, KubePath:kubePath, UserAccess:[]*User{creator}, Machines:machines}
	cid, err := Ormer().Insert(cluster)
	if err != nil {
		beego.Error(err)
		return -1, errors.New("insert cluster into database error")
	}
	return int(cid), nil
}


func (*clusterModel) GetAllCluster() (clusters []*Cluster, err error) {
	// return all cluster
	cluster := new(Cluster)
	_, err = Ormer().QueryTable(cluster).All(&clusters)
	if err != nil {
		beego.Error(err)
		return nil, errors.New("get all clusters error")
	}
	return clusters, nil
}

func (*clusterModel) GetClusterById(cluster_id int) (cluster *Cluster, err error){
	cluster = &Cluster{Id: cluster_id}
	if err = Ormer().Read(cluster); err != nil {
		beego.Error("get cluster by id error")
		return nil, err
	} else {
		return cluster, nil
	}
}

func (*clusterModel) GetMachineById(machine_id int) (machine *Machine, err error) {
	machine = &Machine{Id: machine_id}
	if err = Ormer().Read(machine); err != nil {
		beego.Error("get machine by id error")
		return nil, err
	} else {
		return machine, nil
	}
}

func (*clusterModel) SearchCluster(key string) (cluster []*Cluster, err error) {
	var list []*Cluster
	_, err = Ormer().QueryTable("cluster").Filter("name__icontains", key).All(&list)
	if err != nil {
		beego.Error(err)
		return nil , errors.New("query cluster by name error")
	}
	return list, nil
}

func (*clusterModel) ExtractKey(machine *Machine) (key *MachineKey, err error) {
	if machine == nil {
		beego.Error("params empty error")
		return nil, errors.New("machine is empty error")
	}
	k := &MachineKey{Host:machine.Ip, Port:machine.Port, UserName:machine.UserName, Password:machine.Password}
	return k, nil
}

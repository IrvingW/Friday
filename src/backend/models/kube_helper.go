package models

import (
	"errors"
	"fmt"
	"github.com/astaxie/beego"
	"golang.org/x/crypto/ssh"
	"net"
	"os/exec"
	"strconv"
)

type kubeHelper struct {
}

type MachineKey struct {
	Host string
	UserName string
	Password string
	Port int
}

func sshConnect(key MachineKey) (client *ssh.Client, err error){
	logError := func(err error, msg string) {
		if err != nil {
			beego.Error("%s error: %v", msg, err)
		}
	}
	config := &ssh.ClientConfig{}
	config.SetDefaults()
	config.User = key.UserName
	config.Auth = []ssh.AuthMethod{ssh.Password(key.Password)}
	config.HostKeyCallback = func(hostname string, remote net.Addr, key ssh.PublicKey) error {
		return nil
	}
	addr := fmt.Sprintf("%s:%d", key.Host, key.Port)
	client, err = ssh.Dial("tcp", addr, config)
	if err != nil {
		logError(err, "connect with machine")
		return nil, err
	}
	return client, nil
}

func (*kubeHelper) CheckConnection(key MachineKey) (valid bool, err error){
	// check normal machine (slave) ssh connection
	client, err := sshConnect(key)
	if err != nil {
		beego.Error(err)
		return false, errors.New("create ssh connection error")
	}
	// check create new session
	session, err := client.NewSession()
	if err != nil {
		beego.Error(err)
		return false, errors.New("create new session error")
	}
	// close session
	session.Close()
	return true, nil
}


func (*kubeHelper) CheckMaster(key MachineKey, kubePath string, kubePort int) (valid bool, err error){
	// check connection
	_, err = KubeHelper.CheckConnection(key)
	if err != nil {
		beego.Error(err)
		return false, errors.New("create connection error")
	}

	// check kube config file
	client, err := sshConnect(key)
	if err != nil {
		beego.Error(err)
		return false, errors.New("create ssh connection error")
	}
	// check create new session
	session, err := client.NewSession()
	if err != nil {
		beego.Error(err)
		return false, errors.New("create new session error")
	}
	defer session.Close()
	buf, err := session.CombinedOutput("if [ -f " + kubePath + " ];then echo 0; else echo 1;fi")
	if err != nil {
		beego.Error(err)
		return false, errors.New("execute check kubeconfig file command error")
	}
	if string(buf) == "0\n" {
		beego.Info("kubeconfig file exists.")
	}else if string(buf) == "1" {
		return false, errors.New("kubeconfig file not exist")
	}else {
		beego.Debug(buf)
		return false, errors.New("unexpected output")
	}

	// check kube master function
	buf, err = session.CombinedOutput("kubectl get nodes | wc -l")
	if err != nil {
		beego.Error(err)
		return false, errors.New("run kubectl command error")
	}
	lines, err := strconv.Atoi(string(buf))
	if err != nil {
		beego.Error(err)
		return false, errors.New("convert string into int error")
	}
	if lines <= 1 {
		return false, errors.New("can not get nodes error")
	}else {
		return true, nil
	}
}

func (*kubeHelper) AddMachine(key MachineKey, cluster *Cluster) (id int, err error) {
	// check machine key
	_, err = KubeHelper.CheckConnection(key)
	if err != nil {
		beego.Error(err)
		return -1, errors.New("create connection error")
	}

	// add into database
	machine := &Machine{Ip: key.Host, UserName:key.UserName, Password:key.Password, Port:key.Port, Status:&MachineStatus{}}

	id64, err := Ormer().Insert(machine)
	if err != nil {
		beego.Error(err)
		return -1, errors.New("insert machine into database error")
	}
	return int(id64), nil
}



func (*kubeHelper) RunCommand(command string, key MachineKey) (result string, err error){
	// check kube config file
	client, err := sshConnect(key)
	if err != nil {
		beego.Error(err)
		return "", errors.New("create ssh connection error")
	}
	// check create new session
	session, err := client.NewSession()
	if err != nil {
		beego.Error(err)
		return "", errors.New("create new session error")
	}
	defer session.Close()
	buf, err := session.CombinedOutput(command)
	if err != nil {
		beego.Error(err)
		return "", errors.New("execute command error")
	}
	return string(buf), nil
}

func (*kubeHelper) RunLocalCommand(command string) (result string, err error) {
	cmd := exec.Command("/bin/bash", "-c", command)
	out, err := cmd.Output()
	if err != nil {
		beego.Error(err)
		return "", errors.New("execute command error")
	}
	return string(out), nil
}
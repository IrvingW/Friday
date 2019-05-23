package models

import (
	"crypto/md5"
	"errors"
	"fmt"
	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

type userModel struct {}

type User struct {
	Id int `orm:"pk;auto" json:"id"`
	UserName string `orm:"unique;" json:"user_name"`
	Password string `json:"-"`
	Admin bool `json:"admin"`
	AccessCluster []*Cluster `orm:"rel(m2m);rel_table(UserAccessCluster)" json:"access_cluster"`
}

func (*userModel) AddUser(u *User, token string) (id int64, err error) {
	// check access cluster
	if u.AccessCluster != nil {
		return -1, errors.New("access cluster must be empty")
	}
	// check token
	conn, err := RedisModel.Connect(beego.AppConfig.String("redis_address"))
	if err != nil {
		beego.Error(err)
		return -1, errors.New("connect redis server error")
	}
	defer conn.Close()
	value, err := RedisModel.GetValue(token, conn)
	if err != nil || value == "" {
		beego.Error(err)
		return -1, errors.New("verify token error")
	}
	if value == "admin" {
		u.Admin = true
		user_id, err := Ormer().Insert(u)
		if err != nil {
			beego.Error(err)
			return -1, errors.New("insert user into database error")
		}
		return user_id, nil
	} else {
		user_id, err := Ormer().Insert(u)
		if err != nil {
			beego.Error(err)
			return -1, errors.New("insert user into database error")
		}
		clusters := strings.Split(value, ",")
		var cluster_list []int
		for _, idStr := range clusters {
			id, err := strconv.Atoi(idStr)
			if err != nil {
				beego.Error(err)
				return -1, errors.New("transform id string into int error")
			}
			cluster_list = append(cluster_list, id)
		}
		_, err = UserModel.AuthUser(int(user_id), cluster_list)
		if err != nil {
			beego.Error(err)
			return -1, errors.New("auth clusters to user error")
		}
		return user_id, nil
	}
}

func (*userModel) AuthUser(user_id int, cluster_list []int) (rtn int, err error) {
	// check user_id
	user := &User{Id:user_id}
	if err = Ormer().Read(user); err != nil {
		err = errors.New("get user by id error")
		beego.Error(err)
		return -1, err
	}

	// read origin auth
	_, err = Ormer().LoadRelated(user, "AccessCluster")
	if err != nil {
		beego.Error(err)
		return -1, err
	}
	var need_auth_id[]int = make([]int, len(cluster_list))
	for _, old_auth := range user.AccessCluster {
		for _, new_auth_id := range cluster_list {
			// already auth
			if old_auth.Id == new_auth_id {
				continue
			}
			// add need auth
			_ = append(need_auth_id, new_auth_id)
		}
	}
	// check need_auth list and add auth
	for cluster_id := range need_auth_id {
		cluster := &Cluster{Id:cluster_id}
		if err = Ormer().Read(cluster); err != nil {
			err = errors.New("get cluster by id error")
			beego.Error(err)
			return -1, err
		}
		_ = append(user.AccessCluster, cluster)
	}
	_, err = Ormer().Update(user)
	if err != nil {
		beego.Error(err)
		return -1, errors.New("update user.AccessCluster error")
	}
	return 0, nil
}

func (*userModel) GetUserById(id int) (u *User, err error) {
	u = &User{Id: id}
	if err = Ormer().Read(u); err != nil {
		beego.Error("get user by id error")
		return nil, err
	} else {
		return u, nil
	}
}

func (*userModel) GetUserByUsername(name string) (u *User, err error) {
	u = &User{UserName: name}
	if err = Ormer().Read(u, "UserName"); err != nil {
		beego.Error("get user by user_name error")
		return nil, err
	} else {
		return u, nil
	}
}

func (*userModel) IsAdmin(id int) (admin bool, err error) {
	user := &User{Id: id}
	if err = Ormer().Read(user); err != nil {
		beego.Error("query user by user id error")
		return false, errors.New("userModel: query user by user id " + strconv.Itoa(id) + " failed")
	} else {
		return user.Admin, nil
	}
}

func (*userModel) IsAccessible (user_id int, cluster_id int) (accessible bool, err error){
	// check user_id and cluster id
	user := User{Id:user_id}
	cluster := Cluster{Id:cluster_id}
	err = Ormer().Read(user)
	if err != nil {
		beego.Error("get user by id error")
		return false, errors.New("get user by id " + strconv.Itoa(user_id) + " error")
	}
	err = Ormer().Read(cluster)
	if err != nil {
		beego.Error("get cluster by id error")
		return false, errors.New("get cluster by id " + strconv.Itoa(cluster_id) + " error")
	}
	// check if accessible
	// iterate access cluster list, if id equals cluster_id then return true

	return true, nil
}

func (*userModel) GetAccessibleCluster(user_id int) (clusters []*Cluster, err error) {
	// check user id
	user := User{Id:user_id}
	err = Ormer().Read(user)
	if err != nil {
		beego.Error("get user by id error")
		return nil, errors.New("get user by id " + strconv.Itoa(user_id) + " error")
	}
	// if admin, return all cluster
	if user.Admin {
		clusters, err = ClusterModel.GetAllCluster()
		if err != nil {
			beego.Error("get all cluster error")
			return nil, err
		}else{
			return clusters, nil
		}
	} else {
		// if not admin
		//_, err = Ormer().QueryTable(new(Cluster)).Filter("UserAccess__User__Id", user.Id).All(&clusters)
		_, err = Ormer().LoadRelated(&user, "AccessCluster")
		if err != nil {
			beego.Error(err)
			return nil, errors.New("get accessible clusters error")
		}
		return user.AccessCluster, nil
	}
}

func (*userModel) GenerateToken(auth string) (token string, err error){
	randNum := fmt.Sprintf("%d", rand.Intn(9999)+1000 )
	hashName := md5.Sum( []byte( time.Now().Format("2006_01_02_15_04_05_") + randNum ) )
	key := fmt.Sprintf("%x",hashName)
	conn, err := RedisModel.Connect(beego.AppConfig.String("redis_address"))
	if err != nil {
		beego.Error(err)
		return "", errors.New("make redis connection error")
	}
	defer conn.Close()
	_, err = RedisModel.SetKV(key, auth, 300, conn)
	if err != nil {
		beego.Error(err)
		return "", errors.New("insert token string into redis error")
	}
	return key, nil
}

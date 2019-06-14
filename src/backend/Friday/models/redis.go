package models

import (
	"errors"
	"github.com/astaxie/beego"
	"github.com/garyburd/redigo/redis"
	"strconv"
)

type redisModel struct {
}

func(*redisModel) Connect(address string) (conn redis.Conn, err error){
	conn, err = redis.Dial("tcp", address)
	if err != nil {
		beego.Error(err)
		return nil, errors.New("connect to redis server error")
	}
	return conn, nil
}

func(*redisModel) SetKV(key string, value string, expire int, conn redis.Conn) (rtn int, err error) {
	_, err = conn.Do("SET", key, value, "EX", strconv.Itoa(expire))
	if err != nil {
		beego.Error(err)
		return -1, errors.New("execute redis command error")
	}
	return 0, nil
}

func (*redisModel) GetValue(key string, conn redis.Conn) (value string, err error) {
	value, err = redis.String(conn.Do("GET", key))
	if err != nil {
		beego.Error(err)
		return "", errors.New("get value by key error")
	}
	return value, nil
}
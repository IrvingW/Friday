package controllers

import (
	"crypto/md5"
	"fmt"
	"github.com/astaxie/beego"
	"io/ioutil"
	"math/rand"
	"path"
	"time"
)

type PictureController struct {
	beego.Controller
}

type savePicRtn struct {
	Rtn int
	Msg string
	PictureName string
}

func (c *PictureController) GetPicture() {
	fileName := c.GetString("picture_name")
	if fileName == "" {
		fileName = "default_image.png"
	}
	ext := path.Ext(fileName)
	readPath := beego.AppConfig.String("picture_path")
	img:= path.Join(readPath, fileName)

	c.Ctx.Output.Header("Content-Type", "image/" + ext)
	c.Ctx.Output.Header("Content-Disposition",fmt.Sprintf("inline; filename=\"%s\"",fileName))
	file, err := ioutil.ReadFile(img)
	if err != nil {
		beego.Error("file do not exist: " + readPath)
		return
	}
	c.Ctx.WriteString(string(file))
}

func (c *PictureController) SavePicture() {
	f, h, err := c.GetFile("file")//获取上传的文件
	if err != nil {
		beego.Error(err)
		c.Data["json"] = &savePicRtn{Rtn: -1, Msg: "获取上传图片失败", PictureName: ""}
		c.ServeJSON()
		return
	}
	ext := path.Ext(h.Filename)
	if ext != ".png" && ext != ".jpg" && ext != ".jpeg"{
		beego.Error("file format not valid: " + ext)
		c.Data["json"] = &savePicRtn{Rtn: -1, Msg: "后缀名不符合要求", PictureName: ""}
		c.ServeJSON()
		return
	}
	//构造文件名称
	rand.Seed(time.Now().UnixNano())
	randNum := fmt.Sprintf("%d", rand.Intn(9999)+1000 )
	hashName := md5.Sum( []byte( time.Now().Format("2006_01_02_15_04_05_") + randNum ) )
	fileName := fmt.Sprintf("%x",hashName) + ext
	filePath := path.Join(beego.AppConfig.String("picture_path"), fileName)
	defer f.Close()//关闭上传的文件，不然的话会出现临时文件不能清除的情况
	err = c.SaveToFile("file", filePath)
	if err != nil {
		beego.Error(err)
		c.Data["json"] = &savePicRtn{Rtn: -1, Msg: "文件保存失败", PictureName: ""}
		c.ServeJSON()
		return
	}
	c.Data["json"] = &savePicRtn{Rtn: 0, Msg:"success", PictureName: fileName}
	c.ServeJSON()
}
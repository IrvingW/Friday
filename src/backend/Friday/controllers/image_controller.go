package controllers

import (
	"Friday/models"
	"crypto/md5"
	"fmt"
	"github.com/astaxie/beego"
	"math/rand"
	"os"
	"path"
	"strings"
	"time"
)

type ImageController struct {
	beego.Controller
}

type uploadFileRtn struct {
	Rtn int
	Msg string
	FilePath string
}

type createImageRtn struct {
	Rtn int
	Msg string
}


func (c *ImageController) UploadPackage() {
	f, h, err := c.GetFile("file")//获取上传的文件
	if err != nil {
		beego.Error(err)
		c.Data["json"] = &uploadFileRtn{Rtn: -1, Msg: "获取上传文件失败", FilePath: ""}
		c.ServeJSON()
		return
	}
	ext := path.Ext(h.Filename)
	if ext != ".tar" {
		beego.Error("file format not valid: " + ext)
		c.Data["json"] = &uploadFileRtn{Rtn: -1, Msg: "后缀名不符合要求", FilePath: ""}
		c.ServeJSON()
		return
	}

	//创建目录
	uploadDir := path.Join(beego.AppConfig.String("upload_path"), time.Now().Format("2006-01-02"))
	err = os.MkdirAll( uploadDir , 0777)
	if err != nil {
		beego.Error(err)
		c.Data["json"] = &uploadFileRtn{Rtn: -1, Msg: "创建目录失败", FilePath: ""}
		c.ServeJSON()
		return
	}
	//构造文件名称
	rand.Seed(time.Now().UnixNano())
	randNum := fmt.Sprintf("%d", rand.Intn(9999)+1000 )
	hashName := md5.Sum( []byte( time.Now().Format("2006_01_02_15_04_05_") + randNum ) )
	fileName := fmt.Sprintf("%x",hashName) + ext

	filePath := path.Join(uploadDir, fileName)
	defer f.Close()//关闭上传的文件，不然的话会出现临时文件不能清除的情况
	err = c.SaveToFile("file", filePath)
	if err != nil {
		beego.Error(err)
		c.Data["json"] = &uploadFileRtn{Rtn: -1, Msg: "文件保存失败", FilePath: ""}
		c.ServeJSON()
		return
	}
	c.Data["json"] = &uploadFileRtn{Rtn: 0, Msg: "success", FilePath: filePath}
	c.ServeJSON()
	return
}

func (c *ImageController) UploadDockerfile() {
	f, h, err := c.GetFile("file")//获取上传的文件
	if err != nil {
		beego.Error(err)
		c.Data["json"] = &uploadFileRtn{Rtn: -1, Msg: "获取上传文件失败", FilePath: ""}
		c.ServeJSON()
		return
	}
	ext := path.Ext(h.Filename)
	if ext != ".zip"{
		beego.Error("file format not valid: " + ext)
		c.Data["json"] = &uploadFileRtn{Rtn: -1, Msg: "后缀名不符合要求", FilePath: ""}
		c.ServeJSON()
		return
	}

	//创建目录
	uploadDir := path.Join(beego.AppConfig.String("build_image_path"), time.Now().Format("2006-01-02"))
	err = os.MkdirAll( uploadDir , 0777)
	if err != nil {
		beego.Error(err)
		c.Data["json"] = &uploadFileRtn{Rtn: -1, Msg: "创建目录失败", FilePath: ""}
		c.ServeJSON()
		return
	}
	//构造文件名称
	rand.Seed(time.Now().UnixNano())
	randNum := fmt.Sprintf("%d", rand.Intn(9999)+1000 )
	hashName := md5.Sum( []byte( time.Now().Format("2006_01_02_15_04_05_") + randNum ) )
	fileName := fmt.Sprintf("%x",hashName) + ext

	filePath := path.Join(uploadDir, fileName)
	defer f.Close()//关闭上传的文件，不然的话会出现临时文件不能清除的情况
	err = c.SaveToFile("file", filePath)
	if err != nil {
		beego.Error(err)
		c.Data["json"] = &uploadFileRtn{Rtn: -1, Msg: "文件保存失败", FilePath: ""}
		c.ServeJSON()
		return
	}
	c.Data["json"] = &uploadFileRtn{Rtn: 0, Msg: "success", FilePath: filePath}
	c.ServeJSON()
}

func (c *ImageController) AddImageFromDockerfile() {
	repoId, err := c.GetInt("repo_id")
	if err != nil{
		beego.Error(err)
		c.Data["json"] = &createImageRtn{Rtn:-1, Msg:"无法解析镜像名"}
		c.ServeJSON()
		return
	}
	tagName := c.GetString("tag_name")
	if tagName == "" {
		beego.Error("image name is empty")
		c.Data["json"] = &createImageRtn{Rtn:-1, Msg:"无法解析镜像名"}
		c.ServeJSON()
		return
	}
	path := c.GetString("path")
	if path == "" {
		beego.Error("file name is empty")
		c.Data["json"] = &createImageRtn{Rtn: -1, Msg: "无法获取上传文件名称"}
		c.ServeJSON()
		return
	}
	description := c.GetString("description")

	repo, _ := models.ImageModel.GetRepoById(repoId)
	if repo == nil {
		beego.Error("get repo by name failed")
		c.Data["json"] = &createImageRtn{Rtn: -1, Msg: "无法获取镜像仓库"}
		c.ServeJSON()
		return
	}
	if c.GetSession("user_id") != repo.Owner.Id {
		beego.Error("user_id not equal to repo owner's id")
		c.Data["json"] = &createImageRtn{Rtn: -1, Msg: "无权对该镜像仓库进行操作"}
		c.ServeJSON()
		return
	}

	// import and upload image into registry

	image := &models.Image{Tag: tagName, Description: description, Repository:repo}
	rtn, err := models.ImageModel.AddImage(image, repo.Id, repo.Owner.Id)
	if err != nil || rtn != 0 {
		beego.Error(err)
		c.Data["json"] = &createImageRtn{Rtn: -1, Msg: "添加镜像失败"}
		c.ServeJSON()
		return
	}
	c.Data["json"] = &createImageRtn{Rtn: 0, Msg: "success"}
	c.ServeJSON()
	return
}

func (c *ImageController) AddImageFromPackage() {
	repoId, err := c.GetInt("repo_id")
	if err != nil{
		beego.Error(err)
		c.Data["json"] = &createImageRtn{Rtn:-1, Msg:"无法解析镜像名"}
		c.ServeJSON()
	}
	tagName := c.GetString("tag_name")
	if tagName == "" {
		beego.Error("image name is empty")
		c.Data["json"] = &createImageRtn{Rtn:-1, Msg:"无法解析镜像名"}
		c.ServeJSON()
		return
	}
	path := c.GetString("path")
	if path == "" {
		beego.Error("path name is empty")
		c.Data["json"] = &createImageRtn{Rtn: -1, Msg: "无法获取上传文件路径"}
		c.ServeJSON()
		return
	}
	description := c.GetString("description")

    repo, _ := models.ImageModel.GetRepoById(repoId)
    if repo == nil {
    	beego.Error("get repo by name failed")
		c.Data["json"] = &createImageRtn{Rtn: -1, Msg: "无法获取镜像仓库"}
		c.ServeJSON()
		return
	}
    if c.GetSession("user_id") != repo.Owner.Id {
    	beego.Error("user_id not equal to repo owner's id")
		c.Data["json"] = &createImageRtn{Rtn: -1, Msg: "无权对该镜像仓库进行操作"}
		c.ServeJSON()
		return
	}

	image := &models.Image{Tag: tagName, Description: description, Repository:repo}
	rtn, err := models.ImageModel.AddImage(image, repo.Id, repo.Owner.Id)
	if err != nil || rtn != 0 {
		beego.Error(err)
		c.Data["json"] = &createImageRtn{Rtn: -1, Msg: "添加镜像失败"}
		c.ServeJSON()
		return
	}
	c.Data["json"] = &createImageRtn{Rtn: 0, Msg: "success"}
	c.ServeJSON()
}

func (c *ImageController) CreateRepo() {
	repoName := c.GetString("repo_name")
	if repoName == "" {
		beego.Error("repository name is empty")
		c.Data["json"] = &createImageRtn{Rtn: -1, Msg: "镜像仓库名不能为空"}
		c.ServeJSON()
		return
	}
	picture := c.GetString("picture")
	description := c.GetString("description")
	introduce := c.GetString("introduce")
	labels := strings.Split(c.GetString("labels"), ",")
	user_id := c.GetSession("user_id").(int)
	repo := &models.Repository{Name: repoName, Picture: picture, Description: description, Introduce: introduce}
	_, err := models.ImageModel.CreateRepository(repo, user_id, labels)
	if err != nil {
		beego.Error(err)
		c.Data["json"] = &createImageRtn{Rtn: -1, Msg: "创建镜像仓库失败"}
		c.ServeJSON()
		return
	}

	c.Data["json"] = &createImageRtn{Rtn: 0, Msg: "success"}
	c.ServeJSON()
}


type searchRepoRtn struct {
	Rtn int
	Msg string
	RepoList []*models.Repository
}

func (c *ImageController) SearchRepo() {
	key := c.GetString("key")
	repoList, err := models.ImageModel.SearchRepo(key)
	if err != nil {
		beego.Error(err)
		c.Data["json"] = &searchRepoRtn{-1, "获取镜像仓库列表失败", nil}
		c.ServeJSON()
		return
	}
	c.Data["json"] = &searchRepoRtn{0, "success", repoList}
	c.ServeJSON()
}

type getLabelsRtn struct {
	Rtn int
	Msg string
	Labels []*models.Label
}

func (c *ImageController) GetAllLabels() {
	labelList, err := models.ImageModel.GetAllLabels()
	if err != nil {
		beego.Error(err)
		c.Data["json"] = &getLabelsRtn{-1, "获取标签列表失败", nil}
		c.ServeJSON()
		return
	}
	c.Data["json"] = &getLabelsRtn{0, "success", labelList}
	c.ServeJSON()
}

type repoInfoRtn struct {
	Rtn int
	Msg string
	Repo *models.Repository
}

func (c *ImageController) ShowRepoInfo() {
	repoId, err := c.GetInt("repo_id")
	if err != nil {
		beego.Error(err)
		c.Data["json"] = &repoInfoRtn{-1, "获取仓库ID失败", nil}
		c.ServeJSON()
		return
	}
	repo, err := models.ImageModel.GetRepoById(repoId)
	if err != nil {
		beego.Error(err)
		c.Data["json"] = &repoInfoRtn{-1, "获取仓库信息失败", nil}
		c.ServeJSON()
		return
	}
	c.Data["json"] = &repoInfoRtn{0, "success", repo}
	c.ServeJSON()
	return
}
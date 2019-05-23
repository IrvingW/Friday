package models

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/astaxie/beego"
	"net/http"
	"os"
	"strings"
)

type registryModel struct {
}

const basicUrl string = "https://dclab.registry.docker.com:5000/v2/"
const listRepoUrl string = "_catalog"
const listTagUrl string = "/tags/list"
const manifestUrl string = "manifests/"

const logTitle string = "[models/registry]"

type RepoListResult struct {
	Repositories []string
}

type TagListResult struct {
	Name string
	Tags []string
}

func (*registryModel) GetAllRepository() (repoList []string, err error) {
	response, err := http.Get(basicUrl + listRepoUrl)
	if err != nil {
		beego.Error(logTitle + "http get repository list error")
		beego.Error(err)
		return nil, errors.New(logTitle + "http request submit error")
	}
	defer response.Body.Close()
	// check status code
	if response.StatusCode != http.StatusOK {
		beego.Error(logTitle + "http response status code not ok")
		return nil, errors.New(logTitle + "http response return error")
	}
	var result RepoListResult
	if err := json.NewDecoder(response.Body).Decode(&result); err != nil {
		beego.Error(err)
		return nil, errors.New(logTitle + "decode response body error")
	}

	return result.Repositories, nil
}

func (*registryModel) GetAllTags(repoName string) (tagList []string, err error) {
	if repoName == "" {
		beego.Error(logTitle + "repository name can not be empty")
		return nil, errors.New(logTitle + "repository name can not be empty")
	}
	response, err := http.Get(basicUrl + repoName + listTagUrl)
	if err != nil {
		beego.Error(err)
		return nil, errors.New(logTitle + "http request submit error")
	}
	defer response.Body.Close()
	// check status code
	if response.StatusCode != http.StatusOK {
		beego.Error(logTitle + "http response status code not ok")
		return nil, errors.New(logTitle + "http response status code not ok")
	}
	var result TagListResult
	if err := json.NewDecoder(response.Body).Decode(&result); err != nil {
		beego.Error(err)
		return nil, errors.New(logTitle + "decode response body error")
	}
	return result.Tags, nil
}

const registryName = "dclab.registry.docker.com:5000/"



func pushImage(repoName string, tagName string) (rtn int, err error) {
	imageName := fmt.Sprint("%s:%s", repoName, tagName)
	output, err := KubeHelper.RunLocalCommand("docker push " + registryName + imageName)
	if err != nil {
		beego.Error(err)
		return -1, errors.New("execute shell command error")
	}
	signal := tagName + ":" + " digest:"
	success := strings.Contains(output, signal)
	if success {
		return 0, nil
	} else {
		beego.Error(output)
		return -1, errors.New("docker push image error")
	}
}

func pathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

const imagePath = "/home/irving/friday/images/"
const dockerFilePath = "/home/irving/friday/dockerfiles/"

func (*registryModel) CompileImage(repoName string, tagName string) (rtn int, err error) {
	path := dockerFilePath + repoName + "/" + tagName
	exist, err := pathExists(path)
	if !exist {
		beego.Error(err)
		return -1, errors.New("image path not exist error")
	}
	imageName := fmt.Sprint("%s:%s", repoName, tagName)
	output, err := KubeHelper.RunLocalCommand("docker build -t " + registryName + imageName)
	signal := "Build an image from a Dockerfile"
	success := strings.Contains(output, signal)
	if !success {
		beego.Error(output)
		return -1, errors.New("docker build image error")
	}
	_, err = pushImage(repoName, tagName)
	if err != nil {
		beego.Error(err)
		return -1, errors.New("push image error")
	}
	return 0, nil
}

func (*registryModel) ImportImage(fileName string, repoName string, tagName string) (rtn int, err error){
	// check file exist
	exist, err := pathExists(imagePath + fileName)
	if !exist {
		beego.Error("image file not exist")
		return -1, errors.New("image file not exist error")
	}
	imageName := fmt.Sprintf("%s:%s", repoName, tagName)
	_, err = KubeHelper.RunLocalCommand("docker import - " + registryName + imageName + " < " + imagePath + fileName)
	if err != nil {
		beego.Error(err)
		return -1, errors.New("import image error")
	}
	_, err = pushImage(repoName, tagName)
	if err != nil {
		beego.Error(err)
		return -1, errors.New("push image error")
	}
	return 0, nil
}
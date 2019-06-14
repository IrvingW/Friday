package models

import (
	"errors"
	"github.com/astaxie/beego"
	"time"
)

type imageModel struct {
}

type Label struct {
	Id int `orm:"pk;auto" json:"id"`
	Name string `orm:"unique" json:"name"`
}

type Image struct {
	Id int `orm:"pk;auto" json:"id"`
	Repository *Repository `orm:"rel(fk)" json:"-"`
	Tag string `json:"tag"`
	CreateTime time.Time `orm:"auto_now_add;type(datetime)" json:"create_time"`
	Description string `orm:"null;type(text)" json:"description"`
	Download int `orm:"default(0)" json:"download"`
}

type Repository struct {
	Id int	`orm:"pk;auto" json:"id"`
	Name string `orm:"unique" json:"name"`
	Owner *User `orm:"rel(fk)" json:"owner"`
	Introduce string `orm:"type(text)" json:"introduce"`
	Description string `orm:"null;type(text)" json:"description"`
	CreateTime time.Time `orm:"auto_now_add;type(datetime)" json:"create_time"`
	UpdateTime time.Time `orm:"auto_now;type(datetime)" json:"update_time"`
	Labels []*Label 	`orm:"rel(m2m);rel_table(repo_label)" json:"labels"`
	Images []*Image `orm:"reverse(many)" json:"images"`
	Picture string `json:"picture"`
}

func (*imageModel) GetLabelByName(labelName string) (label *Label, err error){
	if labelName == "" {
		beego.Error("label name can not be empty")
		return nil, errors.New("label name empty error")
	}
	l := &Label{Name: labelName}
	if err = Ormer().Read(l, "Name"); err != nil {
		beego.Error(err)
		return nil, errors.New("get label by name error")
	}
	return l, nil
}

func (*imageModel) CreateLabelIfNotExist(labelName string) (label *Label, err error) {
	// check labelName
	if labelName == ""{
		beego.Error("label name can not be empty")
		return nil, errors.New("label name empty error")
	}
	label, err = ImageModel.GetLabelByName(labelName)
	if err == nil {
		// label already exist
		return label, nil
	}else {
		// label not exist
		label = &Label{Name: labelName}
		id, err := Ormer().Insert(label)
		if err != nil {
			beego.Error(err)
			return nil, errors.New("create new label error")
		}
		label.Id = int(id)
		return label, nil
	}
}

func (*imageModel) GetRepoByName(repoName string) (repo *Repository, err error) {
	if repoName == "" {
		beego.Error("repo name can not be empty")
		return nil, errors.New("repository name empty error")
	}
	r := &Repository{Name: repoName}
	if err = Ormer().Read(r, "Name"); err != nil {
		beego.Error(err)
		return nil, errors.New("get repository by name error")
	}
	_, err = Ormer().LoadRelated(r,"labels")
	if err != nil {
		beego.Error(err)
		return nil, errors.New("load labels error")
	}
	_, err = Ormer().LoadRelated(r,"owner")
	if err != nil {
		beego.Error(err)
		return nil, errors.New("load owner error")
	}
	_, err = Ormer().LoadRelated(r,"images")
	if err != nil {
		beego.Error(err)
		return nil, errors.New("load images error")
	}
	return r, nil
}

func (*imageModel) GetRepoById(repoId int) (repo* Repository, err error) {
	if repoId < 0 {
		beego.Error("invalid repoId")
		return nil, errors.New("invalid repo id error")
	}
	r := &Repository{Id: repoId}
	if err = Ormer().Read(r); err != nil {
		beego.Error(err, repoId)
		return nil, errors.New("get repository by id error")
	}
	_, err = Ormer().LoadRelated(r,"labels")
	if err != nil {
		beego.Error(err)
		return nil, errors.New("load labels error")
	}
	_, err = Ormer().LoadRelated(r,"owner")
	if err != nil {
		beego.Error(err)
		return nil, errors.New("load owner error")
	}
	_, err = Ormer().LoadRelated(r,"images")
	if err != nil {
		beego.Error(err)
		return nil, errors.New("load images error")
	}
	return r, nil
}

func (*imageModel) CreateRepository(repo *Repository, user_id int, labels []string) (repo_id int, err error) {
	// check repository name not empty
	if repo.Name == "" {
		beego.Error("repository name can not be empty")
		return -1, errors.New("repository name empty error")
	}

	// check if repository exist
	_, err = ImageModel.GetRepoByName(repo.Name)
	if err == nil {
		beego.Error("repository already exist")
		return -1, errors.New("repository already exist error")
	}

	// insert repository into database
	repo.Images = []*Image{}

	// get user and update author
	user := &User{Id: user_id}
	if err = Ormer().Read(user); err != nil {
		beego.Error("get user by id error")
		return -1, errors.New("get user by id error")
	}
	repo.Owner = user

	// create if not exist labels
	repoLabels := []*Label{}
	for _,labelName := range labels {
		// continue if label name is empty
		if labelName == "" {
			continue
		}
		label, err := ImageModel.CreateLabelIfNotExist(labelName)
		if err != nil {
			beego.Error(err)
			return -1, errors.New("create label if not exist error")
		}
		repoLabels = append(repoLabels, label)
	}
	repo.Labels = repoLabels

	// insert repo
	id, err := Ormer().Insert(repo)
	if err != nil {
		beego.Error(err)
		return -1, errors.New("insert repository error")
	}

	// success
	return int(id), nil
}

func (*imageModel) AddImage(image* Image, repoId int, user_id int) (rtn int, err error) {
	// get repo by repoId
	repo := &Repository{Id: repoId}
	err = Ormer().Read(repo)
	if err != nil {
		beego.Error(err)
		return -1, errors.New("get repo by id error")
	}
	_, err = Ormer().LoadRelated(repo, "owner")
	if err != nil {
		beego.Error(err)
		return -1, errors.New("load owner info error")
	}

	// check author
	user := &User{Id: user_id}
	err = Ormer().Read(user)
	if err != nil {
		beego.Error(err)
		return -1, errors.New("get user by id error")
	}
	if repo.Owner.UserName != user.UserName {
		beego.Error("owner name: ", repo.Owner.UserName, " user name: ", user.UserName)
		return -1, errors.New("no access error")
	}

	// check tag name
	_, err = Ormer().LoadRelated(repo, "Images")
	for _, img := range repo.Images {
		if img.Tag == image.Tag {
			// tag already exist
			beego.Error("tag name already exist")
			return -1, errors.New("tag name already exist")
		}
	}
	// insert image
	image.Repository = repo
	id, err := Ormer().Insert(image)
	if err != nil {
		beego.Error(err)
		return -1, errors.New("insert image error")
	}
	image.Id = int(id)

	// update repo
	repo.Images = append(repo.Images, image)
	_, err = Ormer().Update(repo)
	if err != nil {
		beego.Error(err)
		return -1, errors.New("update repository image list error")
	}

	// success
	return 0, nil
}

func (*imageModel) SearchRepo(key string) (repoList []*Repository, err error){
	var list []*Repository
	_, err = Ormer().QueryTable("repository").Filter("name__icontains", key).All(&list)
	if err != nil {
		beego.Error(err)
		return nil , errors.New("search repository error")
	}
	for _, repo := range list {
		Ormer().LoadRelated(repo, "Owner")
		Ormer().LoadRelated(repo, "Labels")
	}
	return list, nil
}

func (*imageModel) GetAllLabels() (labelList []*Label, err error){
	var list []*Label
	_, err = Ormer().QueryTable("label").All(&list)
	if err != nil {
		beego.Error(err)
		return nil, errors.New("get all labels error")
	}
	return list, nil
}
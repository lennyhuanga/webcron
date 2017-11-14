package controllers

import (
	"github.com/astaxie/beego"
	"strconv"
	"strings"
	"webcron/app/models"
	"webcron/app/libs"
)

type RoleController struct {
	BaseController
}

func (this *RoleController) List() {
	page, _ := this.GetInt("page")
	if page < 1 {
		page = 1
	}

	list, count := models.RoleGetList(page, this.pageSize)

	this.Data["pageTitle"] = "角色列表"
	this.Data["list"] = list
	this.Data["pageBar"] = libs.NewPager(page, int(count), this.pageSize, beego.URLFor("RoleController.List"), true).ToString()
	this.display()
}

func (this *RoleController) Add() {
	if this.isPost() {
		role := new(models.Roles)
		role.RoleName = strings.TrimSpace(this.GetString("role_name"))
		role.UserId = this.userId
		role.Description = strings.TrimSpace(this.GetString("description"))

		_, err := models.RoleAdd(role)
		if err != nil {
			this.ajaxMsg(err.Error(), MSG_ERR)
		}
		this.ajaxMsg("", MSG_OK)
	}

	this.Data["pageTitle"] = "添加分组"
	this.display()
}

func (this *RoleController) Edit() {
	id, _ := this.GetInt("id")

	role, err := models.RoleGetById(id)
	if err != nil {
		this.showMsg(err.Error())
	}

	if this.isPost() {
		role.RoleName = strings.TrimSpace(this.GetString("role_name"))
		role.Description = strings.TrimSpace(this.GetString("description"))
		err := role.Update()
		if err != nil {
			this.ajaxMsg(err.Error(), MSG_ERR)
		}
		this.ajaxMsg("", MSG_OK)
	}

	this.Data["pageTitle"] = "编辑分组"
	this.Data["role"] = role
	this.display()
}

func (this *RoleController) Batch() {
	action := this.GetString("action")
	ids := this.GetStrings("ids")
	if len(ids) < 1 {
		this.ajaxMsg("请选择要操作的项目", MSG_ERR)
	}

	for _, v := range ids {
		id, _ := strconv.Atoi(v)
		if id < 1 {
			continue
		}
		switch action {
		case "delete":
			models.RoleDelById(id)
			//models.TaskResetGroupId(id)
		}
	}

	this.ajaxMsg("", MSG_OK)
}

package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
)

type Roles struct {
	Id          int
	UserId      int
	RoleName   string
	Description string
	CreateTime  int64
}

func (t *Roles) TableName() string {
	return TableName("roles")
}

func (t *Roles) Update(fields ...string) error {
	if t.RoleName == "" {
		return fmt.Errorf("角色名不能为空")
	}
	if _, err := orm.NewOrm().Update(t, fields...); err != nil {
		return err
	}
	return nil
}

func RoleAdd(obj *Roles) (int64, error) {
	if obj.RoleName == "" {
		return 0, fmt.Errorf("角色名不能为空")
	}
	return orm.NewOrm().Insert(obj)
}

func RoleGetById(id int) (*Roles, error) {
	obj := &Roles{
		Id: id,
	}

	err := orm.NewOrm().Read(obj)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

func RoleDelById(id int) error {
	_, err := orm.NewOrm().QueryTable(TableName("roles")).Filter("id", id).Delete()
	return err
}

func RoleGetList(page, pageSize int) ([]*Roles, int64) {
	offset := (page - 1) * pageSize

	list := make([]*Roles, 0)
	query := orm.NewOrm().QueryTable(TableName("roles"))
	total, _ := query.Count()
	query.OrderBy("-id").Limit(pageSize, offset).All(&list)

	return list, total
}

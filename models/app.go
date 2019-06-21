package models

import (
	"time"

	"github.com/astaxie/beego/orm"
)

type App struct {
	AppCode     string    `orm:"column(app_code);size(255)"`
	AppName     string    `orm:"column(app_name);size(255)"`
	CreateDate  time.Time `orm:"column(create_date);type(datetime)"`
	Id          int       `orm:"column(id);auto"`
	PublishDate time.Time `orm:"column(publish_date);type(date);null"`
}

func (t *App) TableName() string {
	return "app"
}

func init() {
	orm.RegisterModel(new(App))
}


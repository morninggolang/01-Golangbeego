package models

import "github.com/astaxie/beego/orm"

type Fruit struct {
	Id int `orm:"column(id):auto"`
	FruitName string `orm:"column(fruit_name)"`
	FruitColor string `orm:"column(fruit_color)"`
}

func  (t *Fruit) TableName() string  {
	return "fruit"
}

func init()  {
	orm.RegisterModel(new(Fruit))
}

package models

import (
    "github.com/astaxie/beego/orm"
)

type ArticleComment struct {
    Base `orm:"-"`
    Id int
    Article *Article `orm:"rel(one)"`
    Content string
}

func init() {
    // Need to register model in init
    orm.RegisterModel(new(ArticleComment))
}

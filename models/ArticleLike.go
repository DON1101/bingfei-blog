package models

import (
    "github.com/astaxie/beego/orm"
)

type ArticleLike struct {
    Base `orm:"-"`
    Id int
    Article *Article `orm:"rel(one)"`
}

func init() {
    // Need to register model in init
    orm.RegisterModel(new(ArticleLike))
}

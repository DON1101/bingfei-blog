package models

import (
    "github.com/astaxie/beego/orm"
)

type ArticleCategory struct {
    Base `orm:"-"`
    Id int
    Name string
}

func init() {
    // Need to register model in init
    orm.RegisterModel(new(ArticleCategory))
}

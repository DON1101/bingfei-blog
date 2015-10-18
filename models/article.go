package models

import (
    "github.com/astaxie/beego/orm"
)

type Article struct {
    Base `orm:"-"`
    Id int
    Title string
    AuthorName string
    Content string
    ArticleCategories []*ArticleCategory `orm:"rel(m2m)"`
}

func init() {
    // Need to register model in init
    orm.RegisterModel(new(Article))
}

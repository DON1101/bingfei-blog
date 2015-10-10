package article

import (
    "github.com/astaxie/beego"
)

type ArticleListController struct {
    beego.Controller
}

func (this *ArticleListController) Get() {
    this.TplNames = "article/article_list.html"
}
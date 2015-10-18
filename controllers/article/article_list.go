package article

import (
    "github.com/astaxie/beego"
    "github.com/astaxie/beego/orm"

    "models"
)

type ArticleListController struct {
    beego.Controller
}

func (this *ArticleListController) Get() {
    var articles []*models.Article
    o := orm.NewOrm()
    num, _ := o.QueryTable("article").All(&articles)
    this.TplNames = "article/article_list.html"
    this.Data["articles"] = articles
    this.Data["article_num"] = num
}
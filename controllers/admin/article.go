package admin

import (
    // "fmt"
    // "github.com/astaxie/beego"
    "github.com/astaxie/beego/orm"

    "models"
    "controllers/auth"
)

type ArticleAdminController struct {
    auth.AuthController
}

func (this *ArticleAdminController) ArticleList() {
    var articles []*models.Article

    o := orm.NewOrm()
    o.QueryTable("article").All(&articles)

    this.TplNames = "article/admin_article_list.html"
    this.Data["articles"] = articles
}

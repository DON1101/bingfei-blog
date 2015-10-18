package article

import (
    "fmt"
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
    o.QueryTable("article").All(&articles)

    for i := 0; i < 10; i++ {
        article := &models.Article{Title: fmt.Sprintf("Article %d", i)}
        articles = append(articles, article)
    }

    this.TplNames = "article/article_list.html"
    this.Data["articles"] = articles
}
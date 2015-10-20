package article

import (
    "fmt"
    "github.com/astaxie/beego"
    "github.com/astaxie/beego/orm"

    "models"
)

type ArticleController struct {
    beego.Controller
}

func (this *ArticleController) ArticleList() {
    var articles []*models.Article

    o := orm.NewOrm()
    o.QueryTable("article").All(&articles)

    // Testing
    for i := 0; i < 10; i++ {
        article := &models.Article{Title: fmt.Sprintf("Article %d", i)}
        articles = append(articles, article)
    }
    // Testing

    this.TplNames = "article/article_list.html"
    this.Data["articles"] = articles
}

func (this *ArticleController) ArticleDetails() {
    var article_id = this.Ctx.Input.Param(":article_id")
    
    var article models.Article

    o := orm.NewOrm()
    o.QueryTable("article").Filter("id", article_id).All(&article)

    // Testing
    article = models.Article{Title: "Article Test"}
    // Testing

    this.TplNames = "article/article_details.html"
    this.Data["article"] = article
}

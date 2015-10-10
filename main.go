package main

import (
    "fmt"
    "github.com/astaxie/beego"
    "github.com/astaxie/beego/orm"
    _ "github.com/lib/pq"

    "controllers"
    "controllers/article"
    _ "models"
)

func init() {
    orm.RegisterDriver("postgres", orm.DR_Postgres)
    orm.RegisterDataBase(
        "default",
        "postgres",
        fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s sslmode=disable",
                   DB_USER,
                   DB_PASSWORD,
                   DB_HOST,
                   DB_PORT,
                   DB_NAME))
}

func main() {
    SyncDB()

    beego.SetStaticPath("/site_media/static/","static/dist/")

    beego.Router("/", &controllers.IndexController{})
    beego.Router("/article/list/", &article.ArticleListController{})

    beego.Run()
}

func SyncDB() {
    // Database alias.
    name := "default"

    // Drop table and re-create.
    force := false

    // Print log.
    verbose := true

    // Error.
    err := orm.RunSyncdb(name, force, verbose)
    if err != nil {
        beego.Error(err)
    }
}


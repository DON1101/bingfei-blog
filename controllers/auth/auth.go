package auth

import (
    "fmt"
    "github.com/astaxie/beego"
    "github.com/astaxie/beego/orm"

    "models"
)

type AuthController struct {
    beego.Controller
}

func Authentication(email string, password string) *models.User {
    var user *models.User

    o := orm.NewOrm()
    o.QueryTable("user").Filter("email", email).All(user)

    beego.Info(fmt.Sprintf("%s", user.Email))
    return user
}

func (this *AuthController) Prepare() {
    sess_username, _ := this.GetSession("username").(string)
    if sess_username == "" {
        this.TplNames = "auth/login.html"
        this.Render()
    } else { 
        this.Ctx.Redirect(302, "/")
    }
}

func (this *AuthController) LoginGet() {
    this.TplNames = "auth/login.html"
}

func (this *AuthController) LoginPost() {
    beego.Info("haha")
    email := this.GetString("email")
    password := this.GetString("password")
    user := Authentication(email, password)

    if user != nil {
        this.SetSession("username", user.UserName)
        this.Redirect("/", 200)
    }
}

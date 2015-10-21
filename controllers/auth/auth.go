package auth

import (
    "github.com/astaxie/beego"
)

type AuthController struct {
    beego.Controller
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

func (this *AuthController) Login() {
    this.TplNames = "auth/login.html"
}

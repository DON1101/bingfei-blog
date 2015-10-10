package defaultTests

import (
    "fmt"
    "os"
    "testing"
    "github.com/astaxie/beego/orm"
    _ "github.com/lib/pq"

    "models"
)

func TestUser(t *testing.T) {
    o := orm.NewOrm()
    o.Using("test")

    profile := new(models.Profile)
    profile.Age = 30

    user := new(models.User)
    user.Profile = profile
    user.Name = "guojing"

    fmt.Println(o.Insert(profile))
    fmt.Println(o.Insert(user))
}

func init() {
    orm.RegisterDataBase("default", os.Getenv("ORM_DRIVER"), os.Getenv("ORM_SOURCE"))
    orm.RegisterDataBase("test", os.Getenv("ORM_DRIVER"), os.Getenv("ORM_SOURCE"))
}
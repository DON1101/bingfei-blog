package test

import (
    "fmt"
    "os"
    "testing"
    "github.com/astaxie/beego/orm"
    _ "github.com/lib/pq"
    _ "models"
)

func TestInit(t *testing.T) {}

func SyncDB() {
    // Database alias.
    name := "test"

    // Drop table and re-create.
    force := false

    // Print log.
    verbose := true

    // Error.
    err := orm.RunSyncdb(name, force, verbose)
    if err != nil {
        fmt.Println(err)
    }
}

func init() {
    orm.RegisterDataBase("default", os.Getenv("ORM_DRIVER"), os.Getenv("ORM_SOURCE"))
    orm.RegisterDataBase("test", os.Getenv("ORM_DRIVER"), os.Getenv("ORM_SOURCE"))
    SyncDB()
}
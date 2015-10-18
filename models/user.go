package models

import (
    "github.com/astaxie/beego/orm"
)

type GenderType int

const (  
   MALE GenderType = 1 + iota  
   FEMALE
)

type User struct {
    Base `orm:"-"`
    Id int
    Email string
    UserName string
    Profile *Profile `orm:"rel(one)"` // OneToOne relation
}

type Profile struct {
    Base `orm:"-"`
    Id int
    Gender GenderType
    Age int16
    FirtName string
    MiddleName string
    LastName string
    User *User `orm:"reverse(one)"` // Reverse relationship (optional)
}

func init() {
    // Need to register model in init
    orm.RegisterModel(new(User), new(Profile))
}

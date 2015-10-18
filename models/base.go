package models

import (
    "time"
)

type Base struct {
    CreateTime time.Time `orm:"auto_now_add;type(datetime)"`
    CreatedBy *User `orm:"rel(fk)"` // RelForeignKey relation
    UpdateTime time.Time `orm:"auto_now;type(datetime)"`
    UpdatedBy *User `orm:"rel(fk)"` // RelForeignKey relation
}

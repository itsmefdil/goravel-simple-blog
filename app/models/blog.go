package models

import (
	"github.com/goravel/framework/database/orm"
)

type Blog struct {
	orm.Model
	Id_category int
	Title       string
	Body        string
	Author      string
	orm.SoftDeletes
}

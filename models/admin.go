package models

import "github.com/goravel/framework/database/orm"

type Admin struct {
	orm.Model
	orm.Timestamps
	Email    string `gorm:"column:email" json:"email"`
	Password string `gorm:"column:password" json:"password"`
}

func (r *Admin) TableName() string {
	return "admins"
}

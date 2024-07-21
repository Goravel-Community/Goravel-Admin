package models

import (
	"crypto/sha256"
	"fmt"

	"github.com/goravel/framework/database/orm"
)

type Admin struct {
	orm.Model
	orm.Timestamps
	Email    string `gorm:"column:email" json:"email"`
	Password string `gorm:"column:password" json:"password"`
}

func (r *Admin) TableName() string {
	return "admins"
}

func (r *Admin) GetAvatar() string {
	h := sha256.New()
	h.Write([]byte(r.Email))
	bs := h.Sum(nil)
	return "https://www.gravatar.com/avatar/" + fmt.Sprintf("%x", bs) + "?d=mp"
}

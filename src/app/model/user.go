package model

import "github.com/guregu/null"

type User struct {
	CreateTime null.Time   `gorm:"column:create_time"`
	Email      null.String `gorm:"column:email"`
	Header     null.String `gorm:"column:header"`
	Password   string      `gorm:"column:password"`
	Phone      string      `gorm:"column:phone"`
	Status     null.Int    `gorm:"column:status"`
	UID        int         `gorm:"column:uid;primary_key"`
	UpdateTime null.Time   `gorm:"column:update_time"`
	Username   string      `gorm:"column:username"`
}

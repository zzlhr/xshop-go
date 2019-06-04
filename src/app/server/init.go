package server

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)
import _err "xshop/src/app/util"

var userName = "root"
var password = "root"
var dbName = "xshop"

var db *gorm.DB

func Init() {
	_db, err := gorm.Open("mysql",
		userName+":"+password+"@/"+dbName+"?charset=utf8&parseTime=True&loc=Local")
	_err.CheckErr(err)
	db = _db
	// 全局禁用表名复数
	db.SingularTable(true)
}

func GetOkJSON(data interface{}) gin.H {
	return gin.H{
		"code": 0,
		"msg": "ok",
		"data": data,
	}
}

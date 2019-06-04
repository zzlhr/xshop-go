package server

import (
	log "github.com/Sirupsen/logrus"
	"github.com/gin-gonic/gin"
	"xshop/src/app/model"
)

/**
获取用户详情
*/
func GetUserInfo(c *gin.Context) {
	uid := c.Param("uid")

	var user model.User
	log.Info(uid)
	log.Info(db)
	db.First(&user, "uid=?", uid).Select("auth_group", "create_time", "email", "header",
		"phone", "status", "uid", "update_time", "username")
	log.Info(user)
	c.JSON(200, GetOkJSON(user))
}

/**
获取用户列表
*/
func GetUserList(c *gin.Context) {
	userName := c.Query("userName")
	phone := c.Query("phone")
	email := c.Query("email")

	var user model.User
	var d = db
	if userName != "" {
		d = d.Where("username like ?", "%"+userName+"%")
	}
	if phone != "" {
		d = d.Where("phone like ?", "%"+phone+"%")
	}
	if email != "" {
		d = d.Where("email like ?", "%"+email+"%")
	}
	d.Find(&user)
	log.Info(user)
	c.JSON(200, GetOkJSON(user))
}

/**
添加用户
*/
func AddUser(c *gin.Context) {
	
}

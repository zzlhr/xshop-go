package http

import (
	"github.com/gin-gonic/gin"
)
import log "github.com/Sirupsen/logrus"

var route *gin.Engine

func Init(r *gin.Engine) {
	route = r
	router()
	userRouter()
}


func router() {
	if route == nil {
		log.Error("route初始化失败")
	}
	route.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"Hello": "World"})
	})
}

package main

import (
	log "github.com/Sirupsen/logrus"
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"os"
	"xshop/src/app/http"
	"xshop/src/app/server"
)

func logInit() {
	log.SetFormatter(&log.TextFormatter{
		DisableColors: true,
		FullTimestamp: true,
	})
	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	log.SetOutput(os.Stdout)

	// Only log the warning severity or above.
	log.SetLevel(log.WarnLevel)
}
func Init() {
	logInit()
}

func main() {
	route := gin.Default()
	http.Init(route)
	server.Init()
	_ = route.Run(":8801")
}

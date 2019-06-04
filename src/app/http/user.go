package http

import "xshop/src/app/server"

func userRouter() {
	userRoute := route.Group("/user")
	userRoute.GET("/userInfo/:uid", server.GetUserInfo)
	userRoute.GET("/users", server.GetUserList)
	userRoute.POST("/add", server.AddUser)
}

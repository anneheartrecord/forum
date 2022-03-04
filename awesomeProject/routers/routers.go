package routers

import (
	"golangstudy/jike/awesomeProject/controllers"
	"golangstudy/jike/awesomeProject/logger"
	"golangstudy/jike/awesomeProject/middleware"

	"github.com/gin-gonic/gin"
)

func Setup(mode string) *gin.Engine {
	if mode == gin.ReleaseMode {
		gin.SetMode(gin.ReleaseMode) //gin设置成发布者模式
	}

	r := gin.New()
	r.Use(logger.GinLogger(), logger.GinRecovery(true))
	//注册业务路由
	v1 := r.Group("/api//v1")
	{
		v1.Use(middleware.JWTAuthMiddleware()) //应用认证中间件
		v1.GET("/community", controllers.CommunityHandler)
		v1.GET("/community/:id", controllers.CommunityDetailHandler)
		v1.POST("/post", controllers.CreatePostHandler)
		v1.GET("/post/:id", controllers.GetPostDetailHandler)
		v1.GET("/posts/", controllers.GetPostListHandler)
		v1.POST("/vote", controllers.PostVoteController)
	}
	r.POST("/signup", controllers.SignUpHandler)
	r.POST("/login", controllers.LoginHandler)
	r.GET("/ping", func(c *gin.Context) {
		//判断是否登录 请求头中是否有token
		isLogin := true
		c.Request.Header.Get("Authorization")
		if isLogin {
			c.String(200, "pong")
		} else {
			c.String(200, "404")
		}
	})
	return r
}

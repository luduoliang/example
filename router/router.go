package router

import (
	"example/controller/admin"
	"example/controller/api"
	"example/middleware"

	"github.com/gin-gonic/gin"
)

//初始化http服务
func InitHttpServer(httpPort string) {
	//New不使用中间间，Default使用中间间
	r := gin.New()

	//r.Use(gin.Logger())
	//r.Use(gin.Recovery())

	//初始化SESSION
	middleware.InitSession(r, "test_cookie_name")
	//视图目录
	r.LoadHTMLGlob("view/**/**/*")
	//静态资源目录
	r.Static("/static", "./static")
	//初始化路由
	InitRouter(r)

	r.Run(":" + httpPort) // listen and serve on 0.0.0.0:8080
}

//初始化路由
func InitRouter(r *gin.Engine) {
	//登录后台
	r.Any("/admin/login", admin.Login)
	r.Any("/admin/logout", admin.Logout)
	r.Any("/api/jwt", api.GetJwtToken)
	//后台路由，使用检查是否登录中间间
	admin_router := r.Group("/admin", middleware.Cors(), middleware.CheckAdminLogin())
	{
		//后台首页
		admin_router.Any("/test", admin.Test)
	}

	//接口路由分组，使用设置跨域中间间
	api_router := r.Group("/api", middleware.Cors(), middleware.JWT())
	{
		//资讯分类列表，参数：
		api_router.Any("/test", api.Test)
	}
}

package router

import (
	"chag/bean"
	"chag/controllers"
	"chag/db"
	"chag/util"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

// CorsMiddleware adds CORS headers to the response.
func CorsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*") // replace "*" with the domain(s) you want to allow
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, PATCH, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}

		c.Next()
	}
}

// AuthMiddleware 检查 cookie 中的 myidentify 字段
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		//获取token
		token := c.GetHeader("Authorization")
		token1 := ""
		//http://localhost/******?token=9mrtzn;ojneb
		qToken := c.Request.URL.RawQuery
		if len(qToken) > 0 && strings.Contains(qToken, "token=") {
			token1 = strings.Split(qToken, "token=")[1]
		}

		if len(token) == 0 && len(token1) == 0 {
			controllers.ReturnError(c, controllers.ERROR, "请登录")
			c.Abort()
			return
		}
		if len(token) == 0 {
			token = token1
		}
		dtoken := util.Decrypt(token)
		info := strings.Split(dtoken, ":")
		if len(info) != 2 {
			controllers.ReturnError(c, controllers.ERROR, "请登录")
			c.Abort()
			return
		}
		username := info[0]
		password := info[1]

		if len(username) == 0 || len(password) == 0 {
			controllers.ReturnError(c, controllers.ERROR, "请登录")
			c.Abort()
			return
		}
		var user bean.User
		if db.DB.Where("username = ?", username).Where("password = ?", password).First(&user).RowsAffected == 0 {
			controllers.ReturnError(c, controllers.ERROR, "请登录")
			c.Abort()
			return
		}

		// 如果 myidentify 符合要求，继续执行后续的 handler 函数
		c.Next()
	}
}

func Router() *gin.Engine {
	r := gin.Default()

	r.Use(CorsMiddleware())

	// 不受保护的路由
	r.POST("/login", controllers.UserController{}.PostLogin)

	// 受保护的路由
	// 使用 AuthMiddleware 作为前置中间件
	auth := r.Group("/")
	auth.Use(AuthMiddleware())
	{
		auth.GET("/users", controllers.UserController{}.GetList)
		user := auth.Group("/user")
		{
			user.GET("/:id", controllers.UserController{}.Get)
			user.POST("/ranking", controllers.UserController{}.PostRanking)
			user.PUT("", controllers.UserController{}.ModifyPassword)
		}

		auth.GET("/contests", controllers.ContestController{}.GetList)
		contest := auth.Group("contest")
		{
			contest.GET("/:id", controllers.ContestController{}.Get)
		}

		problem := auth.Group("/problem")
		{
			problem.GET("/:id", controllers.ProblemController{}.Get)
			problem.POST("", controllers.ProblemController{}.Post)
		}

		submit := auth.Group("/submit")
		{
			submit.GET("/:userid/:problemid", controllers.SubmitController{}.Get)
			submit.POST("", controllers.SubmitController{}.Post)
		}

		auth.GET("/needhelps", controllers.NeedHelpDaysController{}.GetList)
		needhelp := auth.Group("/needhelp")
		{
			needhelp.GET("/:days", controllers.NeedHelpController{}.GetList)
			needhelp.POST("/context", controllers.NeedHelpController{}.PostContext)
			//needhelp.POST("/recontext", controllers.NeedHelpController{}.PostRecontext)
		}

		auth.GET("/recontext/:needhelpid", controllers.RecontextController{}.GetList)
		auth.POST("/recontext", controllers.RecontextController{}.Post)

		auth.GET("/savefiles", controllers.SaveFileController{}.GetList)
		savefile := auth.Group("/savefile")
		{
			savefile.GET("/:filename", controllers.SaveFileController{}.Get)
		}
	}

	return r
}

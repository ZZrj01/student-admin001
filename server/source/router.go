/*
 * @Author: WUMIANHUA
 * @Date: 2021-10-31 16:28:58
 * @LastEditTime: 2021-12-13 11:37:40
 * @Description:
 */
// 声明包
package source

import (
	"fmt"
	"net/http"
	"server/config"
	"server/source/api"
	"strings"

	"github.com/gin-gonic/gin"
	ginsession "github.com/go-session/gin-session"
)

// 创建一个公开方法给main.go 调用
func CreateRouter() {
	router := gin.Default()
	// 调用seesion中间件
	router.Use(ginsession.New())
	// 解决跨域
	router.Use(Cors())
	// 分组路由
	web := router.Group("/api")
	{
		web.POST("/login", api.Login)
		// 添加班級 api 接口 192.168.0.1/api/class
		web.Use(hasLogin()) //只有下面的方法需要使用该组件
		//退出班级接口
		web.POST("/logout", api.LogOut)
		web.POST("/addclass", api.AddClass)
		web.POST("/class", api.GetClass)
		web.POST("/delclass", api.DelClass)
		web.POST("/modifyclass", api.ModifyClass)

		//添加学生 api 接口 api/student
		web.POST("/addstudent", api.AddStudent)
		web.POST("/student", api.GetStudent)
		web.POST("/delstudent", api.DelStudent)
		web.POST("/modifystudent", api.ModifyStudent)

		//添加考勤 api 接口 api/attendance
		web.POST("/addattendance", api.AddAttendance)
		web.POST("/attendance", api.GetAttendance)
		web.POST("/delattendance", api.DelAttendance)
		// web.POST("/modifyattendance", api.ModifyAttendance)
	}
	// 监听端口
	router.Run(fmt.Sprintf(":%d", config.Conf.Server.Port))

}

/**
 * @description: 解决跨域中间件
 * @param {*}
 * @return {*}
 */
func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		c.Header("Access-Control-Allow-Origin", c.Request.Header.Get("Origin"))
		c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
		c.Header("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, Authorization")
		c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Cache-Control, Content-Language, Content-Type")
		c.Header("Access-Control-Allow-Credentials", "true")
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}
		c.Next()
	}
}
func hasLogin() gin.HandlerFunc {
	return func(c *gin.Context) {
		url := c.Request.RequestURI
		//如果是登录的接口就放行
		if strings.Contains(url, "/login") {
			c.Next()
			return
		}
		//判断是否有token
		token := c.GetHeader("Authorization")
		if token == "" {
			c.JSON(200, map[string]interface{}{
				"code": 1008,
			})
			c.Abort()
			return
		}
		// 判断session里是否有该用户
		store := ginsession.FromContext(c)
		obj, b := store.Get(token)
		fmt.Println(obj)
		//如果b false 断定是没登录
		if !b {
			c.JSON(200, map[string]interface{}{
				"code": 1008,
			})
			c.Abort()
		}
		//通过gin传递参数
		c.Set("token", token)
		c.Next()
	}
}

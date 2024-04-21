package server

import (
	"api2/server/middlewares"
	"api2/server/service"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	//一些基础配置
	config := cors.DefaultConfig()
	config.ExposeHeaders = []string{"Authorization"}
	config.AllowCredentials = true
	config.AllowAllOrigins = true
	config.AllowHeaders = []string{"Origin", "Content-Length", "Content-Type", "Authorization"}
	r.Use(cors.New(config))

	api := r.Group("api")
	api.Use(gin.Recovery())
	api.Use(middlewares.Logger())
	{
		//学生信息查询接口
		student := api.Group("student") 
		{
			// GET /api/student?name=?&page=?&page_size=?&birth_start=?&birth_end=? | 查询学生信息
			student.GET("", service.HandlerBindQuery(&service.StudentService{}))
		}
	}
	return r
}

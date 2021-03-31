package router

import (
	"ekgo/app/admin/admins"
	"ekgo/app/middleware"
	"github.com/gin-gonic/gin"
)

//后台管理
func Admin(Router *gin.RouterGroup) {

	Group := Router.Group("/admin").Use(middleware.AdminAuth())
	{
		//管理员
		Group.GET("/admins", admins.Index)         //分页
		Group.GET("/admins/:id", admins.Show)      //修改管理员
		Group.POST("/admins", admins.Store)        //添加管理员
		Group.DELETE("/admins/:id", admins.Delete) //删除管理员
		Group.PUT("/admins/:id", admins.Update)    //修改管理员
	}

}

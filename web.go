package main

import (
	"github.com/gin-gonic/gin"
	"sample/movies"
	"sample/auth"
)

var (
	r *gin.Engine
)

// InitTables 初始化不同模块的表格
func InitTables() {
	auth.InitTables() // 用户管理，登录验证模块
	movies.InitTables() //功能模块
}

// Register 注册不同模块的路由
func Register() {
	// 包含登录验证中间件
	auth.Register(r)
	// 功能路由
	movies.Register(r)
}

func main() {

	// 创建一个默认的无路由引擎
	r = gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	InitTables()
	Register()
	
	r.Run(":80")
}

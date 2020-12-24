package auth

import (
	"github.com/gin-gonic/gin"
)

// Register 注册路由函数...
func Register(r *gin.Engine) {
	// users
	r.POST(LOGIN, login)
	r.POST(REGISTER, register)
	// 注册路由中间件 校验cookies 对之后的路由实行校验规则
	r.Use(VerifyLoginRouterMiddle())

}

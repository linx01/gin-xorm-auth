package movies

import (
	"github.com/gin-gonic/gin"
)

// Register 注册路由函数...
func Register(r *gin.Engine) {
	// movie
	r.GET(MOVIE, getMovie)
	r.POST(MOVIE, postMovie)
	r.PUT(MOVIE, putMovie)
	r.DELETE(MOVIE, deleteMovie)

}

package auth

import (
	"net/http"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/gin-gonic/gin"

)

// login 验证用户，返回cookie
func login(c *gin.Context) {

	db := initDB()
	redisDB := InitRedis()

	// 关闭连接
	defer db.Close()
	defer redisDB.Close()

	var (
		username    string
		password    string
		userInfo    UserInfo
		session     string
	)

	// 获取参数 POST /Movie
	username = c.DefaultPostForm("username", "")
	password = c.DefaultPostForm("password", "")


	if username == "" || password == ""{
		c.JSON(http.StatusOK, gin.H{
			"status":  "error",
			"message": "please enter the username or the password",
		})
	} else {
		// 搜索用户名和密码 如果吻合则登录成功 设置cookie给请求头
		userInfo = UserInfo{}
		db.Where("username=?", username).First(&userInfo)
		// 没搜索到 返回空白
		if userInfo.UserName == ""{
			c.JSON(http.StatusOK, gin.H{
				"status":  "error",
				"message": "there is no user called "+username,
			})
		}else if md5V(password) != userInfo.PassWord {
			c.JSON(http.StatusOK, gin.H{
				"status":  "error",
				"message": "wrong password, auth failed",
			})
		}else{
			// 登录成功
			session = md5V(password)
			// 更新浏览器的cookies与后端session的有效时长
			UpdateCookies(session, c, redisDB)
			c.JSON(http.StatusOK, gin.H{
				"status":  "ok",
				"message": "login ok",
			})
		}

	}

}

// register 注册
func register(c *gin.Context) {

	db := initDB()
	// 关闭连接
	defer db.Close()

	var (
		username    string
		password string
		userinfo UserInfo
	)

	// 获取参数 POST /Movie
	username = c.DefaultPostForm("username", "")
	password = c.DefaultPostForm("password", "")


	if username == "" || password == ""{
		c.JSON(http.StatusOK, gin.H{
			"status":  "error",
			"message": "please enter the username or the password",
		})
	} else {
		// 插入数据
		// password md5加密
		userinfo = UserInfo{UserName: username, PassWord: md5V(password)} // 生成结构体

		result := db.Create(&userinfo) // 通过数据的指针来创建
		
		if result.Error != nil { // 返回 error
			c.JSON(http.StatusOK, gin.H{
				"status":  "error",
				"message": "register error:" + result.Error.Error(),
			})
		}
		c.JSON(http.StatusOK, gin.H{
			"status":  "ok",
			"message": "register ok, now please login the system by your register info",
		})

	}

}


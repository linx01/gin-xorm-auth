package auth

import (
    "github.com/jinzhu/gorm"
	"fmt"
    "crypto/md5"
	"encoding/hex"
	"github.com/go-redis/redis"
	"time"
	"github.com/gin-gonic/gin"
	"net/http"
)

//  md5V 加密
func md5V(str string) string  {
    h := md5.New()
    h.Write([]byte(str))
    return hex.EncodeToString(h.Sum(nil))
}

// SetDataInToRedis ...
func SetDataInToRedis(x string, y interface{}, time time.Duration, redisDB *redis.Client) error {
	err := redisDB.Set(x, y, time).Err()
	return err
}

// GetDataFromRedis ...
func GetDataFromRedis(x string, redisDB *redis.Client) (string, error) {
	res, err := redisDB.Get(x).Result()
	return res, err
}


// updateCookies 更新浏览器的cookies与后端session的有效时长
func UpdateCookies(session string, c *gin.Context, redisDB *redis.Client){
	// 给请求头设置session
	c.SetCookie("session", session, EXPIREDTIME, "/", "localhost", false, true)
	// 并将cookie缓存
	SetDataInToRedis(session, 1, time.Second*EXPIREDTIME, redisDB)
}

// InitRedis 初始化redis连接
func InitRedis() *redis.Client{

	var err error
	var client *redis.Client

	// 加入循环 如果无法连接 则无法往下走业务逻辑
	for {
			client = redis.NewClient(&redis.Options{
					Addr:     "localhost:6379",
					Password: "",
					DB:       1,
			})
			_, err = client.Ping().Result()
			if err != nil {
				panic("init redis error:" + err.Error())
			}
			break
	}

	return client
}


// initDB 初始化数据库连接
func initDB() *gorm.DB {
	// 初始化数据库
	dsn := USERNAME + ":" + PASSWORD + "@(127.0.0.1:3306)/sample?charset=utf8"
	db, err := gorm.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}
	return db
}

// inittable 初始化表格
func InitTables() {
	db := initDB()
	defer db.Close()
	// 表名
	db.Table("users_info").CreateTable(&UserInfo{})
	// 自动迁移
	db.AutoMigrate(&UserInfo{})
	// 初始化管理员用户名和密码
	userinfo := UserInfo{UserName: ADMINNAME, PassWord: md5V(ADMINPASSWORD)}
	result := db.Create(&userinfo)
	if result.Error != nil{
		fmt.Println(result.Error.Error())
	}

}

// VerifyRouterMiddle 校验登录中间件
/*
拿取请求头cookie并判断：
如果无cookie，返回请登录；
如果cookie过期，返回提示信息；
如果cookie没有过期，刷新客户端和缓存中的cookie缓存时长
*/
func VerifyLoginRouterMiddle() gin.HandlerFunc{
	return func(c *gin.Context) {
		redisDB := InitRedis()
		// 获取请求cookie
		sessionValue, err := c.Cookie("session")
		if err != nil{
			// 无cookie，不再调用后面的函数处理，返回请登录
			c.Abort()
			c.JSON(http.StatusOK, gin.H{
				"status":"error",
				"message":"please login.",
			})
		}else if _,erR := GetDataFromRedis(sessionValue,redisDB);erR != nil{
			// cookie过期，不再调用后面的处理函数，重新登录
			c.Abort()
			c.JSON(http.StatusOK, gin.H{
				"status":"error",
				"message":"your account's session is expired.please relogin to access.",
			})
		}else if erR == nil{
			// 没过期，更新cookie
			UpdateCookies(sessionValue, c, redisDB)
			// 继续调用后面的处理函数
			c.Next()
		}
	}
}

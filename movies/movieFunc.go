package movies

import (
	"net/http"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/gin-gonic/gin"
)

// getMovie 查询
func getMovie(c *gin.Context) {

	db := initDB()
	// 关闭连接
	defer db.Close()

	var (
		info     map[string]map[string]string
		detail   map[string]string
		movieName string
		movieInfo MovieInfo
		movieInfos []MovieInfo
	)

	info = make(map[string]map[string]string)

	
	// 获取参数 GET /Movie?name=xxx
	movieName = c.DefaultQuery("name", "")

	if movieName == "" {
		// 查询所有
		db.Find(&movieInfos)
		if len(movieInfos) != 0{
			for _, m := range movieInfos{
				detail = make(map[string]string)
				detail["nation"] = m.Nation
				detail["directior"] = m.Director
				info[m.Name] = detail
			}
			c.JSON(http.StatusOK, gin.H{
				"status":  "ok",
				"message": info,
			})
		}else{
			c.JSON(http.StatusOK, gin.H{
				"status":  "ok",
				"message": "empty",
			})
		}

	} else {
		// 查询单个
		db.Where("name=?", movieName).First(&movieInfo)
		if movieInfo.Name != ""{
			detail = make(map[string]string)
			detail["nation"] = movieInfo.Nation
			detail["directior"] = movieInfo.Director
			info[movieName] = detail
			c.JSON(http.StatusOK, gin.H{
				"status":  "ok",
				"message": info,
			})
		}else{
			c.JSON(http.StatusOK, gin.H{
				"status":  "ok",
				"message": "there is no result for movie " + movieName,
			})
		}
	}
}

// postMovie 增加
func postMovie(c *gin.Context) {

	db := initDB()
	// 关闭连接
	defer db.Close()

	var (
		name    string
		nation string
		director  string
		movie MovieInfo
	)

	// 获取参数 POST /Movie
	name = c.DefaultPostForm("name", "")
	nation = c.DefaultPostForm("nation", "")
	director = c.DefaultPostForm("director", "")

	if name == "" {
		c.JSON(http.StatusOK, gin.H{
			"status":  "error",
			"message": "please enter the Movie name",
		})
	} else {
		// 插入数据

		movie = MovieInfo{Name: name, Nation: nation, Director: director} // 生成结构体

		result := db.Create(&movie) // 通过数据的指针来创建
		
		if result.Error != nil { // 返回 error
			c.JSON(http.StatusOK, gin.H{
				"status":  "error",
				"message": result.Error.Error(),
			})
		}
		c.JSON(http.StatusOK, gin.H{
			"status":  "ok",
			"message": "insert ok",
		})

	}

}

// putMovie 修改
func putMovie(c *gin.Context) {

	db := initDB()
	// 关闭连接
	defer db.Close()

	var (
		name    string
		director string
		nation  string
		movieInfo MovieInfo
	)

	// 获取参数 PUT /Movie
	name = c.DefaultPostForm("name", "")
	nation = c.DefaultPostForm("nation", "")
	director = c.DefaultPostForm("director", "")

	if name == "" {
		c.JSON(http.StatusOK, gin.H{
			"status":  "error",
			"message": "please enter the Movie name",
		})
	} else {
		// 更新数据
		db.Where("name=?", name).First(&movieInfo)
		if movieInfo.Name == ""{
			c.JSON(http.StatusOK, gin.H{
				"status":  "error",
				"message": "there is no result for movie "+name,
			})
		}else{
			// 更新所有字段
			movieInfo.Nation = nation
			movieInfo.Director = director
			result := db.Save(&movieInfo)

			if result.Error != nil {
				c.JSON(http.StatusOK, gin.H{
					"status":  "error",
					"message": result.Error.Error(),
				})
			}
			c.JSON(http.StatusOK, gin.H{
				"status":  "ok",
				"message": "edit ok",
			})
		}


	}

}

// deleteMovie 删除
func deleteMovie(c *gin.Context) {

	db := initDB()
	// 关闭连接
	defer db.Close()

	var (
		name string
		movieInfo MovieInfo
	)

	// 获取参数 DELETE /Movie
	name = c.DefaultPostForm("name", "")

	if name == "" {
		c.JSON(http.StatusOK, gin.H{
			"status":  "error",
			"message": "please enter the Movie name",
		})
	} else {
		// 删除数据
		db.Where("name=?", name).First(&movieInfo)
		if movieInfo.Name == ""{
			c.JSON(http.StatusOK, gin.H{
				"status":  "error",
				"message": "there is no result for movie "+name,
			})
		}else{
			result := db.Delete(&movieInfo)
			if result.Error != nil {
				c.JSON(http.StatusOK, gin.H{
					"status":  "error",
					"message": result.Error.Error(),
				})
			}
			c.JSON(http.StatusOK, gin.H{
				"status":  "ok",
				"message": "delete ok",
			})
		}
	}

}
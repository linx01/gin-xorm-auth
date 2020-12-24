package movies

import (
    "github.com/jinzhu/gorm"
	// _ "github.com/jinzhu/gorm/dialects/mysql"
)

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
	db.Table("movies_info").CreateTable(&MovieInfo{})
	// 表名
	db.Table("heros_info").CreateTable(&HeroInfo{})
	// 自动迁移
	db.AutoMigrate(&MovieInfo{})
	// 自动迁移
	db.AutoMigrate(&HeroInfo{})
}

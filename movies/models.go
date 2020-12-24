package movies


// MovieInfo ...
type MovieInfo struct {
	ID       int64  `gorm:"column:id;primary_key;AUTO_INCREMENT"` // column:id 列名
	Name     string `gorm:"column:name;type:varchar(100);index:name;unique;not_null"` // index:name 设置索引 type:varchar(100) 大小
	Nation   string `gorm:"column:nation"`
	Director string `gorm:"column:director"`
}

// HeroInfo ...
type HeroInfo struct {
	ID          int64  `gorm:"column:id;primary_key;AUTO_INCREMENT"` // column:id 列名
	Name        string `gorm:"column:name;type:varchar(100);index:name;unique;not_null"` // index:name 设置索引 type:varchar(100) 大小
	MovieName   string `gorm:"column:movie_name;not_null;association_foreignkey:movies_info.name"`
}
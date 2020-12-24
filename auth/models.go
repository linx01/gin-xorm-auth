package auth


// UserInfo ...
type UserInfo struct {
	ID         int64  `gorm:"column:id;primary_key;AUTO_INCREMENT"` // column:id 列名
	UserName   string `gorm:"column:username;type:varchar(100);index:name;unique;not_null"` // index:name 设置索引 type:varchar(100) 大小
	PassWord   string `gorm:"column:password;type:varchar(100);not_null"`
}

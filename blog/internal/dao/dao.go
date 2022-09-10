package dao

/*
	数据访问层(Database Access Object)
	所有与数据相关的操作都会在 dao 层进行, 例如 MySQL、ElasticSearch 等。
*/

import "github.com/jinzhu/gorm"

type Dao struct {
	engine *gorm.DB
}

func New(engine *gorm.DB) *Dao {
	return &Dao{engine: engine}
}

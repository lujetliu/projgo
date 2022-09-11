package model

import "blog/global"

func Migrate() {
	global.DBEngine.AutoMigrate(&Tag{})
	global.DBEngine.AutoMigrate(&Article{})
	global.DBEngine.AutoMigrate(&ArticleTag{})
}

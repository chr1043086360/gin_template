package models

import (
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql" // 这里是导进来不用，只是调用mysql包里的init函数
)

// 数据库的单例
var DB *gorm.DB

func Datebase(connString string) {
	db, err := gorm.Open("mysql", connString)

	//db.LogMode(true)
	if err != nil {
		// panic("failed to connect database")
		fmt.Print(err)
	}
	db.DB().SetMaxIdleConns(50)
	db.DB().SetMaxOpenConns(100)
	db.DB().SetConnMaxLifetime(time.Second * 30)
	DB = db

	migration()

}


func migration() {
	// 自动迁移模式
	DB.AutoMigrate(&User{})
}

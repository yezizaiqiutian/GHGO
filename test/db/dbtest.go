package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"log"
	"os"
	"time"
)

func main() {
	dbinit()
	//dbinsert()
	dbselect()
}

var (
	DB *gorm.DB
)

//插入
func dbinsert() {
	info := UserInfo{UserName: "test002",PassWord: "111112",Phone: "123456"}
	DB.Create(&info)
}

//查询
func dbselect() {
	var userlist []UserInfo

	//var newslist []bean.News = make([]int, 0)
	DB.Find(&userlist)

	fmt.Println(userlist)
}

//建立链接
func dbinit() {
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: time.Second, // 慢 SQL 阈值
			LogLevel:      logger.Info, // Log level
			Colorful:      true,        // 禁用彩色打印
		},
	)

	//dsn := "root:gh426486@tcp(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"
	dsn := "root:123456@tcp(82.157.165.250:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"
	// 全局模式
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
		Logger: newLogger,
	})
	if err != nil {
		panic(err)
	}
}

//建表
func dbcreate() {
	_ = DB.AutoMigrate(&UserInfo{})
}
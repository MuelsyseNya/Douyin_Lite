package repository

import (
	"fmt"
	"github.com/Muelsyse/Douyin_Lite/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func InitDatabase() {

	// Configure MySQL connection parameters
	username := "root"      //账号
	password := "123456"    //密码
	host := "127.0.0.1"     //数据库地址，可以是Ip或者域名
	port := 3306            //数据库端口
	Dbname := "douyin_lite" //数据库名
	timeout := "10s"        //连接超时，10秒

	// dsn := "root:Mdmdfuck123..@tcp(114.132.198.20:3306)/douyin_demo?charset=utf8mb4&parseTime=True&loc=Local"

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local&timeout=%s", username, password, host, port, Dbname, timeout)

	DB, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		SkipDefaultTransaction: true, // 禁用默认事务操作来提升性能
	})
	// 数据库链接错误
	if err != nil {
		panic(err)
	}

	// 数据库自动迁移，通过gorm自动创建表
	err = DB.AutoMigrate(&model.User{}, &model.Video{}, &model.Comment{}, &model.Follow{}, &model.Favorite{}, &model.Message{})
	if err != nil {
		panic(err)
	}
}

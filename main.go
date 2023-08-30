package main

import (
	"github.com/Muelsyse/Douyin_Lite/repository"
	"github.com/Muelsyse/Douyin_Lite/service"
	"github.com/gin-gonic/gin"
)

func main() {
	go service.RunMessageServer()

	// 初始化数据库
	repository.InitDatabase()

	r := gin.Default()

	initRouter(r)

	err := r.Run()
	if err != nil {
		return
	} // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

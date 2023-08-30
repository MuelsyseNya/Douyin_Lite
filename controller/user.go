package controller

import (
	"github.com/Muelsyse/Douyin_Lite/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

// Register 用户注册
func Register(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")

	result, err := service.NewUser().Register(username, password)
	if err != nil {
		// 发生错误，不返回用户ID和Token
		c.JSON(http.StatusOK, result.Response)
		return
	}

	// 成功返回
	c.JSON(http.StatusOK, result)
	return
}

func Login(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")

	result, err := service.NewUser().Login(username, password)
	if err != nil {
		// 发生错误，不返回用户ID和Token
		c.JSON(http.StatusOK, result.Response)
		return
	}

	// 成功返回
	c.JSON(http.StatusOK, result)
	return
}

func UserInfo(c *gin.Context) {

}

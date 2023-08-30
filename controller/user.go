package controller

import (
	"github.com/Muelsyse/Douyin_Lite/response"
	"github.com/Muelsyse/Douyin_Lite/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// Register 用户注册
func Register(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")

	result := service.NewUser().Register(username, password)

	// 返回响应
	c.JSON(http.StatusOK, result)
	return
}

func Login(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")

	result := service.NewUser().Login(username, password)

	// 返回响应
	c.JSON(http.StatusOK, result)
	return
}

func UserInfo(c *gin.Context) {

	// 需要查询的用户的ID
	toUserID, err := strconv.ParseInt(c.Query("user_id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusOK, response.UserResponse{
			Response: response.ErrParseInt,
		})
	}
	// 获取当前用户的ID
	UserId, ok := c.Keys["current_id"].(int64)
	if !ok {
		c.JSON(http.StatusOK, response.UserResponse{
			Response: response.ErrGetUserID,
		})
	}
	result := service.NewUser().QueryUserInfo(toUserID, UserId)

	// 返回响应
	c.JSON(http.StatusOK, result)
	return
}

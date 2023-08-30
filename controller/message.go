package controller

import (
	//"fmt"
	"github.com/gin-gonic/gin"
)

var tempChat = map[string][]Message{}

var messageIdSequence = int64(1)

type ChatResponse struct {
	Response
	MessageList []Message `json:"message_list"`
}

// MessageAction no practical effect, just check if token is valid
func MessageAction(c *gin.Context) {

}

// MessageChat all users have same follow list
func MessageChat(c *gin.Context) {

}

func genChatKey(userIdA int64, userIdB int64) string {
	return ""
}

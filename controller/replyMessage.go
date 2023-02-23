package controller

import (
	"Bing-QQBot/config"
	"Bing-QQBot/util"
	"fmt"
	"log"
	"net/http"
	"strings"
)

//	func reply(context *gin.Context, content string) {
//		context.JSON(http.StatusOK, gin.H{
//			"reply": content,
//		})
//	}
type ReplySender struct {
	// 把发送消息作为一个类进行封装
	// todo
}

func ReplyPrivateMessage(message string, userID string) {
	url := fmt.Sprintf("http://localhost:%s/send_private_msg?user_id=%s&message=%s", config.MyConfig.Bot, userID, util.HandlerAnswer(message))
	_, err := http.Get(url)
	if err != nil {
		log.Printf("Send Error:%v", err)
	}

}
func ReplyGroupMessage(message string, groupID string, senderID string) {
	answer := util.HandlerAnswer(message)
	url := fmt.Sprintf("http://localhost:%s/send_group_msg?group_id=%v&message=[CQ:at,qq=%s]%s", config.MyConfig.Bot, groupID, senderID, answer)
	_, err := http.Get(url)
	if err != nil {
		fmt.Printf("Send Error:%v", err)
	}
}
func ReplyNotice(message string, groupID string, senderID string) {
	message = strings.Replace(message, " ", "", -1)
	message = strings.TrimSpace(message)

	if groupID != "" {
		url := fmt.Sprintf("http://localhost:%s/send_group_msg?group_id=%v&message=%s", config.MyConfig.Bot, groupID, message)
		log.Println(url)
		_, err := http.Get(url)
		if err != nil {
			fmt.Printf("Send Error:%v", err)
		}
	} else {
		url := fmt.Sprintf("http://localhost:%s/send_private_msg?user_id=%s&message=%s", config.MyConfig.Bot, senderID, message)
		_, err := http.Get(url)
		if err != nil {
			log.Printf("Send Error:%v", err)
		}
	}

}

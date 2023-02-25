package router

import (
	"Bing-QQBot/config"
	"Bing-QQBot/controller"
	"Bing-QQBot/model"
	"Bing-QQBot/util"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/tidwall/gjson"
	"io"
	"log"
	"net/http"
)

var R *gin.Engine
var flag = false

func NewRouter() {

	R = gin.Default()
	//gin.SetMode(gin.ReleaseMode)
	R.POST("/", Message)
	R.POST("/answer", Answer)
	R.GET("/question", Question)
	url := fmt.Sprintf("0.0.0.0:%s", config.MyConfig.Bridge)
	_ = R.Run(url)
}

var myMessage model.MessageStruct

func Message(c *gin.Context) {
	dataReader := c.Request.Body
	rawData, _ := io.ReadAll(dataReader)
	stringData := string(rawData)
	postType := gjson.Get(stringData, "post_type").String()
	if postType == "message" {
		log.Println(stringData)
		myMessage.GroupID = gjson.Get(stringData, "group_id").String()
		myMessage.SelfQQ = gjson.Get(stringData, "self_id").String()
		myMessage.Message = gjson.Get(stringData, "message").String()
		myMessage.Sender = gjson.Get(stringData, "sender.nickname").String()
		myMessage.SenderQQ = gjson.Get(stringData, "sender.user_id").String()
		myMessage.MessageType = gjson.Get(stringData, "message_type").String()
		myMessage.GroupName = util.GetGroupInfo(myMessage.GroupID)
		if myMessage.MessageType == "private" {
			log.Printf("%s发送了信息：%v", myMessage.Sender, myMessage.Message)
		} else if myMessage.MessageType == "group" {
			log.Printf("群“%s”中%s向你发送了信息：%v", myMessage.GroupName, myMessage.Sender, myMessage.Message)
		}

	}
}
func Question(c *gin.Context) {
	question := myMessage.Message
	messageType := myMessage.MessageType
	self := myMessage.SelfQQ
	// 也许需要在这判断flag的值防止创建过多对话，但也许能实现同时多个会话，问题在于还需要获得对话id
	if !flag {
		log.Println("Haven‘t begining to chat ")
		start := controller.StartBingBot(myMessage.MessageType, myMessage.Message, myMessage.GroupID, myMessage.SenderQQ, myMessage.SelfQQ)
		if start == "!start" {
			c.JSON(http.StatusOK, gin.H{
				"question": start,
				"type":     messageType,
				"self":     self,
			})
			flag = true
		} else {
			c.JSON(http.StatusOK, gin.H{
				"question": start,
				"type":     messageType,
				"self":     self,
			})
		}

	} else {
		if util.HandlerQuesiton(question) == "!exit" {
			flag = false
		}
		c.JSON(http.StatusOK, gin.H{
			"question": question,
			"type":     messageType,
			"self":     self,
		})
	}

}

func Answer(c *gin.Context) {
	dataReader := c.Request.Body
	rawData, _ := io.ReadAll(dataReader)
	answer := string(rawData)
	log.Println(answer)
	if answer == "bing正在生成回答" {
		controller.ReplyNotice(answer, myMessage.GroupID, myMessage.SenderQQ)
	} else {
		if myMessage.MessageType == "private" {
			controller.ReplyPrivateMessage(answer, myMessage.SenderQQ)
		}
		//} else if myMessage.MessageType == "group" {
		//	controller.ReplyGroupMessage(answer, myMessage.GroupID, myMessage.SenderQQ)
		//}
	}

}

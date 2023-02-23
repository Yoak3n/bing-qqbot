package controller

import (
	"Bing-QQBot/util"
)

func StartBingBot(messageType string, question string, groupID string, senderID string, self string) (startMessage string) {
	if messageType == "group" && util.CheckGroupAtMe(question, self) {
		question = util.HandlerQuesiton(question)
		if question == "和bing聊天" || question == "与bing聊天" || question == "chatwithbing" {
			ReplyNotice("开始与bing聊天吧!", groupID, "")
		}
	} else {
		if question == "和bing聊天" || question == "与bing聊天" || question == "chatwithbing" {
			ReplyNotice("开始与bing聊天吧!", "", senderID)
		}
	}
	startMessage = "!start"
	return
}

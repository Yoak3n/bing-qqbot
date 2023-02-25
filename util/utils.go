package util

import (
	"fmt"
	"github.com/tidwall/gjson"
	"io"
	"net/http"
	"regexp"
	"strconv"
	"strings"
)

func GetGroupInfo(groupID string) (GroupName string) {
	baseURL := "http://127.0.0.1:5700"
	getURL := fmt.Sprintf("%s/get_group_info?group_id=%s", baseURL, groupID)
	res, _ := http.Get(getURL)
	jsonByte, _ := io.ReadAll(res.Body)
	jsonString := string(jsonByte)
	GroupName = gjson.Get(jsonString, "data.group_name").String()
	return GroupName
}

func CheckGroupAtMe(message string, selfQQ string) (flag bool) {
	if strings.Contains(message, "[CQ:") {
		messageWithCQ := strings.Split(message, " ")
		if messageWithCQ[0] == fmt.Sprintf("[CQ:at,qq=%s]", selfQQ) {
			flag = true
		} else {
			flag = false
		}
	}
	return flag
}

func HandlerQuesiton(message string) (question string) {
	re, _ := regexp.Compile(`\[\S*\]`)
	finds := re.FindAllString(message, -1)
	if len(finds) != 0 {
		for _, find := range finds {
			message = strings.Replace(message, find, "", -1)
		}
	}
	return strings.TrimSpace(message)
}

func HandlerAnswer(message string) (answer string) {
	answer = strconv.Quote(message)
	answer = answer[1 : len(answer)-1]
	answer = strings.Replace(answer, " ", "%20", -1)
	answer = strings.Replace(answer, `\n`, "%0A", -1)
	answer = strings.Replace(answer, `\`, "", -1)
	re, _ := regexp.Compile(`\[\^\d*\^\]`)
	finds := re.FindAllString(answer, -1)
	if len(finds) != 0 {
		for _, find := range finds {
			answer = strings.Replace(answer, find, "", -1)
		}
	}

	return
}

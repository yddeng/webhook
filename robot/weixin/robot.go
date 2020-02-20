package weixin

import (
	"fmt"
	"github.com/yddeng/webhook/conf"
	"github.com/yddeng/webhook/util"
)

/*
 企业微信机器人

{
    "msgtype": "text",
    "text": {
        "content": "广州今日天气：29度，大部分多云，降雨概率：60%",
        "mentioned_list":["wangqing","@all"],
        "mentioned_mobile_list":["13800001111","@all"]
    }
}

{
    "msgtype": "markdown",
    "markdown": {
        "content": "实时新增用户反馈<font color=\"warning\">132例</font>，请相关同事注意。\n
         >类型:<font color=\"comment\">用户反馈</font>
         >普通用户反馈:<font color=\"comment\">117例</font>
         >VIP用户反馈:<font color=\"comment\">15例</font>"
    }
}
*/

type Message struct {
	MsgType  string            `json:"msgtype"`
	Text     map[string]string `json:"text"`
	Markdown map[string]string `json:"markdown"`
}

func SendToClient(name, msg string) {
	req := Message{MsgType: "markdown",
		Markdown: map[string]string{"content": msg}}
	//fmt.Println(req)

	robots := conf.GetConfig().Robot
	for _, r := range robots {
		if r.Name == name {
			resp, err := util.PostJson(r.Url, req, 0)
			if err != nil {
				fmt.Printf("sendToClient name:%s err:%s\n", r.Name, err)
				continue
			}
			if resp.StatusCode != 200 {
				fmt.Printf("sendToClient name:%s code:%d\n", r.Name, resp.StatusCode)
				continue
			}
			fmt.Printf("sendToClient name:%s ok\n", r.Name)
		}
	}

}

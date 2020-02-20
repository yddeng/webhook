package weixin

import (
	"fmt"
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
*/

type Message struct {
	MsgType string            `json:"msgtype"`
	Text    map[string]string `json:"text"`
}

func SendToClient(msg string) {
	req := Message{MsgType: "text",
		Text: map[string]string{"content": msg}}
	fmt.Println(req)

	/*robots := conf.GetConfig().Robot
	for _, r := range robots {
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
	*/
}

package workweixin

import (
	"fmt"
	"github.com/yddeng/webhook/core/event"
	"github.com/yddeng/webhook/core/robot"
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

var (
	push_tmp = `**%s**  通知:
<font color="info">%s</font> 推送了提交到 <font color="info">%s</font> 分支。`

	merge_open_tmp = `**%s**  通知:
<font color="info">%s</font> 创建了从 <font color="info">%s</font> 到 <font color="info">%s</font> 的合并请求。`

	merge_close_tmp = `**%s**  通知:
<font color="info">%s</font> 关闭了从 <font color="info">%s</font> 到 <font color="info">%s</font> 的合并请求。`

	merge_merge_tmp = `**%s**  通知:
<font color="info">%s</font> 通过了从 <font color="info">%s</font> 到 <font color="info">%s</font> 的合并请求。`
)

func MakePushMsg(project, name, branch string) string {
	str := fmt.Sprintf(push_tmp, project, name, branch)
	return str
}

func MakeMergeMsg(project, action, name, s_branch, t_branch string) string {
	ret := ""

	switch action {
	case "open":
		ret = fmt.Sprintf(merge_open_tmp, project, name, s_branch, t_branch)
	case "close":
		ret = fmt.Sprintf(merge_close_tmp, project, name, s_branch, t_branch)
	case "merge":
		ret = fmt.Sprintf(merge_merge_tmp, project, name, s_branch, t_branch)
	default:
	}
	return ret
}

type Message struct {
	MsgType  string            `json:"msgtype"`
	Text     map[string]string `json:"text"`
	Markdown map[string]string `json:"markdown"`
}

type Robot struct {
	tt  string
	url string
}

func (this *Robot) SendToClient(cmd string, args ...string) {

	msg := ""
	switch cmd {
	case event.PushEvent:
		msg = MakePushMsg(args[0], args[1], args[2])
	case event.MergeRequest:
		msg = MakeMergeMsg(args[0], args[1], args[2], args[3], args[4])
	default:
		fmt.Println(this.tt, "no cmd", cmd)
		return
	}

	req := Message{MsgType: "markdown",
		Markdown: map[string]string{"content": msg}}

	resp, err := util.PostJson(this.url, req, 0)
	if err != nil {
		fmt.Printf("sendToClient url:%s err:%s\n", this.url, err)
		return
	}
	if resp.StatusCode != 200 {
		fmt.Printf("sendToClient url:%s code:%d\n", this.url, resp.StatusCode)
		return
	}
	fmt.Printf("sendToClient url:%s ok\n", this.url)

}

type WeixinMaker struct {
	name string
}

func (this *WeixinMaker) Type() string {
	return this.name
}

func (this *WeixinMaker) Make(url string) robot.RobotI {
	return &Robot{
		tt:  this.name,
		url: url,
	}
}

func init() {
	m := &WeixinMaker{
		name: "workweixin",
	}
	robot.RegisterMaker(m)
}

package workweixin

import (
	"bytes"
	"fmt"
	"github.com/yddeng/dutil/dhttp"
	"github.com/yddeng/webhook/common"
	"github.com/yddeng/webhook/proxy/robot"
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
<font color="info">%s</font> 推送提交到 <font color="info">%s</font> 分支。`

	merge_open_tmp = `**%s**  通知:
<font color="info">%s</font> 创建从 <font color="info">%s</font> 到 <font color="info">%s</font> 的合并请求。`

	merge_close_tmp = `**%s**  通知:
<font color="info">%s</font> 关闭了从 <font color="info">%s</font> 到 <font color="info">%s</font> 的合并请求。`

	merge_merge_tmp = `**%s**  通知:
<font color="info">%s</font> 通过了从 <font color="info">%s</font> 到 <font color="info">%s</font> 的合并请求。`

	message_tmp = `**Message** : %s`
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

func MakeMessage(args ...string) string {
	buffer := bytes.Buffer{}
	for _, str := range args {
		buffer.WriteString(str)
	}
	return buffer.String()
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

func (this *Robot) SendToClient(cmd string, args ...string) error {

	msg := ""
	switch cmd {
	case common.Message:
		msg = MakeMessage(args...)
	case common.PushEvent:
		msg = MakePushMsg(args[0], args[1], args[2])
	case common.MergeRequest:
		msg = MakeMergeMsg(args[0], args[1], args[2], args[3], args[4])
	default:
		return fmt.Errorf("invailed cmd :%s", cmd)
	}

	req := Message{MsgType: "markdown",
		Markdown: map[string]string{"content": msg}}

	resp, err := dhttp.PostJson(this.url, req, 0)
	if err != nil {
		return err
	}
	if resp.StatusCode != 200 {
		return fmt.Errorf("statuscode %d", resp.StatusCode)
	}
	return nil
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

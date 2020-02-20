package message

import "fmt"

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

package message

import "fmt"

func MakePushMsg(project, name, branch string) string {
	str := `项目:%s 推送通知
	事件:Push
	提交者: %s
	分支: %s`

	str = fmt.Sprintf(str, project, name, branch)
	return str
}

func MakeMergeMsg(project, name, s_branch, t_branch string) string {
	str := `项目:%s 合并请求
	事件:MergeRequest
	提交者: %s
	源分支: %s
	目标分支: %s`

	str = fmt.Sprintf(str, project, name, s_branch, t_branch)
	return str
}

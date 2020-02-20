package message

import "fmt"

/*
  push
	项目:xxx
	事件:Push
	提交者: xxx
	分支: xxx
	提交次数:x

  merge
	项目xxx收到一次MergeRequest请求
	xxx请求合并分支xxx到xxx
	时间:
*/

func MakePushMsg(project, name, branch string, count int) string {
	str := `项目:%s
	事件:Push
	提交者: %s
	分支: %s
	提交次数:%d`

	str = fmt.Sprintf(str, project, name, branch, count)
	return str
}

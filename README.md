# webhook

1. git事件通知到群机器人。
2. 项目代码自动拉取编译重启（生产期自动部署，正式上线后建议不要使用）。

## 功能支持
git 支持：
- gitlab

机器人支持：
- 企业微信群机器人

事件支持：
- push
- merge_request


## 功能设计

1. 事件通知

当我们push 代码到线上仓库，线上仓库必然知道这个push操作，就会hook（回调）这里的url。将事件通知给群机器人。

```
机器人相关配置
[[Robot]]
RobotType   = "workweixin"                          机器人类型
Homepage    = "https://github.com/yddeng/webhook"   项目路径
RobotUrl    = "fffff"                               机器人的地址
NotifyCmd   = ["push","merge_request"]              # 通知到机器人的事件
```

git将事件处理后，推送到队列，由机器人消费队列事件。
```
func pcall(i interface{}) {
	e := i.(*Event)

	fmt.Println("pcall", e)
	r, ok := robots[e.Homepage]
	if !ok {
		fmt.Println("no robot", e.Homepage)
		return
	}

	if r.checkCmd(e.Cmd) {
		r.instance.SendToClient(e.Cmd, e.Args...)
	}

}
```

2. 自动化部署 todo

假设，我们有三个环境：

- 线上仓库。如Github、GitLab或Gitee（开源中国）
- 本地仓库。日常开发用的。
- 服务器仓库。一般是自动在测试服务器，或者生产服务器。

这里我们要达到的目的是，当有新的本地 commit push 到线上仓库时，服务器仓库自动pull最线上仓库新的代码。

这里采用代理模式。可能部署了多台服务器。
仓库的事件推送到这里，再由这里通知到各个服务器，服务器上部署的client程序执行本地脚本。

## 配置

webhook配置 [webhook配置](doc/webhook配置.md)

机器人配置 [机器人配置](doc/机器人配置.md)

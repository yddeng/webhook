# webhook

gitlab 与 企业微信机器人

监听事件：
- push
- merge_request

teacher message:
**deng** merged _dev_ to _master_ 。




 
git proxy

一个代理，处理git消息，
聊天机器人注册到代理上，有git消息时，推送到对应的机器人，

client
事件处理客户端，实现自动拉取代码，更新服务器等功能。
轮询？推送？ 等方式实现事件拉取，得到代码更新时，调用脚本程序，实现拉取代码，更新等操作
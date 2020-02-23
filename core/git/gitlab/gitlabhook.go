package gitlab

import (
	"encoding/json"
	"fmt"
	"github.com/yddeng/webhook/core/message"
	"github.com/yddeng/webhook/core/robot/workweixin"
	"github.com/yddeng/webhook/core/verify"
	"io/ioutil"
	"net/http"
	"strings"
)

// gitlab
func Hook(w http.ResponseWriter, r *http.Request) {
	_ = r.ParseForm()

	if r.Method != "POST" {
		fmt.Println("method id err:", r.Method)
		return
	}

	//跨域
	w.Header().Set("Access-Control-Allow-Origin", "*")             //允许访问所有域
	w.Header().Add("Access-Control-Allow-Headers", "Content-Type") //header的类型
	w.Header().Set("content-type", "application/json")             //返回数据格式是json

	//fmt.Println(r.Header, r.RemoteAddr)

	// access验证
	event := r.Header.Get("X-Gitlab-Event")
	token := r.Header.Get("X-Gitlab-Token")
	ip := strings.Split(r.RemoteAddr, ":")[0]
	if event == "" || !verify.VerifyAccess(ip, token) {
		fmt.Println("wrong x-gitlab-event OR x-gitlab-token")
		return
	}

	var data, err = ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Printf("Failed to read request: %s\n", err)
		return
	}

	//var f interface{}
	//_ = json.Unmarshal(data, &f)
	//fmt.Println(f)

	switch event {
	case "Push Hook":
		if verify.VerifyCommand("push") {
			PushEvent(data)
		}
	case "Merge Request Hook":
		if verify.VerifyCommand("merge_request") {
			MergeEvent(data)
		}
	default:
		fmt.Printf("event invaild %s\n", event)
	}

}

type GitlabRepository struct {
	Name     string `json:"name"`
	Homepage string `json:"homepage"`
}

type Commit struct {
	ID        string `json:"id"`
	Message   string `json:"message"`
	Timestamp string `json:"timestamp"`
	URL       string `json:"url"`
	Author    Author `json:"author"`
}

type Author struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type GitlabPush struct {
	ObjectKind   string           `json:"object_kind"`
	Ref          string           `json:"ref"`
	UserUsername string           `json:"user_username"`
	Repository   GitlabRepository `json:"repository"`
}

func PushEvent(data []byte) {
	var hook GitlabPush
	err := json.Unmarshal(data, &hook)
	if err != nil {
		fmt.Printf("Failed to parse request: %s\n", err)
		return
	}

	fmt.Println(hook)

	sp := strings.Split(hook.Ref, "/")
	branch := sp[len(sp)-1]
	msg := message.MakePushMsg(hook.Repository.Name, hook.UserUsername, branch)
	workweixin.SendToClient(hook.Repository.Name, msg)
}

type GitlabMergeRequest struct {
	ObjectKind       string           `json:"object_kind"`
	ObjectAttributes ObjectAttributes `json:"object_attributes"`
	Repository       GitlabRepository `json:"repository"`
	User             User             `json:"user"`
}

type ObjectAttributes struct {
	Action       string `json:"action"`
	SourceBranch string `json:"source_branch"`
	TargetBranch string `json:"target_branch"`
}

type User struct {
	Name     string `json:"name"`
	Username string `json:"username"`
}

func MergeEvent(data []byte) {
	var hook GitlabMergeRequest
	err := json.Unmarshal(data, &hook)
	if err != nil {
		fmt.Printf("Failed to parse request: %s\n", err)
		return
	}

	//action : open close merge

	fmt.Println(hook)

	msg := message.MakeMergeMsg(hook.Repository.Name, hook.ObjectAttributes.Action, hook.User.Username,
		hook.ObjectAttributes.SourceBranch, hook.ObjectAttributes.TargetBranch)
	workweixin.SendToClient(hook.Repository.Name, msg)
}

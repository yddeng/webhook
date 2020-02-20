package gitlab

import (
	"encoding/json"
	"fmt"
	"github.com/yddeng/webhook/access"
	"github.com/yddeng/webhook/message"
	"github.com/yddeng/webhook/robot/weixin"
	"io/ioutil"
	"net/http"
)

//GitlabRepository represents repository information from the webhook
type GitlabRepository struct {
	Name        string `json:"name"`
	URL         string `json:"url"`
	Description string `json:"description"`
	Home        string `json:"home"`
}

//Commit represents commit information from the webhook
type Commit struct {
	ID        string `json:"id"`
	Message   string `json:"message"`
	Timestamp string `json:"timestamp"`
	URL       string `json:"url"`
	Author    Author `json:"author"`
}

//Author represents author information from the webhook
type Author struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

//Webhook represents push information from the webhook
type Gitlab struct {
	ObjectKind        string `json:"object_kind"`
	EventName         string `json:"event_name"`
	Ref               string `json:"ref"`
	UserUsername      string `json:"user_username"`
	UserID            int
	ProjectID         int
	Repository        GitlabRepository `json:"repository"`
	Commits           []Commit         `json:"commits"`
	TotalCommitsCount int              `json:"total_commits_count"`
}

// gitlab
func GitlabHook(w http.ResponseWriter, r *http.Request) {
	_ = r.ParseForm()

	if r.Method != "POST" {
		fmt.Println("method id err:", r.Method)
		return
	}

	//跨域
	w.Header().Set("Access-Control-Allow-Origin", "*")             //允许访问所有域
	w.Header().Add("Access-Control-Allow-Headers", "Content-Type") //header的类型
	w.Header().Set("content-type", "application/json")             //返回数据格式是json

	fmt.Println(r.Header, r.RemoteAddr)

	// access验证
	event := r.Header.Get("X-Gitlab-Event")
	token := r.Header.Get("X-Gitlab-Token")
	if event == "" || !access.VerifyAccess(r.RemoteAddr, token) {
		fmt.Println("wrong x-gitlab-event OR x-gitlab-token")
		return
	}

	// 检测事件

	var hook Gitlab

	//read request body
	var data, err = ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Printf("Failed to read request: %s\n", err)
		return
	}

	//unmarshal request body
	err = json.Unmarshal(data, &hook)
	if err != nil {
		fmt.Printf("Failed to parse request: %s\n", err)
		return
	}
	fmt.Println("------------")
	fmt.Println(hook)

	var f interface{}
	_ = json.Unmarshal(data, &f)

	fmt.Println("------------")
	fmt.Println(f)

	msg := message.MakePushMsg(hook.Repository.Name, hook.UserUsername, hook.Ref, hook.TotalCommitsCount)
	weixin.SendToClient(msg)

}

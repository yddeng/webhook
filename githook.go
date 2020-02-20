package webhook

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

//GitlabRepository represents repository information from the webhook
type GitlabRepository struct {
	Name        string
	URL         string
	Description string
	Home        string
}

//Commit represents commit information from the webhook
type Commit struct {
	ID        string
	Message   string
	Timestamp string
	URL       string
	Author    Author
}

//Author represents author information from the webhook
type Author struct {
	Name  string
	Email string
}

//Webhook represents push information from the webhook
type Webhook struct {
	Before            string
	After             string
	Ref               string
	Username          string
	UserID            int
	ProjectID         int
	Repository        GitlabRepository
	Commits           []Commit
	TotalCommitsCount int
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

	fmt.Println(r.Header)

	// access验证
	event := r.Header.Get("X-Gitlab-Event")
	token := r.Header.Get("X-Gitlab-Token")
	if event == "" || !VerifyAccess("", token) {
		fmt.Println("wrong x-gitlab-event OR x-gitlab-token")
		return
	}

	// 检测事件

	var hook Webhook

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

}

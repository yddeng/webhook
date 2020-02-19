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

func GitHook(w http.ResponseWriter, r *http.Request) {
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

	var hook Webhook

	//read request body
	var data, err = ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Printf("Failed to read request: %s", err)
		return
	}

	//unmarshal request body
	err = json.Unmarshal(data, &hook)
	if err != nil {
		fmt.Printf("Failed to parse request: %s", err)
		return
	}

	fmt.Println(hook)

	var f interface{}
	_ = json.Unmarshal(data, &f)

	fmt.Println("------------")
	fmt.Println(f)

}

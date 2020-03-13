package gitlab

import (
	"encoding/json"
	"fmt"
	"github.com/yddeng/webhook/common"
	"github.com/yddeng/webhook/proxy"
	"io/ioutil"
	"net/http"
	"strings"
)

// Header constants
const (
	XGitlabToken = "X-Gitlab-Token"
	XGitlabEvent = "X-Gitlab-Event"
	GitlabName   = "gitlab"
)

const (
	GitlabPushEvent         = "Push Hook"
	GitlabMergeRequestEvent = "Merge Request Hook"
)

func Hook(w http.ResponseWriter, r *http.Request) {
	_ = r.ParseForm()

	if r.Method != "POST" {
		fmt.Println("method id err:", r.Method)
		return
	}

	// access验证
	event := r.Header.Get(XGitlabEvent)
	token := r.Header.Get(XGitlabToken)
	if event == "" {
		fmt.Println("x-gitlab-event is nil")
		return
	}

	ip := strings.Split(r.RemoteAddr, ":")[0]
	if err := proxy.VerifyAccess(ip, token); err != nil {
		fmt.Println(err)
		return
	}

	var data, err = ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Printf("Failed to read request: %s\n", err)
		return
	}
	//fmt.Println(string(data))

	//var f interface{}
	//_ = json.Unmarshal(data, &f)
	//fmt.Println(f)

	var obj GitlabObj
	err = json.Unmarshal(data, &obj)
	if err != nil {
		fmt.Printf("Failed to parse request: %s\n", err)
		return
	}
	//fmt.Println(obj)

	sp := strings.Split(obj.Ref, "/")
	branch := sp[len(sp)-1]

	e := &proxy.Event{
		Homepage: obj.Repository.Homepage,
		Branch:   branch,
	}

	switch event {
	case GitlabPushEvent:
		e.Cmd = common.PushEvent
		e.Args = []string{obj.Repository.Name, obj.UserUsername, branch}
	case GitlabMergeRequestEvent:
		e.Cmd = common.MergeRequest
		e.Args = []string{obj.Repository.Name, obj.ObjectAttributes.Action, obj.User.Username,
			obj.ObjectAttributes.SourceBranch, obj.ObjectAttributes.TargetBranch}
	default:
		fmt.Printf("event invaild %s\n", event)
		return
	}

	proxy.Push(e)
}

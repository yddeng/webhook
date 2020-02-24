package gitlab

import (
	"encoding/json"
	"fmt"
	"github.com/yddeng/webhook/core/event"
	"github.com/yddeng/webhook/core/robot"
	"github.com/yddeng/webhook/core/verify"
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

type GitlabHook struct {
	name string
}

var gitlabHook = &GitlabHook{
	name: GitlabName,
}

func Hook(w http.ResponseWriter, r *http.Request) {
	_ = r.ParseForm()

	if r.Method != "POST" {
		fmt.Println("method id err:", r.Method)
		return
	}

	// access验证
	event := r.Header.Get(XGitlabEvent)
	token := r.Header.Get(XGitlabToken)
	ip := strings.Split(r.RemoteAddr, ":")[0]
	if event == "" || !verify.VerifyAccess(ip, token) {
		fmt.Println("error x-gitlab-event OR x-gitlab-token")
		return
	}

	var data, err = ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Printf("Failed to read request: %s\n", err)
		return
	}
	fmt.Println(string(data))

	//var f interface{}
	//_ = json.Unmarshal(data, &f)
	//fmt.Println(f)

	var obj GitlabObj
	err = json.Unmarshal(data, &obj)
	if err != nil {
		fmt.Printf("Failed to parse request: %s\n", err)
		return
	}
	fmt.Println(obj)

	switch event {
	case GitlabPushEvent:
		PushEvent(obj)
	case GitlabMergeRequestEvent:
		MergeEvent(obj)
	default:
		fmt.Printf("event invaild %s\n", event)
	}

}

func PushEvent(obj GitlabObj) {

	sp := strings.Split(obj.Ref, "/")
	branch := sp[len(sp)-1]

	args := []string{obj.Repository.Name, obj.UserUsername, branch}

	robot.PushEvent(&robot.Event{
		Homepage: obj.Repository.Homepage,
		Cmd:      event.PushEvent,
		Args:     args,
	})
}

func MergeEvent(obj GitlabObj) {
	args := []string{obj.Repository.Name, obj.ObjectAttributes.Action, obj.User.Username,
		obj.ObjectAttributes.SourceBranch, obj.ObjectAttributes.TargetBranch}

	robot.PushEvent(&robot.Event{
		Homepage: obj.Repository.Homepage,
		Cmd:      event.MergeRequest,
		Args:     args,
	})
}

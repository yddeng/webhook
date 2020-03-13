package main

import (
	"bytes"
	"fmt"
	"net/http"
)

var pushStr = `{"object_kind":"push","event_name":"push","before":"57a2c34539d39a903ade1efb03aa159e24e2416c","after":"180c7e65d7253a67103d3b039a6da20d839057ee","ref":"refs/heads/master","checkout_sha":"180c7e65d7253a67103d3b039a6da20d839057ee","message":null,"user_id":9,"user_name":"yidongdeng","user_username":"deng","user_email":"248244142@qq.com","user_avatar":"https://secure.gravatar.com/avatar/d956cbbf319398c55de2bb46770efa40?s=80\u0026d=identicon","project_id":60,"project":{"id":60,"name":"documents","description":"","web_url":"https://tech.feiyuapi.com/deng/documents","avatar_url":null,"git_ssh_url":"ssh://git@tech.feiyuapi.com:1022/deng/documents.git","git_http_url":"https://tech.feiyuapi.com/deng/documents.git","namespace":"deng","visibility_level":0,"path_with_namespace":"deng/documents","default_branch":"master","ci_config_path":null,"homepage":"https://tech.feiyuapi.com/deng/documents","url":"ssh://git@tech.feiyuapi.com:1022/deng/documents.git","ssh_url":"ssh://git@tech.feiyuapi.com:1022/deng/documents.git","http_url":"https://tech.feiyuapi.com/deng/documents.git"},"commits":[{"id":"180c7e65d7253a67103d3b039a6da20d839057ee","message":"Merge branch 'dev' into 'master'\n\nnone\n\nSee merge request deng/documents!9","timestamp":"2020-02-20T08:13:46Z","url":"https://tech.feiyuapi.com/deng/documents/commit/180c7e65d7253a67103d3b039a6da20d839057ee","author":{"name":"yidongdeng","email":"248244142@qq.com"},"added":[],"modified":["README.md"],"removed":[]},{"id":"e2c2f17b1ce327396ff8d77b61afba3d689cdb6f","message":"none\n","timestamp":"2020-02-20T07:44:00Z","url":"https://tech.feiyuapi.com/deng/documents/commit/e2c2f17b1ce327396ff8d77b61afba3d689cdb6f","author":{"name":"deng","email":"248244142@qq.com"},"added":[],"modified":["README.md"],"removed":[]},{"id":"57a2c34539d39a903ade1efb03aa159e24e2416c","message":"Merge branch 'dev' into 'master'\n\ndev\n\nSee merge request deng/documents!6","timestamp":"2020-02-20T07:06:50Z","url":"https://tech.feiyuapi.com/deng/documents/commit/57a2c34539d39a903ade1efb03aa159e24e2416c","author":{"name":"yidongdeng","email":"248244142@qq.com"},"added":[],"modified":["README.md"],"removed":[]}],"total_commits_count":3,"repository":{"name":"documents","url":"ssh://git@tech.feiyuapi.com:1022/deng/documents.git","description":"","homepage":"https://tech.feiyuapi.com/deng/documents","git_http_url":"https://tech.feiyuapi.com/deng/documents.git","git_ssh_url":"ssh://git@tech.feiyuapi.com:1022/deng/documents.git","visibility_level":0}}`

func main() {
	url := "http://127.0.0.1:7845/hook/gitlab"
	data := []byte(pushStr)

	client := &http.Client{}
	req, err := http.NewRequest("POST", url, bytes.NewReader(data))
	if err != nil {
		fmt.Println(1, err)
		return
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Gitlab-Event", "Push Hook")
	req.Header.Set("X-Gitlab-Token", "token")

	req.Close = true

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(2, err)
		return
	}

	fmt.Println(resp.StatusCode)
}

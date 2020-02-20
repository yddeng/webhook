# webhook

gitlab 与 企业微信机器人

监听事件：
- push
- merge_request

map[Connection:[close] Content-Length:[1732] Content-Type:[application/json] X-Gitlab-Event:[Push Hook]]
{8fcd09fae82a56f041e191a84e8836a3ee046df2 65ef8dd02cb762f733f9241b6d2e25c071dff668 refs/heads/master  0 0 {documents ssh://git@tech.feiyuapi.com:1022/deng/documents.git  } [{65ef8dd02cb762f733f9241b6d2e25c071dff668 none
 2020-02-19T10:48:08Z https://tech.feiyuapi.com/deng/documents/commit/65ef8dd02cb762f733f9241b6d2e25c071dff668 {deng 248244142@qq.com}}] 0}
------------
map[after:65ef8dd02cb762f733f9241b6d2e25c071dff668 before:8fcd09fae82a56f041e191a84e8836a3ee046df2 checkout_sha:65ef8dd02cb762f733f9241b6d2e25c071dff668 commits:[map[added:[] author:map[email:248244142@qq.com name:deng] id:65ef8dd02cb762f733f9241b6d2e25c071dff668 message:none
 modified:[README.md] removed:[] timestamp:2020-02-19T10:48:08Z url:https://tech.feiyuapi.com/deng/documents/commit/65ef8dd02cb762f733f9241b6d2e25c071dff668]] event_name:push message:<nil> object_kind:push project:map[avatar_url:<nil> ci_config_path:<nil> default_branch:master description: git_http_url:https://tech.feiyuapi.com/deng/documents.git git_ssh_url:ssh://git@tech.feiyuapi.com:1022/deng/documents.git homepage:https://tech.feiyuapi.com/deng/documents http_url:https://tech.feiyuapi.com/deng/documents.git id:60 name:documents namespace:deng path_with_namespace:deng/documents ssh_url:ssh://git@tech.feiyuapi.com:1022/deng/documents.git url:ssh://git@tech.feiyuapi.com:1022/deng/documents.git visibility_level:0 web_url:https://tech.feiyuapi.com/deng/documents] project_id:60 ref:refs/heads/master repository:map[description: git_http_url:https://tech.feiyuapi.com/deng/documents.git git_ssh_url:ssh://git@tech.feiyuapi.com:1022/deng/documents.git homepage:https://tech.feiyuapi.com/deng/documents name:documents url:ssh://git@tech.feiyuapi.com:1022/deng/documents.git visibility_level:0] total_commits_count:1 user_avatar:https://secure.gravatar.com/avatar/d956cbbf319398c55de2bb46770efa40?s=80&d=identicon user_email:248244142@qq.com user_id:9 user_name:yidongdeng user_username:deng]
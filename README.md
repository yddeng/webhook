# webhook

gitlab 与 企业微信机器人

监听事件：
- push
- merge_request

teacher message:
**deng** merged _dev_ to _master_ 。



map[Connection:[close] Content-Length:[2336] Content-Type:[application/json] X-Gitlab-Event:[Push Hook] X-Gitlab-Token:[feiyu2020]]
------------

map[
after:65ef8dd02cb762f733f9241b6d2e25c071dff668 
before:6a4da40fc4afd1822e4f69d38fb9e9af3336958c 
checkout_sha:65ef8dd02cb762f733f9241b6d2e25c071dff668 
commits:[
    map[
        added:[] 
        author:map[email:248244142@qq.com name:deng] 
        id:65ef8dd02cb762f733f9241b6d2e25c071dff668 
        message:none 
        modified:[README.md] 
        removed:[] 
        timestamp:2020-02-19T10:48:08Z 
        url:https://tech.feiyuapi.com/deng/documents/commit/65ef8dd02cb762f733f9241b6d2e25c071dff668
        ] 
    map[added:[] author:map[email:248244142@qq.com name:deng] id:8fcd09fae82a56f041e191a84e8836a3ee046df2 message:none modified:[README.md] removed:[] timestamp:2020-02-19T09:18:37Z url:https://tech.feiyuapi.com/deng/documents/commit/8fcd09fae82a56f041e191a84e8836a3ee046df2] 
    map[added:[] author:map[email:248244142@qq.com name:deng] id:6a4da40fc4afd1822e4f69d38fb9e9af3336958c message:none modified:[README.md] removed:[] timestamp:2020-02-19T09:06:16Z url:https://tech.feiyuapi.com/deng/documents/commit/6a4da40fc4afd1822e4f69d38fb9e9af3336958c]
] 
 event_name:push 
 message:<nil> 
 object_kind:push 
 project:map[
    avatar_url:<nil> 
    ci_config_path:<nil> 
    default_branch:master 
    description: 
    git_http_url:https://tech.feiyuapi.com/deng/documents.git 
    git_ssh_url:ssh://git@tech.feiyuapi.com:1022/deng/documents.git 
    homepage:https://tech.feiyuapi.com/deng/documents 
    http_url:https://tech.feiyuapi.com/deng/documents.git 
    id:60 
    name:documents 
    namespace:deng 
    path_with_namespace:deng/documents ssh_url:ssh://git@tech.feiyuapi.com:1022/deng/documents.git 
    url:ssh://git@tech.feiyuapi.com:1022/deng/documents.git 
    visibility_level:0 
    web_url:https://tech.feiyuapi.com/deng/documents
    ] 
 project_id:60 
 ref:refs/heads/master 
 repository:map[
    description: git_http_url:https://tech.feiyuapi.com/deng/documents.git 
    git_ssh_url:ssh://git@tech.feiyuapi.com:1022/deng/documents.git 
    homepage:https://tech.feiyuapi.com/deng/documents 
    name:documents 
    url:ssh://git@tech.feiyuapi.com:1022/deng/documents.git 
    visibility_level:0
    ] 
 total_commits_count:3 
 user_avatar:https://secure.gravatar.com/avatar/d956cbbf319398c55de2bb46770efa40?s=80&d=identicon 
 user_email:24824142@qq.com 
 user_id:9 
 user_name:xxx 
 user_username:deng]
 
 
 map[
 changes:map[
    author_id:map[current:9 previous:<nil>] 
    created_at:map[current:2020-02-20 12:11:25 +0800 previous:<nil>] 
    description:map[current: previous:<nil>] 
    id:map[current:650 previous:<nil>] 
    iid:map[current:1 previous:<nil>] 
    merge_params:map[
        current:map[force_remove_source_branch:0] 
        previous:map[]
        ] 
    source_branch:map[current:dev previous:<nil>] 
    source_project_id:map[current:60 previous:<nil>] 
    target_branch:map[current:master previous:<nil>] 
    target_project_id:map[current:60 previous:<nil>] 
    title:map[current:dev previous:<nil>] 
    total_time_spent:map[current:0 previous:<nil>] 
    updated_at:map[current:2020-02-20 12:11:25 +0800 previous:<nil>]
    ] 
  event_type:merge_request 
  labels:[] 
  object_attributes:map[
  action:open 
  assignee_id:<nil> 
  author_id:9 
  created_at:2020-02-20 12:11:25 +0800 
  description: 
  head_pipeline_id:<nil> 
  human_time_estimate:<nil> 
  human_total_time_spent:<nil> 
  id:650 iid:1 
  last_commit:map[author:map[email:248244142@qq.com name:deng] 
  id:6828c6bc3f004f6cb613a13967ab49975176f0e9 
  message:dev
  timestamp:2020-02-20T04:09:02Z 
  url:https://tech.feiyuapi.com/deng/documents/commit/6828c6bc3f004f6cb613a13967ab49975176f0e9] 
  last_edited_at:<nil> 
  last_edited_by_id:<nil> 
  merge_commit_sha:<nil> 
  merge_error:<nil> 
  merge_params:map[force_remove_source_branch:0] 
  merge_status:unchecked 
  merge_user_id:<nil> 
  merge_when_pipeline_succeeds:false milestone_id:<nil> 
  source:map[avatar_url:<nil> 
    ci_config_path:<nil> 
    default_branch:master 
    description: git_http_url:https://tech.feiyuapi.com/deng/documents.git 
    git_ssh_url:ssh://git@tech.feiyuapi.com:1022/deng/documents.git 
    homepage:https://tech.feiyuapi.com/deng/documents 
    http_url:https://tech.feiyuapi.com/deng/documents.git 
    id:60 
    name:documents 
    namespace:deng 
    path_with_namespace:deng/documents 
    ssh_url:ssh://git@tech.feiyuapi.com:1022/deng/documents.git 
    url:ssh://git@tech.feiyuapi.com:1022/deng/documents.git 
    visibility_level:0 web_url:https://tech.feiyuapi.com/deng/documents
  ] 
  source_branch:dev 
  source_project_id:60 
  state:opened 
  target:map[
    avatar_url:<nil> 
    ci_config_path:<nil> 
    default_branch:master 
    description: 
    git_http_url:https://tech.feiyuapi.com/deng/documents.git 
    git_ssh_url:ssh://git@tech.feiyuapi.com:1022/deng/documents.git 
    homepage:https://tech.feiyuapi.com/deng/documents 
    http_url:https://tech.feiyuapi.com/deng/documents.git 
    id:60 
    name:documents 
    namespace:deng 
    path_with_namespace:deng/documents ssh_url:ssh://git@tech.feiyuapi.com:1022/deng/documents.git 
    url:ssh://git@tech.feiyuapi.com:1022/deng/documents.git 
    visibility_level:0 
    web_url:https://tech.feiyuapi.com/deng/documents
  ] 
  target_branch:master 
  target_project_id:60 
  time_estimate:0 
  title:dev 
  total_time_spent:0 
  updated_at:2020-02-20 12:11:25 +0800 
  updated_by_id:<nil> 
  url:https://tech.feiyuapi.com/deng/documents/merge_requests/1 
  work_in_progress:false
  ] 
  object_kind:merge_request 
  project:map[avatar_url:<nil> ci_config_path:<nil> 
  default_branch:master 
  description: 
  git_http_url:https://tech.feiyuapi.com/deng/documents.git 
  git_ssh_url:ssh://git@tech.feiyuapi.com:1022/deng/documents.git 
  homepage:https://tech.feiyuapi.com/deng/documents http_url:https://tech.feiyuapi.com/deng/documents.git 
  id:60 
  name:documents 
  namespace:deng 
  path_with_namespace:deng/documents ssh_url:ssh://git@tech.feiyuapi.com:1022/deng/documents.git 
  url:ssh://git@tech.feiyuapi.com:1022/deng/documents.git 
  visibility_level:0 
  web_url:https://tech.feiyuapi.com/deng/documents] 
  repository:map[
  description: 
  homepage:https://tech.feiyuapi.com/deng/documents 
  name:documents 
  url:ssh://git@tech.feiyuapi.com:1022/deng/documents.git] 
  user:map[avatar_url:https://secure.gravatar.com/avatar/d956cbbf319398c55de2bb46770efa40?s=80&d=identicon 
  name:yidongdeng 
  username:deng]]
 
 
 
 
 map[changes:map[author_id:map[current:9 previous:<nil>] created_at:map[current:2020-02-20 14:49:23 +0800 previous:<nil>] description:map[current: previous:<nil>] id:map[current:653 previous:<nil>] iid:map[current:4 previous:<nil>] merge_params:map[current:map[force_remove_source_branch:0] previous:map[]] source_branch:map[current:dev previous:<nil>] source_project_id:map[current:60 previous:<nil>] target_branch:map[current:master previous:<nil>] target_project_id:map[current:60 previous:<nil>] title:map[current:dev previous:<nil>] total_time_spent:map[current:0 previous:<nil>] updated_at:map[current:2020-02-20 14:49:23 +0800 previous:<nil>]] event_type:merge_request labels:[] object_attributes:map[action:open assignee_id:<nil> author_id:9 created_at:2020-02-20 14:49:23 +0800 description: head_pipeline_id:<nil> human_time_estimate:<nil> human_total_time_spent:<nil> id:653 iid:4 last_commit:map[author:map[email:248244142@qq.com name:deng] id:6828c6bc3f004f6cb613a13967ab49975176f0e9 message:dev
  timestamp:2020-02-20T04:09:02Z url:https://tech.feiyuapi.com/deng/documents/commit/6828c6bc3f004f6cb613a13967ab49975176f0e9] last_edited_at:<nil> last_edited_by_id:<nil> merge_commit_sha:<nil> merge_error:<nil> merge_params:map[force_remove_source_branch:0] merge_status:unchecked merge_user_id:<nil> merge_when_pipeline_succeeds:false milestone_id:<nil> source:map[avatar_url:<nil> ci_config_path:<nil> default_branch:master description: git_http_url:https://tech.feiyuapi.com/deng/documents.git git_ssh_url:ssh://git@tech.feiyuapi.com:1022/deng/documents.git homepage:https://tech.feiyuapi.com/deng/documents http_url:https://tech.feiyuapi.com/deng/documents.git id:60 name:documents namespace:deng path_with_namespace:deng/documents ssh_url:ssh://git@tech.feiyuapi.com:1022/deng/documents.git url:ssh://git@tech.feiyuapi.com:1022/deng/documents.git visibility_level:0 web_url:https://tech.feiyuapi.com/deng/documents] source_branch:dev source_project_id:60 state:opened target:map[avatar_url:<nil> ci_config_path:<nil> default_branch:master description: git_http_url:https://tech.feiyuapi.com/deng/documents.git git_ssh_url:ssh://git@tech.feiyuapi.com:1022/deng/documents.git homepage:https://tech.feiyuapi.com/deng/documents http_url:https://tech.feiyuapi.com/deng/documents.git id:60 name:documents namespace:deng path_with_namespace:deng/documents ssh_url:ssh://git@tech.feiyuapi.com:1022/deng/documents.git url:ssh://git@tech.feiyuapi.com:1022/deng/documents.git visibility_level:0 web_url:https://tech.feiyuapi.com/deng/documents] target_branch:master target_project_id:60 time_estimate:0 title:dev total_time_spent:0 updated_at:2020-02-20 14:49:23 +0800 updated_by_id:<nil> url:https://tech.feiyuapi.com/deng/documents/merge_requests/4 work_in_progress:false] object_kind:merge_request project:map[avatar_url:<nil> ci_config_path:<nil> default_branch:master description: git_http_url:https://tech.feiyuapi.com/deng/documents.git git_ssh_url:ssh://git@tech.feiyuapi.com:1022/deng/documents.git homepage:https://tech.feiyuapi.com/deng/documents http_url:https://tech.feiyuapi.com/deng/documents.git id:60 name:documents namespace:deng path_with_namespace:deng/documents ssh_url:ssh://git@tech.feiyuapi.com:1022/deng/documents.git url:ssh://git@tech.feiyuapi.com:1022/deng/documents.git visibility_level:0 web_url:https://tech.feiyuapi.com/deng/documents] repository:map[description: homepage:https://tech.feiyuapi.com/deng/documents name:documents url:ssh://git@tech.feiyuapi.com:1022/deng/documents.git] user:map[avatar_url:https://secure.gravatar.com/avatar/d956cbbf319398c55de2bb46770efa40?s=80&d=identicon name:yidongdeng username:deng]]
 {merge_request {dev master} {documents ssh://git@tech.feiyuapi.com:1022/deng/documents.git} {yidongdeng deng}}
 {text map[content:项目:documents 合并请求
 事件:MergeRequest
 提交者: deng
 源分支: dev
 目标分支: master]}
 map[changes:map[
 state:map[current:closed previous:opened] 
 total_time_spent:map[current:0 previous:<nil>] 
 updated_at:map[current:2020-02-20 14:49:51 +0800 previous:2020-02-20 14:49:23 +0800]
 ] 
 event_type:merge_request 
 labels:[] 
 object_attributes:map[
 action:close assignee_id:<nil> author_id:9 created_at:2020-02-20 14:49:23 +0800 description: head_pipeline_id:<nil> human_time_estimate:<nil> human_total_time_spent:<nil> id:653 iid:4 last_commit:map[author:map[email:248244142@qq.com name:deng] id:6828c6bc3f004f6cb613a13967ab49975176f0e9 message:dev
  timestamp:2020-02-20T04:09:02Z url:https://tech.feiyuapi.com/deng/documents/commit/6828c6bc3f004f6cb613a13967ab49975176f0e9] last_edited_at:<nil> last_edited_by_id:<nil> merge_commit_sha:<nil> merge_error:<nil> merge_params:map[force_remove_source_branch:0] merge_status:can_be_merged merge_user_id:<nil> merge_when_pipeline_succeeds:false milestone_id:<nil> source:map[avatar_url:<nil> ci_config_path:<nil> default_branch:master description: git_http_url:https://tech.feiyuapi.com/deng/documents.git git_ssh_url:ssh://git@tech.feiyuapi.com:1022/deng/documents.git homepage:https://tech.feiyuapi.com/deng/documents http_url:https://tech.feiyuapi.com/deng/documents.git id:60 name:documents namespace:deng path_with_namespace:deng/documents ssh_url:ssh://git@tech.feiyuapi.com:1022/deng/documents.git url:ssh://git@tech.feiyuapi.com:1022/deng/documents.git visibility_level:0 web_url:https://tech.feiyuapi.com/deng/documents] source_branch:dev source_project_id:60 state:closed target:map[avatar_url:<nil> ci_config_path:<nil> default_branch:master description: git_http_url:https://tech.feiyuapi.com/deng/documents.git git_ssh_url:ssh://git@tech.feiyuapi.com:1022/deng/documents.git homepage:https://tech.feiyuapi.com/deng/documents http_url:https://tech.feiyuapi.com/deng/documents.git id:60 name:documents namespace:deng path_with_namespace:deng/documents ssh_url:ssh://git@tech.feiyuapi.com:1022/deng/documents.git url:ssh://git@tech.feiyuapi.com:1022/deng/documents.git visibility_level:0 web_url:https://tech.feiyuapi.com/deng/documents] target_branch:master target_project_id:60 time_estimate:0 title:dev total_time_spent:0 updated_at:2020-02-20 14:49:51 +0800 updated_by_id:<nil> url:https://tech.feiyuapi.com/deng/documents/merge_requests/4 work_in_progress:false] object_kind:merge_request project:map[avatar_url:<nil> ci_config_path:<nil> default_branch:master description: git_http_url:https://tech.feiyuapi.com/deng/documents.git git_ssh_url:ssh://git@tech.feiyuapi.com:1022/deng/documents.git homepage:https://tech.feiyuapi.com/deng/documents http_url:https://tech.feiyuapi.com/deng/documents.git id:60 name:documents namespace:deng path_with_namespace:deng/documents ssh_url:ssh://git@tech.feiyuapi.com:1022/deng/documents.git url:ssh://git@tech.feiyuapi.com:1022/deng/documents.git visibility_level:0 web_url:https://tech.feiyuapi.com/deng/documents] repository:map[description: homepage:https://tech.feiyuapi.com/deng/documents name:documents url:ssh://git@tech.feiyuapi.com:1022/deng/documents.git] user:map[avatar_url:https://secure.gravatar.com/avatar/d956cbbf319398c55de2bb46770efa40?s=80&d=identicon name:yidongdeng username:deng]]
 {merge_request {dev master} {documents ssh://git@tech.feiyuapi.com:1022/deng/documents.git} {yidongdeng deng}}
 {text map[content:项目:documents 合并请求
 事件:MergeRequest
 提交者: deng
 源分支: dev
 目标分支: master]}
 
 
 

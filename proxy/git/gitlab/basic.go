package gitlab

/*

type GitStDiff struct {
	Diff        string `json:"diff"`
	NewPath     string `json:"new_path"`
	OldPath     string `json:"old_path"`
	AMode       string `json:"a_mode"`
	BMode       string `json:"b_mode"`
	NewFile     bool   `json:"new_file"`
	RenamedFile bool   `json:"renamed_file"`
	DeletedFile bool   `json:"deleted_file"`
}

type GitIssue struct {
	Id          int       `json:"id"`
	Title       string    `json:"title"`
	AssigneeId  int       `json:"assignee_id"`
	AuthorId    int       `json:"author_id"`
	ProjectId   int       `json:"project_id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	Position    int       `json:"position"`
	BranchName  string    `json:"branch_name"`
	Description string    `json:"description"`
	MilestoneId string    `json:"milestone_id"`
	State       string    `json:"state"`
	Iid         int       `json:"iid"`
}

type GitSnippet struct {
	Id              int       `json:"id"`
	Title           string    `json:"title"`
	Content         string    `json:"content"`
	AuthorId        int       `json:"author_id"`
	ProjectId       int       `json:"project_id"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
	FileName        string    `json:"file_name"`
	ExpiresAt       time.Time `json:"expires_at"`
	Type            string    `json:"type"`
	VisibilityLevel int       `json:"visibility_level"`
}
*/

type GitlabObj struct {
	ObjectKind       string           `json:"object_kind"`
	EventName        string           `json:"event_name"`
	Before           string           `json:"before"`
	After            string           `json:"after"`
	Ref              string           `json:"ref"`
	CheckoutSha      string           `json:"checkout_sha"`
	Message          string           `json:"message"`
	UserId           int              `json:"user_id"`
	Username         string           `json:"username"`
	UserUsername     string           `json:"user_username"`
	UserEmail        string           `json:"user_email"`
	ProjectId        int              `json:"project_id"`
	Repository       Repository       `json:"repository"`
	Commits          []Commit         `json:"commits"`
	TotalCommitCount int              `json:"total_commits_count"`
	ObjectAttributes ObjectAttributes `json:"object_attributes"`
	User             User             `json:"user"`
}

type Repository struct {
	Name            string `json:"name"`
	Homepage        string `json:"homepage"`
	Url             string `json:"url"`
	Description     string `json:"description"`
	GitHttpUrl      string `json:"git_http_url"`
	GitSSHUrl       string `json:"git_ssh_url"`
	VisibilityLevel int    `json:"visibility_level"`
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

type ObjectAttributes struct {
	Action      string `json:"action"`
	AssigneeId  int    `json:"assignee_id"`
	AuthorId    int    `json:"author_id"`
	CreatedAt   string `json:"created_at"`
	Description string `json:"description"`

	Id         int    `json:"id"`
	Iid        int    `json:"iid"`
	LastCommit Commit `json:"last_commit"`

	MergeStatus string `json:"merge_status"`
	MergeUserId string `json:"merge_user_id"`

	Source          GitRequestTarget `json:"source"`
	SourceBranch    string           `json:"source_branch"`
	SourceProjectId int              `json:"source_project_id"`
	State           string           `json:"state"`
	Target          GitRequestTarget `json:"target"`
	TargetBranch    string           `json:"target_branch"`
	TargetProjectId int              `json:"target_project_id"`

	Title       string `json:"title"`
	UpdatedAt   string `json:"updated_at"`
	MilestoneId int    `json:"milestone_id"`
	Position    int    `json:"position"`
}

type GitRequestTarget struct {
	Name            string `json:"name"`
	Namespace       string `json:"namespace"`
	SSHUrl          string `json:"ssh_url"`
	HttpUrl         string `json:"http_url"`
	VisibilityLevel int    `json:"visibility_level"`
}

type User struct {
	Name      string `json:"name"`
	Username  string `json:"username"`
	AvatarUrl string `json:"avatar_url"`
}

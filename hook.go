package webhook

import (
	"github.com/yddeng/webhook/core/robot"
	"net/http"
)

type GitHook interface {
	Type() string
	Hook(w http.ResponseWriter, r *http.Request)
}

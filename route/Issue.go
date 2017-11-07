package route

import (
	"net/http"

	log "github.com/sirupsen/logrus"
	"github.com/train-cat/starter-issue-subscriber/helper"
	"github.com/train-cat/starter-issue-subscriber/model"
)

// Issue record an issue to the database
func Issue(w http.ResponseWriter, r *http.Request) {
	i, err := model.GetIssueFromHTTPRequest(r)

	if helper.HTTPError(w, err) {
		return
	}

	// process
	log.Infof("%+v\n", i)

	w.WriteHeader(http.StatusNoContent)
}

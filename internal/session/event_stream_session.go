package session

import (
	"errors"
	"net/http"

	"github.com/Rayvid/go_firstattempt/internal/session/get"
	"github.com/Rayvid/go_firstattempt/internal/session/post"
	"github.com/julienschmidt/httprouter"
)

// HandleSession handles event stream sessions
func HandleSession(w http.ResponseWriter, r *http.Request, p httprouter.Params) (err error) {
	topic := p.ByName("topic")

	switch r.Method {
	case "POST":
		err = post.HandleRequest(w, r, topic)
	case "GET":
		err = get.HandleRequest(w, r, topic)
	default:
		err = errors.New("Not supported request type")
	}

	return err
}

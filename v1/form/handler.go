package form

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// Service holds the service configuration
// necessary for endpoints to respond to requests
type Service struct {
	Router *mux.Router
}

// NewFormService returns an HTTP router multiplexer with
// attached handler functions
func NewFormService() *Service {
	svc := &Service{
		Router: mux.NewRouter(),
	}

	svc.Router.Methods(http.MethodPost).Path("/submit").HandlerFunc(formHandler)

	return svc
}

func formHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	name := r.Form["name"]
	email := r.Form["email"]
	message := r.Form["message"]

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("name: %s, email: %s, message: %s", name, email, message)))
	return
}

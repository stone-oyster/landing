package form

import (
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

	svc.Router.Methods(http.MethodGet).Path("/submit").HandlerFunc(formHandler)

	return svc
}

func formHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("your details were submitted!"))
	return
}

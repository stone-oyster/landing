package form

import (
	"context"
	"log"
	"net/http"
	"strings"
	"time"

	"cloud.google.com/go/bigquery"
	"github.com/gorilla/mux"
	"github.com/pkg/errors"
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
		log.Println(errors.Wrap(err, "error parsing form"))
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	user := &User{
		Name:       strings.Join(r.Form["name"], ""),
		Email:      strings.Join(r.Form["email"], ""),
		Subscribed: contains(r.Form["subscribe"], "yes"),
		Messages: []Message{
			// TODO: append to this
			{
				Data:      strings.Join(r.Form["message"], ""),
				Timestamp: time.Now(),
			},
		},
	}

	ctx := context.Background()
	client, err := bigquery.NewClient(ctx, "stoneoyster")
	if err != nil {
		log.Println(errors.Wrap(err, "error getting client"))
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	u := client.Dataset("landing_page_messages").Table("people").Uploader()
	err = u.Put(ctx, &bigquery.StructSaver{
		Struct: user,
		Schema: schema,
	})
	if err != nil {
		log.Println(errors.Wrap(err, "error uploading to messages"))
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
	return
}

func contains(slice []string, item string) bool {
	set := make(map[string]struct{}, len(slice))
	for _, s := range slice {
		set[s] = struct{}{}
	}
	_, ok := set[item]
	return ok
}

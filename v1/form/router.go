package form

import (
	"context"
	"log"
	"net/http"

	"cloud.google.com/go/bigquery"
	"github.com/gorilla/mux"
	"github.com/pkg/errors"
)

// Service holds the service configuration
// necessary for endpoints to respond to requests
type Service struct {
	Router          *mux.Router
	BQClient        *bigquery.Client
	SlackWebhookURL string
}

// NewFormService returns an HTTP router multiplexer with
// attached handler functions
func NewFormService() *Service {

	// get big query client
	ctx := context.Background()
	bq, err := bigquery.NewClient(ctx, "stoneoyster")
	if err != nil {
		log.Fatal(errors.Wrap(err, "error getting client"))
	}

	svc := &Service{
		Router:          mux.NewRouter(),
		BQClient:        bq,
		SlackWebhookURL: "", // Secret - do not commit
	}

	svc.Router.Methods(http.MethodPost).Path("/submit").HandlerFunc(svc.formHandler)

	return svc
}

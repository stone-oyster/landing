package main

import (
	"log"
	"fmt"
	"net/http"
	"os"

	"github.com/stone-oyster/landing/v1/form"
)

func main() {
	svc := form.NewFormService()

	if err := http.ListenAndServe(getPort(), svc.Router); err != nil {
		log.Fatal(err)
	}
}

func getPort() string {
	// When running on Google App Engine, the PORT env
	// variable is set by the runtime. If set, we will
	// serve on the port specified there.
	if port := os.Getenv("PORT"); port != "" {
		return fmt.Sprintf(":%s", port)
	}
	return ":8080"
}

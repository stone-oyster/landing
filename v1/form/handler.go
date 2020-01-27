package form

import (
	"context"
	"log"
	"net/http"

	"cloud.google.com/go/bigquery"
	"github.com/pkg/errors"
)

func (svc *Service) formHandler(w http.ResponseWriter, r *http.Request) {
	user, err := userFromRequest(r)
	if err != nil {
		log.Println(errors.Wrap(err, "error building user object"))
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	u := svc.BQClient.Dataset("landing_page_messages").Table("people").Uploader()
	if err = u.Put(context.Background(), &bigquery.StructSaver{
		Struct: user,
		Schema: schema,
	}); err != nil {
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

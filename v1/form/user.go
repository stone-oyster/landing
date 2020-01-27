package form

import (
	"net/http"
	"strings"
	"time"

	"github.com/pkg/errors"
)

// User is the struct representation of our users
type User struct {
	Email      string
	Name       string
	Subscribed bool
	Messages   []Message
}

func userFromRequest(r *http.Request) (*User, error) {
	if err := r.ParseForm(); err != nil {
		return nil, errors.Wrap(err, "could not parse form")
	}

	email := strings.ToLower(strings.Join(r.Form["email"], ""))
	if email == "" {
		return nil, errors.New("no email in form")
	}

	name := strings.Join(r.Form["name"], "")
	if name == "" {
		return nil, errors.New("no name in form")
	}

	subbed := contains(r.Form["subscribe"], "yes")
	message := strings.Join(r.Form["message"], "")

	return &User{
		Email:      email,
		Name:       name,
		Subscribed: subbed,
		Messages: []Message{
			// TODO: append to this
			{
				Data:      message,
				Timestamp: time.Now(),
			},
		},
	}, nil
}

// Message is the struct representation of our messages
type Message struct {
	Data      string
	Timestamp time.Time
}

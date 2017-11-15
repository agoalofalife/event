package examples

import (
	"fmt"
	"github.com/agoalofalife/event"
	"time"
)

type EmailHistory struct {
	history []Email
}

func (e EmailHistory) Push() {
	for _, email := range e.history {
		email.date = time.Now()
		email.Post()
	}
}

type Email struct {
	sender    string
	recipient string
	subject   string
	text      string
	date      time.Time
}

func (e Email) Post() {
	fmt.Printf(
		`Post email :
		   Sender    : %s,
		   recipient : %s,
		   Subject   : %s,
		   Text      : %s,
		   Date      : %v,
		`, e.sender, e.recipient, e.subject, e.text, e.date)
}
func ExampleStructFunc() {
	e := event.New()

	containerEmail := EmailHistory{
		history: []Email{
			Email{
				sender:    "Jo",
				recipient: "Make",
				subject:   "You forget",
				text:      "You forgot to congratulate me",
			},
		},
	}

	e.Add(containerEmail, containerEmail.Push)

	e.Fire(containerEmail)
}

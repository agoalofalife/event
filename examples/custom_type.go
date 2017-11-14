package examples

import (
	"github.com/agoalofalife/event"
	"fmt"
)

type Profession string

func custom() {
	var engineer Profession = "engineer"

	e := event.New()
	// here is callback
	e.Add(engineer, func(who Profession) {
		fmt.Printf("%s decided to change the profession", who)
	}, engineer)

	e.Fire(engineer)
}

package examples

import (
	"github.com/agoalofalife/event"
	"fmt"
)

func ExampleOneToMany() {
	e := event.New()

	e.Add("receiving.message", func() {
		fmt.Println("Post email")
	})

	e.Add("receiving.message", func() {
		fmt.Println("Post in chat")
	})
	e.Add("receiving.message", func() {
		fmt.Println("Create task")
	})

	e.Go("receiving.message")
	//Post email
	//Post in chat
	//Create task
}

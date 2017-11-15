package examples

import (
	"github.com/agoalofalife/event"
	"fmt"
)

func ExampleBase() {
	e := event.New()
	// here is callback
	var counter int
	callback := func() {
		counter += 1
	}

	e.Add("counter.viewing", callback)

	e.Fire("counter.viewing")
	e.Fire("counter.viewing")
	e.Fire("counter.viewing")
	e.Fire("counter.viewing")

	fmt.Println(counter == 4) // true



	e.Untie(callback)
	e.Fire("counter.viewing")
	e.Fire("counter.viewing")
	e.Fire("counter.viewing")

	fmt.Println(counter == 4) // true because callback delete

	e.Destroy("counter.viewing")
	// if there is no subscriber that panic
	defer func() {
		recover()
		fmt.Println(counter == 4) // true
	}()
	// panic
	e.Fire("counter.viewing")
}
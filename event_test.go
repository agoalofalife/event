package event

import (
	"testing"
	"log"
	"os"
)

func createEvent() *Dispatcher {
	return Constructor()
}

func TestDestroy(t *testing.T) {
	event := createEvent()
	event.Add("test", func() {}, []interface{}{})
	event.Destroy("test")

	if event.existEvent("test") {
		t.Error("event not deleted")
	}
}

func TestUntie(t *testing.T)  {
	event := createEvent()
	closure := func() {}
	event.Add("test", closure, []interface{}{})
	event.Untie(closure)
	log.Println(event.listeners)
	os.Exit(2)
	event.Fire("test")
}
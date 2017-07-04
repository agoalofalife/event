package event

import (
	"testing"
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

	if event.existSubscriber(closure) {
		t.Error("Subscriber not deleted")
	}
}
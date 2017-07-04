package event

import (
	"testing"
	"reflect"
)

const nameEvent  = "test"

func createEvent() *Dispatcher {
	return Constructor()
}

func TestDestroy(t *testing.T) {
	event := createEvent()
	event.Add(nameEvent, func() {}, []interface{}{})
	event.Destroy(nameEvent)

	if event.existEvent(nameEvent) {
		t.Error("event not deleted")
	}
}

func TestUntie(t *testing.T)  {
	event := createEvent()
	closure := func() {}
	event.Add(nameEvent, closure, []interface{}{})
	event.Untie(closure)

	if event.existSubscriber(closure) {
		t.Error("Subscriber not deleted")
	}
}

func TestAdd(t *testing.T)  {
	event := createEvent()
	closure := func() {}
	event.Add(nameEvent, closure, []interface{}{})

	if event.listeners[nameEvent][0][typing] != "func" {
		t.Error("not installed typing")
	}

	if reflect.ValueOf(event.listeners[nameEvent][0][structure]).Pointer() != reflect.ValueOf(closure).Pointer() {
		t.Error("not installed structure")
	}
}

// TODO so not work test
//func TestGo(t *testing.T)  {
//	event := createEvent()
//	nameString := "exist"
//	closure := func(test string) string{
//		return  test
//	}
//	event.Add(nameEvent, closure, []interface{}{nameString})
//
//	event.Go(nameEvent)
//}

func TestGetName(t *testing.T)  {
	type Test struct {}

	if GetName(Test{}) != "Test"{
		t.Error("GetName structure error")
	}

}
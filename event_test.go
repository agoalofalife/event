package event

import (
	"fmt"
	"reflect"
	"testing"
)

const nameEvent = "test"

func createEvent() *Dispatcher {
	return New()
}
func TestDestroy(t *testing.T) {
	t.Parallel()
	event := createEvent()
	event.Add(nameEvent, func() {}, []interface{}{})
	event.Destroy(nameEvent)

	if event.existEvent(nameEvent) {
		t.Error("event not deleted")
	}
}
func TestDestroyNotExist(t *testing.T) {
	t.Parallel()
	defer func() {
		str := recover()

		if str != fmt.Sprintf(eventNotExist, "qwer") {
			t.Fatalf("Wrong panic message: %s", str)
		}
	}()
	event := createEvent()
	event.Destroy(`qwer`)
}
func TestUntie(t *testing.T) {
	t.Parallel()
	event := createEvent()
	closure := func() {}
	event.Add(nameEvent, closure, []interface{}{})
	event.Untie(closure)

	if event.existSubscriber(closure) {
		t.Error("Subscriber not deleted")
	}
}
func TestAdd(t *testing.T) {
	t.Parallel()
	event := createEvent()
	closure := func() {}
	event.Add(nameEvent, closure, []interface{}{})

	if event.listeners[nameEvent][0][typing] != "func" {
		t.Error("not installed typing")
	}

	if reflect.ValueOf(event.listeners[nameEvent][0][perform]).Pointer() != reflect.ValueOf(closure).Pointer() {
		t.Error("not installed structure")
	}
}
func TestFire(t *testing.T) {
	t.Parallel()
	event := createEvent()
	nameString := "exist"
	closure := func(test string) string {
		return test
	}
	event.Add(nameEvent, closure, nameString)

	event.Fire(nameEvent)
}
func TestFireNotExist(t *testing.T) {
	t.Parallel()
	defer func() {
		str := recover()
		if str != fmt.Sprintf(eventNotExist, "qwer") {
			t.Fatalf("Wrong panic message: %s", str)
		}
	}()
	event := createEvent()

	event.Fire(`qwer`)
}
func TestFileNotFoundType(t *testing.T) {
	t.Parallel()
	event := createEvent()

	if event.Fire(2).Error() != notFoundName {
		t.Error("Expected error")
	}
}
func TestGetName(t *testing.T) {
	t.Parallel()
	type Test struct{}

	if GetName(Test{}) != "Test" {
		t.Error("GetName structure error")
	}

}
func TestExistSubscriber(t *testing.T) {
	t.Parallel()
	event := createEvent()
	closure := func() {}
	event.Add(nameEvent, closure, []interface{}{})

	if event.listeners[nameEvent][0][typing] != "func" {
		t.Error("not installed typing")
	}
	if !event.existSubscriber(closure) {
		t.Error("Subscriber test fail")
	}
}
func TestFactoryNames(t *testing.T) {
	t.Parallel()
	// fot test
	type customType int
	var digit customType = 2

	type customTypeStr string
	var str customTypeStr = "test"

	cases := []struct {
		Value        interface{}
		ExpectedErr  string
		ExpectedName string
	}{
		{
			Value: struct {
			}{},
			ExpectedErr:  notFoundName,
			ExpectedName: "",
		},
		{
			Value:        "test",
			ExpectedErr:  "",
			ExpectedName: "test",
		},
		{
			Value: &struct {
			}{},
			ExpectedErr:  notFoundName,
			ExpectedName: "",
		},
		{
			Value:        New(),
			ExpectedErr:  "",
			ExpectedName: "Dispatcher",
		},
		{
			Value:        *New(),
			ExpectedErr:  "",
			ExpectedName: "Dispatcher",
		},
		{
			Value:        digit,
			ExpectedErr:  "",
			ExpectedName: "customType",
		},
		{
			Value:        2,
			ExpectedErr:  notFoundName,
			ExpectedName: "",
		},
		{
			Value:        str,
			ExpectedErr:  "",
			ExpectedName: "customTypeStr",
		},
	}

	for _, cas := range cases {
		name, err := factoryNames(cas.Value)
		if name != cas.ExpectedName && err.Error() != cas.ExpectedErr {
			fmt.Println(name, err)
			t.Errorf("Expected value name : %s and error : %s", cas.ExpectedName, cas.ExpectedErr)
		}
	}
}

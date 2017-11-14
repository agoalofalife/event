package event

import (
	"fmt"
	"github.com/pkg/errors"
	"reflect"
)

const (
	typing    = "type"
	perform   = "perform"
	arguments = "arguments"
)

// Dispatcher Base struct store listeners
// list listeners
// display map events :
//     event [name-events] --
//			     -- [number-iterate]
//			         -- type => [type-structure]
//				 -- structure => [structure]
//				 -- arguments => ...arguments slice : interface{}
type Dispatcher struct {
	listeners map[string]map[int]map[string]interface{}
}

// New empty create Dispatcher
func New() *Dispatcher {
	d := &Dispatcher{}
	d.listeners = map[string]map[int]map[string]interface{}{}
	return d
}

// Add new listeners
func (dispatcher *Dispatcher) Add(name interface{}, performing interface{}, parameters ...interface{}) (flag bool, err error) {
	// get name subscriber
	subscriber, err := factoryNames(name)
	if err == nil {
		flag = true
	}

	nameType := getType(performing)

	if _, exist := dispatcher.listeners[subscriber]; !exist {
		dispatcher.listeners[subscriber] = make(map[int]map[string]interface{})
	}

	dispatcher.listeners[subscriber][len(dispatcher.listeners[subscriber])] = map[string]interface{}{
		typing:    nameType.String(),
		perform:   performing,
		arguments: parameters,
	}
	return flag, err
}

// Go alias method Fire
func (dispatcher *Dispatcher) Go(event interface{}) (err error) {
	subscriber, err := factoryNames(event)
	if err != nil {
		return err
	}

	if dispatcher.existEvent(subscriber) {
		for _, iterate := range dispatcher.listeners[subscriber] {
			resolver(iterate[typing].(string), iterate[perform], iterate[arguments].([]interface{}))
		}
	} else {
		panic(fmt.Sprintf(eventNotExist, subscriber))
	}
	return
}

// Fire alias Go method
func (dispatcher *Dispatcher) Fire(event interface{}) (err error) {
	return dispatcher.Go(event)
}

// Destroy or untie event
func (dispatcher *Dispatcher) Destroy(event string) {

	if dispatcher.existEvent(event) {
		delete(dispatcher.listeners, event)
	} else {
		panic(fmt.Sprintf(eventNotExist, event))
	}
}

// Untie Listeners events
func (dispatcher *Dispatcher) Untie(pointer interface{}) {
	for _, event := range dispatcher.listeners {
		for key, iterate := range event {
			if reflect.ValueOf(iterate[perform]).Pointer() == reflect.ValueOf(pointer).Pointer() {
				delete(event, key)
			}
		}
	}
}

// check exist event inner structure
func (dispatcher *Dispatcher) existEvent(event string) bool {
	if _, exist := dispatcher.listeners[event]; exist {
		return true
	}
	return false
}

// check exist subscriber in event
func (dispatcher *Dispatcher) existSubscriber(subscriber interface{}) bool {
	for _, event := range dispatcher.listeners {
		for _, iterate := range event {
			if reflect.ValueOf(iterate[perform]).Pointer() == reflect.ValueOf(subscriber).Pointer() {
				return true
			}
		}
	}
	return false
}

// GetName sugar syntax
func GetName(structure interface{}) string {
	return reflect.TypeOf(structure).Name()
}

// factory name subscribers
func factoryNames(value interface{}) (name string, err error) {
	typeOf := reflect.TypeOf(value)
	switch typeOf.Kind() {
	case reflect.Struct:
		name := typeOf.Name()
		// if struct empty : example struct{}{}
		if name != "" {
			return name, err
		}
		return "", errors.New(notFoundName)
	case reflect.String:
		if typeOf.Name() == typeOf.Kind().String() {
			return value.(string), err
		}
		return typeOf.Name(), err
	case reflect.Ptr:
		name := typeOf.Elem().Name()
		if name != "" {
			return name, err
		}
		return "", errors.New(notFoundName)
	default:
		if typeOf.Name() == typeOf.Kind().String() {
			return "", errors.New(notFoundName)
		}
		return typeOf.Name(), err
	}
}

func resolver(pointerType string, pointer interface{}, parameters []interface{}) {
	switch pointerType {
	// call closure
	case reflect.Func.String():
		in := make([]reflect.Value, len(parameters))

		for key, argument := range parameters {
			in[key] = reflect.ValueOf(argument)
		}
		value := reflect.ValueOf(pointer)
		value.Call(in)
	}
}

// get type  return (func, string ..)
func getType(some interface{}) reflect.Kind {
	return reflect.ValueOf(some).Kind()
}

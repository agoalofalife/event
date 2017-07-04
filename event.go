package event

import (
	"reflect"
)

const (
	typing    = "type"
	structure = "structure"
	arguments = "arguments"
)

type Dispatcher struct {
	// list listeners
	// display map events :
	//     event [name-events] --
	//			     -- [number-iterate]
	//			         -- type => [type-structure]
	//				 -- structure => [structure]
	//				 -- arguments => ...arguments slice : interface{}

	listeners map[string]map[int]map[string]interface{}
}

// start
func Constructor() *Dispatcher {
	d := &Dispatcher{}
	d.listeners = map[string]map[int]map[string]interface{}{}
	return d
}

// add new listeners
func (dispatcher *Dispatcher) Add(name string, performing interface{}, parameters []interface{}) {
	nameType := getType(performing)

	if _, exist := dispatcher.listeners[name]; !exist {
		dispatcher.listeners[name] = map[int]map[string]interface{}{}
	}

	dispatcher.listeners[name][len(dispatcher.listeners[name])] = map[string]interface{}{
		typing:    nameType.String(),
		structure: performing,
		arguments: parameters,
	}
}

func (dispatcher *Dispatcher) Go(event string) {

	if dispatcher.existEvent(event) {
		for _, iterate := range dispatcher.listeners[event] {
			resolver(iterate[typing].(string), iterate[structure], iterate[arguments].([]interface{}))
		}
	} else {
		panic("This is event : '" + event + "'  not exist.")
	}
}

// alias Go method
func (dispatcher *Dispatcher) Fire(event string) {
	dispatcher.Go(event)
}

// remove or untie event
func (dispatcher *Dispatcher) Destroy(event string) {

	if dispatcher.existEvent(event) {
		delete(dispatcher.listeners, event)
	} else {
		panic("This is event : '" + event + "'  not exist.")
	}
}

func (dispatcher *Dispatcher) Untie(pointer interface{}) {
	for _, event := range dispatcher.listeners {
		for key, iterate := range event {
			if reflect.ValueOf(iterate[structure]).Pointer() == reflect.ValueOf(pointer).Pointer() {
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
			if reflect.ValueOf(iterate[structure]).Pointer() == reflect.ValueOf(subscriber).Pointer() {
				return true
			}
		}
	}
	return false
}

// get Name sugar syntax
func GetName(structure interface{}) string {
	return reflect.TypeOf(structure).Name()
}

func resolver(pointerType string, pointer interface{}, parameters []interface{}) {

	switch pointerType {
	// call closure
	case "func":
		in := make([]reflect.Value, 0)

		for _, argument := range parameters {

			in = append(in, reflect.ValueOf(argument))
		}

		value := reflect.ValueOf(pointer)
		value.Call(in)
	}
}

// get type  return (func, string ..)
func getType(some interface{}) reflect.Kind {
	return reflect.ValueOf(some).Kind()
}

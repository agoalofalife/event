package event

import (
	"reflect"
	"log"
	"os"
)

type Event interface {
}

type Listener interface {
	Handler(event Event)
}

type Dispatcher struct {
	// list listeners
	// display map events
	//     event [name-events] --
	//			     -- [number-iterate]
	//			       -- [type-structure] *pointer

	listeners map[string]map[int]map[string]interface{}
}

// start
func Constructor() *Dispatcher {
	d := &Dispatcher{}
	d.listeners = map[string]map[int]map[string]interface{}{}
	return d
}

// add new listeners
func (dispatcher *Dispatcher) AddClosure(name string, performing interface{}) {
	nameType := getType(performing)

	if _, exist := dispatcher.listeners[name]; !exist {
		dispatcher.listeners[name] = map[int]map[string]interface{}{}
	}
	dispatcher.listeners[name][len(dispatcher.listeners[name])] = map[string]interface{}{
		nameType.String(): performing,
	}
}

func (dispatcher *Dispatcher) AddStructure(name interface{}, performing struct{})  {
	nameStructure := getNameStructure(name)
	dispatcher.AddClosure(nameStructure, performing)
}


func (dispatcher *Dispatcher) Go(event string, parameters ...interface{}) {
	if _, exist := dispatcher.listeners[event]; exist {
		for _, iterate := range dispatcher.listeners[event] {
			for typing, pointer := range iterate {
				resolver(typing, pointer, parameters...)
			}
		}
	} else {
		panic("This is event : '" + event + "'  not exist.")
	}
}

// alias Go method
func (dispatcher *Dispatcher) Fire(event string, parameters ...interface{}) {
	dispatcher.Go(event, parameters...)
}

func resolver(pointerType string, pointer interface{}, parameters ...interface{}) {
	log.Println(pointerType)
	os.Exit(2)
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

func getNameStructure(structure interface{}) string {
	return reflect.TypeOf(structure).Name()
}
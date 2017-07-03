package event

import (
	"reflect"
)

type Event interface{
	Add(name string, performing interface{})
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
func Constructor()  *Dispatcher{
	d := &Dispatcher{}
	d.listeners = make(map[string]map[int]map[string]interface{})
	return d
}

// add new event
func (dispatcher *Dispatcher) Add(name string, performing interface{}) {
	nameType := getType(performing)
	//dispatcher.listeners[name] = make(map[int]map[string]interface{})
	//dispatcher.listeners[name][len(dispatcher.listeners[name])] = map[string]interface{}{"w":performing}
	dispatcher.listeners = map[string]map[int]map[string]interface{}{
		name : map[int]map[string]interface{}{
			len(dispatcher.listeners[name]) : map[string]interface{}{
				nameType.String() : performing,
			},
		},
	}

	//dispatcher.listeners[name][len(dispatcher.listeners[name])] = performing
}



// get type  (func(), string ..)
func getType(some interface{}) reflect.Type {
	return reflect.TypeOf(some)
}

package event

import (
	"reflect"
)

type Event interface{

}

// display map events
//     event [name-events] --
//			   -- [type-structure] *link

type Dispatcher struct {
	// list listeners
	listeners map[string]map[string]interface{}
}

func Constructor()  *Dispatcher{
	d := &Dispatcher{}
	d.listeners = make(map[string]map[string]interface{})
	return d
}
// add new event
func (dispatcher *Dispatcher) Add(name string, execute interface{}) {
	nameType := getType(execute)
	dispatcher.listeners[name] = make(map[string]interface{})
	dispatcher.listeners[name][nameType.String()] = execute
}


// get type  (func(), string ..)
func getType(some interface{}) reflect.Type {
	return reflect.TypeOf(some)
}

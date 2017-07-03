package event

import (
	"reflect"
)

type Event interface {
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
func Constructor() *Dispatcher {
	d := &Dispatcher{}
	//d.listeners = make(map[string]map[int]map[string]interface{})
	d.listeners = map[string]map[int]map[string]interface{}{}
	return d
}

// add new listeners
func (dispatcher *Dispatcher) Add(name string, performing interface{}) {
	dispatcher.addListeners(name, performing)
}

func (dispatcher *Dispatcher) addListeners(name string, performing interface{}) {
	nameType := getType(performing)

	//if dispatcher.listeners[name] != nil{
	//       dispatcher.listeners[name][len(dispatcher.listeners[name])][nameType.String()] = performing
	//} else{
	//       dispatcher.listeners = map[string]map[int]map[string]interface{}{
	//	       name : map[int]map[string]interface{}{
	//		       len(dispatcher.listeners[name]) : map[string]interface{}{
	//			       nameType.String() : performing,
	//		       },
	//	       },
	//       }
	//}

	//dispatcher.listeners = map[string]map[int]map[string]interface{}{}

	//dispatcher.listeners[name][len(dispatcher.listeners[name])] = map[string]interface{}{
	//	nameType.String() : performing,
	//}
	//dispatcher.listeners[name][len(dispatcher.listeners[name])] = make(map[string]interface{})

	//dispatcher.listeners[name][len(dispatcher.listeners[name])][ nameType.String()] = performing

	if _, exist := dispatcher.listeners[name]; !exist {
		dispatcher.listeners[name] = map[int]map[string]interface{}{}
	}
	dispatcher.listeners[name][len(dispatcher.listeners[name])] = map[string]interface{}{
		nameType.String(): performing,
	}
}

// get type  (func(), string ..)
func getType(some interface{}) reflect.Type {
	return reflect.TypeOf(some)
}

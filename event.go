package event

import (
	"reflect"
	"fmt"
)

type Event interface{

}

// display map events
//     event [name-events] --
//			   -- [type-structure] *link

type Dispatcher struct {
	// list listeners
	listenersToFunction map[string] func()
}

func (dispatcher *Dispatcher) Add(name string, execute interface{}) {
	fmt.Println(getType(execute))
	//dispatcher.listenersToFunction[name] = closure
}


// get type  (func(), string ..)
func getType(some interface{}) reflect.Type {
	return reflect.TypeOf(some)
}

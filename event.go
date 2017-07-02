package event


type Event interface{

}

type Dispatcher struct {
	// list listeners
	listeners map[string]interface{}
}


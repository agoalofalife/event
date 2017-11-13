

## Event
[![Build Status](https://travis-ci.org/agoalofalife/event.svg?branch=master)](https://travis-ci.org/agoalofalife/event)
[![codecov](https://codecov.io/gh/agoalofalife/event/branch/master/graph/badge.svg)](https://codecov.io/gh/agoalofalife/event)
[![Go Report Card](https://goreportcard.com/badge/github.com/agoalofalife/event)](https://goreportcard.com/report/github.com/agoalofalife/event)
[![GoDoc](http://godoc.org/github.com/agoalofalife/event?status.svg)](http://godoc.org/github.com/agoalofalife/event)

This is package implements [pattern-observer](https://en.wikipedia.org/wiki/Observer_pattern)

### Fast example

```go
import (
	"github.com/agoalofalife/event"
)

func main() {
	// create struct
	e := event.Constructor()

	// subscriber 
	e.Add("push.email", func(text string){
    	// some logic 
    }, []interface{}{"text"})
    
    // init event
    e.Fire("push.email") // or e.Go("push.email")
}
```

let us consider an example:

 * You must first create the structure
 * Next, the first argument declares the name of the event (string type), second argument  executes when the event occurs, the third argument is passed a list of arguments, which are substituted in the parameters of the second argument.
 * In the end you need to run the event. There are two methods available "Go" and his alias "Fire"

### The subscriber function method

```go
func main() {
	// create struct
	e := event.Constructor()
    
    // structure
    type Some struct {}
    // method
    func (s Some) echo(echo string) string {
    	fmt.Println(echo)
    	return echo
    }

	// subscriber 
    temp := Some{}
	e.Add(event.GetName(temp), temp.echo, []interface{}{"second"})
    
    // init event
    e.Fire(event.GetName(temp))
    
    // output in terminal
    .. second
}
```

* In this example we sign the event method structure

 
Read more information at  [WIKI](https://github.com/agoalofalife/event/wiki) :+1:

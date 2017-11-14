package examples

import (
	"fmt"
	"github.com/agoalofalife/event"
	"math/rand"
	"time"
)

type Getter interface {
	Get(msg string)
}

type WareHouse struct{}

func (w WareHouse) Get(msg string) {
	fmt.Printf("there is a new message : '%s' in warehouse\n", msg)
}

type Office struct{}

func (o Office) Get(msg string) {
	fmt.Printf("there is a new message : '%s' in office\n", msg)
}

func eventAndChallels() {
	e := event.New()
	var message = make(chan string)
	eventer := new(Getter)
	wareHouse := new(WareHouse)
	office := new(Office)

	e.Add(eventer, wareHouse.Get)
	e.Add(eventer, office.Get)

	go recipient(e, message, eventer)
	go sender(message)
	//
	time.Sleep(time.Second * 11)
}
func recipient(e *event.Dispatcher, message <-chan string, name *Getter) {
	for {
		select {
		case msg := <-message:
			if msg != "" {
				e.Go(name, msg)
			} else {
				break
			}
		}
	}
}
func sender(message chan<- string) {
	timeLimit := time.Now().Add(time.Second * 10)
	cases := []string{
		"Received a new order",
		"Came shipping",
		"Message from telegram",
	}

	for {
		time.Sleep(time.Second * 2)
		if time.Now().Sub(timeLimit) < 0 {
			message <- cases[rand.Intn(len(cases))]
		} else {
			fmt.Println("Stop")
			time.Sleep(time.Second * 2)
			close(message)
			break
		}
	}
}

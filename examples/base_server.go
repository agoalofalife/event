package examples

import (
	"github.com/agoalofalife/event"
	"net/http"
	"strconv"
)

func ExampleBaseServer() {
	e := event.New()

	type CounterPing int
	var count CounterPing

	e.Add(count, func() { count += 1 })

	http.HandleFunc("/ping", func(writer http.ResponseWriter, request *http.Request) {
		e.Fire(count)
		writer.WriteHeader(http.StatusCreated)
	})

	http.HandleFunc("/count", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte(strconv.Itoa(int(count))))
	})

	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		panic(err)
	}
}

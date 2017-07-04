package event

import (
	"testing"
)

func BenchmarkAdd(b *testing.B){
	e := createEvent()
	for i := 0; i < b.N; i++ {
		e.Add(nameEvent, func() {}, []interface{}{})
	}
}

func BenchmarkGo(b *testing.B) {
	e := createEvent()
	e.Add(nameEvent, func() {}, []interface{}{})
	for i := 0; i < b.N; i++ {
		e.Go(nameEvent)
	}
}
package generic_test

import (
	"testing"

	"github.com/DmitriyVTitov/generic"
)

func TestFanIn(t *testing.T) {
	ch1 := make(chan int)
	ch2 := make(chan int)
	var nilCh chan int // nil channel to test handling of nil channels

	go func() {
		defer close(ch1)
		for i := range 5 {
			ch1 <- i
		}
	}()

	go func() {
		defer close(ch2)
		for i := range 10 {
			ch2 <- i
		}
	}()

	out := generic.FanIn(ch1, ch2, nilCh)

	counter := 0
	for range out {
		counter++
	}

	if counter != 15 {
		t.Errorf("Expected 15 values, but got %d", counter)
	}
}

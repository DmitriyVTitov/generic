package generic

import "sync"

// FanIn takes multiple input channels and returns a single output channel
// that emits values from all input channels. The output channel is closed
// once all input channels are closed.
func FanIn[T any](channels ...<-chan T) <-chan T {
	out := make(chan T)

	var wg sync.WaitGroup

	for _, ch := range channels {
		if ch == nil {
			continue
		}

		wg.Add(1)

		go func(c <-chan T) {
			defer wg.Done()
			for val := range c {
				out <- val
			}
		}(ch)
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}

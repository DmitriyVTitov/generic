package generic

import "sync"

func FanIn[T any](channels ...<-chan T) <-chan T {
	out := make(chan T)

	var wg sync.WaitGroup
	wg.Add(len(channels))

	for _, ch := range channels {
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

func FanOut[T any](in <-chan T, n int) []<-chan T {
	channels := make([]chan T, n)
	for i := range channels {
		channels[i] = make(chan T)
	}

	go func() {
		for val := range in {
			for _, ch := range channels {
				ch <- val
			}
		}

		for _, ch := range channels {
			close(ch)
		}
	}()

	outChannels := make([]<-chan T, n)
	for i, ch := range channels {
		outChannels[i] = ch
	}

	return outChannels
}

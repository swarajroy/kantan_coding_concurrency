package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func repeatFunc[T any, K any](done <-chan K, fn func() T) <-chan T {
	stream := make(chan T)
	go func() {
		defer close(stream)
		for {
			select {
			case <-done:
				return
			default:
				stream <- fn()
			}
		}
	}()
	return stream
}

func isPrime(val int) bool {
	for i := val - 1; i > 1; i-- {
		if val%i == 0 {
			return false
		}
	}
	return true
}

func primesFinder[K any](done <-chan K, stream <-chan int) <-chan int {
	primes := make(chan int)
	go func() {
		defer close(primes)
		for {
			select {
			case <-done:
				return
			case randomInt := <-stream:
				if isPrime(randomInt) {
					primes <- randomInt
				}
			}
		}
	}()
	return primes
}

func take[T any, K any](done <-chan K, stream <-chan T, n int) <-chan T {
	taken := make(chan T)
	go func() {
		defer close(taken)
		for i := 0; i < n; i++ {
			select {
			case <-done:
				return
			case taken <- <-stream:
			}
		}
	}()
	return taken
}

func fanIn[K any](done <-chan K, channels []<-chan int) <-chan int {
	fannedInStream := make(chan int)
	wg := sync.WaitGroup{}

	transfer := func(done <-chan K, c <-chan int) {
		defer wg.Done()
		for i := range c {
			select {
			case <-done:
				return
			case fannedInStream <- i:
			}
		}
	}

	go func() {
		//defer close(fannedInStream)
		for _, c := range channels {
			wg.Add(1)
			go transfer(done, c)
		}
	}()

	go func() {
		wg.Wait()
		close(fannedInStream)
	}()

	return fannedInStream
}

func main1() {
	start := time.Now()
	done := make(chan bool)
	defer close(done)
	randomNumFetcher := func() int { return rand.Intn(500000000) }
	primesStream := primesFinder(done, repeatFunc(done, randomNumFetcher))
	taken := take(done, primesStream, 10)
	for n := range taken {
		fmt.Println(n)
	}

	// fan - out
	// CPUCount := runtime.NumCPU()
	// primeFinderChannels := make([]<-chan int, CPUCount)
	// for i := 0; i < CPUCount; i++ {
	// 	primeFinderChannels[i] = primesFinder(done, repeatFunc(done, randomNumFetcher))
	// }

	// fan - In
	// for val := range take(done, fanIn(done, primeFinderChannels), 10) {
	// 	fmt.Println(val)
	// }

	// fan - in
	fmt.Println("Took = ", time.Since(start))
}

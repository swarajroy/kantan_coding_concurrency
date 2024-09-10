package main

import (
	"fmt"
	"sync"
	"time"
)

func process(data int) int {
	time.Sleep(time.Second * 2)
	val := data * 2
	return val
}

func processData(data int, resultDest *int, wg *sync.WaitGroup) {
	defer wg.Done()

	*resultDest = process(data)
}

func main_2() {
	start := time.Now()
	wg := sync.WaitGroup{}
	input := []int{1, 2, 3, 4, 5}
	result := make([]int, len(input))

	for i, data := range input {
		wg.Add(1)
		go processData(data, &result[i], &wg)
	}
	wg.Wait()

	fmt.Println(result)
	fmt.Println("Took = ", time.Since(start))
}

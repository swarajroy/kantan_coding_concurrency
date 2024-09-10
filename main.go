package main

import "fmt"

func main() {
	// channel1 := make(chan string)
	// channel2 := make(chan string)
	// go func() {
	// 	channel1 <- "channel 1 data"
	// }()

	// go func() {
	// 	channel2 <- "channel 2 data"
	// }()

	// select {
	// case msgFromChannel1 := <-channel1:
	// 	fmt.Printf("msg from channel 1 = %s", msgFromChannel1)
	// case msgFromChannel2 := <-channel2:
	// 	fmt.Printf("msg from channel 2 = %s", msgFromChannel2)
	// }
	// chars := []string{"a", "b", "c"}
	// channel := make(chan string, len(chars))

	// for _, data := range chars {
	// 	channel <- data
	// }

	// close(channel)

	// for data := range channel {
	// 	fmt.Printf("data is %s\n", data)
	// }

	//var forever chan struct{}
	// done := make(chan struct{})
	// defer func() {
	// 	done <- struct{}{}
	// 	//close(done)
	// }()
	// go func(done <-chan struct{}) {
	// 	for {
	// 		select {
	// 		case msg := <-done:
	// 			fmt.Printf("closing the go-routine with msg %s\n", msg)
	// 			return
	// 		default:
	// 			fmt.Println("DOING SOME WORK")
	// 		}
	// 	}
	// }(done)
	// time.Sleep(time.Second * 3)
	//done <- "done"
	//<-forever
	nums := []int{1, 2, 3, 4, 5}

	channel := sliceToChannel(nums)

	resultChannel := sq(channel)

	for data := range resultChannel {
		fmt.Println(data)
	}

}

func sq(in <-chan int) <-chan int {
	out := make(chan int)
	//defer close(out)
	go func() {
		defer close(out)
		for n := range in {
			out <- n * n
		}
	}()

	return out
}

func sliceToChannel(nums []int) <-chan int {
	out := make(chan int)
	//defer close(out)
	go func() {
		defer close(out)
		for _, data := range nums {
			out <- data
		}
	}()
	//close(out)
	//close(out)
	return out
}

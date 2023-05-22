package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Channels in golang")

	mych := make(chan int, 2) // 2 is for number of values you can add to your channel while having single listner
	wg := &sync.WaitGroup{}
	// mych <- 5
	// fmt.Println(<-mych)
	wg.Add(2)

	// By default channels are bidirectional so it can send data or recieve data

	go func(ch chan int, wg *sync.WaitGroup) { // if you want your channel to recieve only in this goroutine then add <-chan int in place of type
		val, ischclosed := <-ch
		fmt.Println(ischclosed)
		fmt.Println(val)
		wg.Done()
	}(mych, wg)

	// For Ex,
	// Recieve only channel
	// go func(ch <-chan int, wg *sync.WaitGroup) {
	// 	val, ischclosed := <-ch
	// 	fmt.Println(ischclosed)
	// 	fmt.Println(val)
	// 	wg.Done()
	// }(mych, wg)

	go func(ch chan int, wg *sync.WaitGroup) { // if you want your channel to send only in this goroutine then add chan<- int in place of type
		close(ch)
		// mych <- 5
		// mych <- 7
		wg.Done()
	}(mych, wg)
	wg.Wait()

	// For, Ex
	// Send only channel
	// go func(ch chan<- int, wg *sync.WaitGroup) {
	// 	close(ch)
	// 	// mych <- 5
	// 	// mych <- 7
	// 	wg.Done()
	// }(mych, wg)
	// wg.Wait()
}

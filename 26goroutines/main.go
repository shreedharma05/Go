package main

import (
	"fmt"
	"net/http"
	"sync"
)

var wg sync.WaitGroup // usually a pointer
var mutex sync.Mutex  // usually a pointer

var signals = []string{"test"} // To add all recieved signals

func main() {
	// go greeter("Um,")
	// greeter("Hello")
	websites := []string{
		"http://lco.dev",
		"http://go.dev",
		"http://google.com",
		"http://roadmap.sh",
	}

	for _, web := range websites {
		go getStatusCode(web)
		wg.Add(1)
	}
	wg.Wait()
	fmt.Println(signals)
}

// func greeter(s string) {
// 	for i := 0; i < 5; i++ {
// 		time.Sleep(1 * time.Second) //pauses the current goroutine for at least the duration
// 		fmt.Println(s)
// 	}
// }

func getStatusCode(website string) {
	defer wg.Done()
	res, err := http.Get(website)
	if err != nil {
		fmt.Println("oops, There is a problem")
	}

	mutex.Lock()
	signals = append(signals, website) // This is a write operation
	// So, while using goroutines there is a high chance of writing multiple values at same time (by mutiple goroutines)
	// To avoid this we use Mutex (mutual exclusion lock) and lock this operation untill it gets finished
	mutex.Unlock()

	fmt.Printf("%v status code for %v\n", res.StatusCode, website)
}

package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Race Conditions in Go")

	wg := &sync.WaitGroup{}
	mut := &sync.Mutex{}

	var score = []int{0}

	wg.Add(3)
	go func(wg *sync.WaitGroup, m *sync.Mutex) {
		fmt.Println("One Routine")
		m.Lock()
		score = append(score, 1)
		m.Unlock()
		wg.Done()
	}(wg, mut)
	go func(wg *sync.WaitGroup, m *sync.Mutex) {
		fmt.Println("Two Routine")
		m.Lock()
		score = append(score, 2)
		m.Unlock()
		wg.Done()
	}(wg, mut)
	go func(wg *sync.WaitGroup, m *sync.Mutex) {
		fmt.Println("Three Routine")
		m.Lock()
		score = append(score, 3)
		m.Unlock()
		wg.Done()
	}(wg, mut)

	wg.Wait()
	fmt.Println(score)
}

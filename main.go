package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	// example1()
	example2()
}

func example2() {

	startTime := time.Now()
	ch := make(chan int)
	wg := sync.WaitGroup{}
	go func() {
		for i := 0; i < 100; i++ {
			wg.Add(1)
			// ch <- i
			go func() {
				defer wg.Done()
				result := job()
				ch <- result
			}()
		}
		wg.Wait()
		close(ch)
	}()

	for v := range ch {
		fmt.Println("Reading channel value : ", v)
	}
	fmt.Println("Taking time : ", time.Since(startTime))
}

func job() int {
	time.Sleep(time.Second)
	return rand.Intn(100)
}

func example1() {
	//# School trip planing
	//Before the School trip you have 3 task
	//1)School Trip Location Search
	//2)Book a Bus
	//3)Trip Activties
	//# Ready for School Trip

	fmt.Println("Planing School Trip...")
	schoolTrip := sync.WaitGroup{}
	schoolTrip.Add(3)
	go func() {
		time.Sleep(1 * time.Second)
		fmt.Println("Rohan searched the location...")
		defer schoolTrip.Done()
	}()
	go func() {
		time.Sleep(2 * time.Second)
		fmt.Println("Rahul Booked the bus...")
		defer schoolTrip.Done()
	}()
	go func() {
		time.Sleep(3 * time.Second)
		fmt.Println("Suraj activity planned...")
		defer schoolTrip.Done()
	}()
	schoolTrip.Wait()
	fmt.Println("Great!!! Ready for the scool trip")
}

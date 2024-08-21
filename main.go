package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
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
		time.Sleep(2 * time.Second)
		fmt.Println("Rohan searched the location...")
		defer schoolTrip.Done()
	}()
	go func() {
		time.Sleep(2 * time.Second)
		fmt.Println("Rahul Booked the bus...")
		defer schoolTrip.Done()
	}()
	go func() {
		time.Sleep(2 * time.Second)
		fmt.Println("Suraj activity planned...")
		defer schoolTrip.Done()
	}()
	schoolTrip.Wait()
	fmt.Println("Great!!! Ready for the scool trip")
}

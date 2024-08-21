package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// Represent a festival task
type Task struct {
	ID   int
	Name string
}

// Represent outcome of task
type Result struct {
	TaskID int
	Output string
}

func job() int {
	time.Sleep(time.Second)
	return rand.Intn(1000)
}

func workers(id int, tasks <-chan Task, results chan<- Result, wg *sync.WaitGroup) {
	// Chan //read-write
	// <-Chan //read-only
	// chan<- //write-only
	defer wg.Done()

	for task := range tasks {
		fmt.Printf("Worker %d started task : %s \n", id, task.Name)
		time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond) // Simulate work
		output := fmt.Sprintf("Worker %d finished '%s'", id, task.Name)
		results <- Result{
			TaskID: task.ID,
			Output: output,
		}
	}
}

func main() {
	// example1()
	// example2()
	// example3()
	example4()
}

func example4() {
	fmt.Println("Organizing School Festival...")

	tasks := []Task{
		{ID: 1, Name: "Decorate the hall"},
		{ID: 2, Name: "Prepare welcome drinks"},
		{ID: 3, Name: "Setup the stage"},
		{ID: 4, Name: "Arrange Chairs"},
		{ID: 5, Name: "Test Sound system"},
		{ID: 6, Name: "Hang Banners"},
		{ID: 7, Name: "Prepare food stalls"},
		{ID: 8, Name: "Setup game boths"},
		{ID: 9, Name: "Coordinate with performers"},
		{ID: 10, Name: "Prepare goodie bags"},
	}
	numWorkers := 3
	taskChan := make(chan Task, len(tasks))
	resultChan := make(chan Result, len(tasks))

	//Create a WaitGroup to wait for all workers to finish task
	var wg sync.WaitGroup

	for i := 1; i <= numWorkers; i++ {
		wg.Add(1)
		go workers(i, taskChan, resultChan, &wg)
	}

	for _, task := range tasks {
		taskChan <- task
	}
	close(taskChan)
	//Start a goroutine to close results channel after all workers finish
	go func() {
		wg.Wait()
		close(resultChan)
	}()
	for v := range resultChan {
		// fmt.Println("Reading channel value : ", v)
		fmt.Printf("Task %d completed : %s \n", v.TaskID, v.Output)
	}

	fmt.Println("All tasks are completed! The School festival is ready to begin!")

}

func example3() {

	startTime := time.Now()
	ch := make(chan int)
	wg := sync.WaitGroup{}
	go func() {
		for i := 0; i < 100000; i++ {
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

func example2() {
	ch := make(chan int)
	go func() {
		for i := 0; i < 100; i++ {
			ch <- job()
		}
		close(ch)
	}()
	for v := range ch {
		fmt.Println("Readin channel value ", v)
	}
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

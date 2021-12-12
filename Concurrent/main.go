package main

import (
	"fmt"
	"sync"
)

/*
	Concurrency: Execute multiple tasks by context switch between them
	Parallelism: Execute multiple tasks simultaneous at the same time
*/
/*
	-GoRoutines: Enable Concurrent Programming
		-have their owen execution stack
        -variable stack space (start from: 2kb)
        -manged by go Runtime(Go Runtime will map goroutines to thread at OS Level )

	-SyncPackage: allow GoRoutines to coordinate their work
        -WaitGroups: coordinate tasks
        -Mutex:Mange Shared Memory(Mutual exclusion lock)
          *only the owner of mutex lock can access the code
	-channels : provide a safe way for go routines to communicate
		-Don't communicate between tasks by sharing memory , share memory by communicating
*/
func main() {
	wg := &sync.WaitGroup{}
	wg.Add(2)
	arr := []int{1, 2, 3}
	//chanel
	ch := make(chan int)
	fmt.Println("Hello from main goroutine")
	go fmt.Println("Hello from another goroutine")

	go f("goroutine", wg)

	go func(msg string, wg *sync.WaitGroup) {
		arr[0] = 6
		for i := 0; i < 3; i++ {
			fmt.Println(msg, ":", i)
		}
		fmt.Println("message Received : ", <-ch)

		wg.Done()
	}("going", wg)

	go func(msg string, wg *sync.WaitGroup) {
		arr[0] = 9
		for i := 0; i < 3; i++ {
			fmt.Println(msg, ":", i)
		}
		ch <- 42
		wg.Done()
	}("Belal", wg)
	//make our main goroutine  waiting , so we other  routines  result
	wg.Wait()
	println(arr[0])
}
func f(from string, wg *sync.WaitGroup) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
	}
	wg.Done()
}

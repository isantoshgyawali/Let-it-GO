package main

import (
    "fmt"
    "sync"
)

func main() {
    // goRoutineHello() //Just running this will cause in program termination before executing printf statements
    WaitGroups()
}

// Functions or methods that run concurrently with other goRoutines
// lightweight in nature and managed by Go Runtime 
// allows concurrent executions but not guarantee the parallel execution
// 
// created using "go" keyword
func goRoutineHello() {
    go fmt.Printf("Hello")
    go fmt.Printf("world")
}
// if we  just executed the function goroutines then, program can close and
// probably will close before even the "printf" statements get executed
// so, 
// we have to wait untill those functions finished processing

// INTRODUCING WAIT_GROUPS
// for every new go_routine that we need to wait 
// do : wg.ADD(1)
// then: defer the wg.Done() so that waitgroup gets closed before the functions end

// ALSO THERE ISN'T ANY SPECIFIC TIMING FOR A FUNCTION CAN FINISH ON ANY TIME_FRAME
// THERE ISN'T ANY ORDER MAINTAINED - BECAUSE THIS IS CONCURRENCY [ out of the normal flow of the program ]
func WaitGroups() {
    var wg sync.WaitGroup

    // method one -- adding all the go functions at once to the WaitGroups
    wg.Add(2)
    go func() {
        defer wg.Done() // Mark this goroutine as done
        fmt.Println("HELLO")
    }()

    go func() {
        defer wg.Done() // Mark this goroutine as done
        fmt.Println("GUYzzzz")
    }()

    // method 2 -- incrementing go functions to WaitGroups rather adding at once
    // more flexible to use, as we can observe
    wg.Add(1)
    go func() {
        defer wg.Done()
        fmt.Println("LOOK WHAT IS THAT!!!!")
    }()

    wg.Add(1)
    go func() {
        defer wg.Done()
        fmt.Println("HOLY MOLY, IT'S a UHMMM....")
    }()

    wg.Wait()
}

func Mutexes() {

}

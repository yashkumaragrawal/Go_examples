package main

import (
    "fmt"
    "time"
    "sync"
)

func worker(id int, job <-chan string, wg *sync.WaitGroup) {
    defer wg.Done()
    for j := range job {
        fmt.Println("worker", id, "started  job", j)
        time.Sleep(time.Millisecond)
    }
}
func main() {
    var wg sync.WaitGroup
    arr := []string{"A","B","C","D","E","F","A1","B1","C1","D1","E1","F1","G","H"}
    n := len(arr) 
    job := make(chan string, n)

    for i := 1; i <= 3; i++ {
	wg.Add(1)
        go worker(i, job, &wg)
    }

    for j := 0; j < n; j++ {
        job <- arr[j]
    }
    close(job)
    wg.Wait()
}

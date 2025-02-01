package main

import (
	"fmt"
	"time"
)

type result struct {
	workerId int
	result   int
}

func main() {
	workers := 3
	data := make(map[int]int)
	resultChannel := make(chan result)
	for i := 0; i < workers; i++ {
		go func() {
			resultChannel <- slowFunc(i)
		}()
	}
	for i := 0; i < workers; i++ {
		result := <-resultChannel
		data[result.workerId] = result.result
	}
	print(len(data))
	fmt.Println("done!")
}

func slowFunc(id int) result {
	println("slow func")
	time.Sleep(time.Second * 1)
	return result{id, id * 4}
}

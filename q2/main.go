package main

import "fmt"

func Finish(i int) {
	fmt.Printf("task %d is finished\n", i)
}

func main() {
	for i := 1; i <= 5; i++ {
		func() {
			fmt.Printf("task %d started\n", i)
			defer Finish(i)
		}()
	}
}

package main

import "fmt"

func main() {
	in := make(chan int)

	go func(out chan int) {
		for i := 0; i <= 4; i++ {
			fmt.Println("Before", i)
			out <- i
			fmt.Println("After", i)
		}
		close(out)
		fmt.Println("Generator finish")
	}(in)

	for i := range in {
		fmt.Println("\tget", i)
	}

	// fmt.Scanln()
}

package main

import "fmt"

func main() {
	for {
		fmt.Println("Loop iteration")
		break
	}

	isRun := true
	for isRun {
		fmt.Println("Loop iteration with condition")
		isRun = false
	}

	for i := 0; i < 2; i++ {
		fmt.Println("Loop iteration", i)
		if i == 1 {
			continue
		}
	}

	sl := []int{1, 2, 3}
	idx := 0

	for idx < len(sl) {
		fmt.Println("while-stype loop edx:", idx, "value:", sl[idx])
		idx++
	}

	for i := 0; i < len(sl); i++ {
		fmt.Println("c-style loop", i, sl[i])
	}
	for idx := range sl {
		fmt.Println("range slice by index", idx)
	}
	for idx, val := range sl {
		fmt.Println("range slice by idx-value", idx, val)
	}

	profile := map[int]string{1: "Hoasker", 2: "RZ"}

	for key := range profile {
		fmt.Println("range map by key", key)
	}

	for key, val := range profile {
		fmt.Println("range map by key-val", key, val)
	}

	for _, val := range profile {
		fmt.Println("range map by val", val)
	}

	str := "Hello, World!"
	for pos, char := range str {
		fmt.Printf("%#U at pos %d\n", char, pos)
	}
}

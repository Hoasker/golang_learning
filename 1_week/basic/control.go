package main

import "fmt"

func main() {
	boolVal := true
	if boolVal {
		fmt.Println("BoolVal is true")
	}

	mapVal := map[string]string{"name": "hoasker"}
	if keyValue, keyExist := mapVal["name"]; keyExist {
		fmt.Println("name =", keyValue)
	}
	if _, keyExist := mapVal["name"]; keyExist {
		fmt.Println("key 'name' exist")
	}

	cond := 1

	if cond == 1 {
		fmt.Println("cond is 1")
	} else if cond == 2 {
		fmt.Println("cond is 2")
	}

	strVal := "name"
	switch strVal {
	case "name":
		fallthrough
	case "test", "lastName":
		// some work
	default:
		//some work
	}

	var val1, val2 = 2, 2
	switch {
	case val1 > 1 || val2 < 11:
		fmt.Println("first block")
	case val2 > 10:
		fmt.Println("second block")
	}

Loop:
	for key, val := range mapVal {
		fmt.Println("switch in loop", key, val)
		switch {
		case key == "lastname":
			break
			fmt.Println("dont pront this")
		case key == "firstName" && val == "Hoasker":
			fmt.Println("switch - break loop here")
			break Loop
		}
	}
}

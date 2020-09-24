package main

import "fmt"

func doNothing() {
	fmt.Println("I'm regular function")
}

func main() {
	func(in string) {
		fmt.Println("Anon func out:", in)
	}("nobody")

	return

	printer := func(in string) {
		fmt.Println("Printer outs:", in)
	}
	printer("As variable")

	return

	type strFuncType func(string)

	worker := func(callback strFuncType) {
		callback("As callback")
	}
	worker(printer)

	return

	prefixer := func(prefix string) strFuncType {
		return func(in string) {
			fmt.Printf("[%s] %s\n", prefix, in)
		}
	}
	successLogger := prefixer("SUCCESS")
	successLogger("Expected behaviour")
}

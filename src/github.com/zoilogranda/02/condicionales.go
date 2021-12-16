package main

import "fmt"

func main() {
	if 5 < 10 && 7 > 1 {
		fmt.Println("Verdadero")
	}
	if 20 > 1 || 2 > 1 {
		fmt.Println("OR")
	}
	if 3 == 3 {
		fmt.Println("==")
	}
	if isValid := true; isValid {
		fmt.Println("con Variable")
		fmt.Printf("%T\n", isValid)
	}
	if false {
		//
	} else {
		fmt.Println("Else")
	}
	fmt.Println("Final")
}

package main

import "fmt"

func main() {
	var a int
	a = 2
	switch a {
	case 1:
		fmt.Println("Vale 1")
	case 2:
		fmt.Println("Vale 2")
	default:
		fmt.Println("Default")
	}
	var b uint8
	b = 9
	switch b {
	case 1:
		fallthrough
	case 2:
		fallthrough
	case 3:
		fallthrough
	case 4:
		fallthrough
	case 5:
		fmt.Println("Es un dia habil")
	case 6:
		fallthrough
	case 7:
		fmt.Println("Es un fin de semana")
	default:
		fmt.Println("No es un dia valido")
	}

	switch c := 3; {
	case c >= 0 && c <= 5:
		fmt.Println("primer case")
	case c >= 6 && c <= 8:
		fmt.Println("segundo case")
	default:
		fmt.Println("default case")
	}
}

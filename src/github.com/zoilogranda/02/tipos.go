package main

import "fmt"

func main() {
	var a int
	var b int8
	a = 1231
	b = 5
	n := "Zoilo"
	n2 := "Granda"

	//casting
	c := a + int(b)
	fmt.Println(c)

	fmt.Printf("hola %s %s %d como estas?\n", n, n2, b)
	fmt.Printf("c es de tipo: %T \n", c)

	var nombre string
	var numero int
	var booleano bool

	fmt.Println(nombre, numero, booleano)
}

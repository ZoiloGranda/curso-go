package main

import "fmt"

func main() {
	var nombres [3]string
	nombres[0] = "Pedro"
	nombres[1] = "Juan"
	nombres[2] = "Maria"

	fmt.Println(nombres)

	apellidos := [3]string{"Garcia", "Chirinos", "Calles"}
	apellido := apellidos[1]

	size := len(nombres)
	fmt.Println(apellidos)
	fmt.Println(apellido)
	fmt.Println(size)
	fmt.Println(nombres[0:2])
}

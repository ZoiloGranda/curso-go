package main

import "fmt"

func main() {
	// var nombres []string
	nombres := make([]string, 0)
	nombres = append(nombres, "Daniel")
	fmt.Printf("tamaño %d capacidad %d\n", len(nombres), cap(nombres))
	nombres = append(nombres, "Pedro")
	fmt.Printf("tamaño %d capacidad %d\n", len(nombres), cap(nombres))
	nombres = append(nombres, "Juan")
	fmt.Printf("tamaño %d capacidad %d\n", len(nombres), cap(nombres))
	nombres = append(nombres, "Maria")
	fmt.Printf("tamaño %d capacidad %d\n", len(nombres), cap(nombres))
	nombres = append(nombres, "Mario")
	fmt.Printf("tamaño %d capacidad %d\n", len(nombres), cap(nombres))
	nombres[2] = "Jose"
	fmt.Println(nombres)

	apellidos := []string{
		"Garcia", "Chirinos", "Correa",
	}
	fmt.Println(apellidos)
}

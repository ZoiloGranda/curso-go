package main

import "fmt"

type Persona struct {
	Nombre string
	Edad   uint8
	Emails []string
}

type Animal struct {
	Raza  string
	Patas uint8
}

type Carro struct {
	Marca  string
	Modelo string
}

func main() {
	var persona1 Persona
	persona1.Nombre = "Pedro"
	persona1.Edad = 20
	fmt.Println(persona1)

	animal1 := Animal{}
	animal1.Raza = "Perro"
	animal1.Patas = 4
	fmt.Println(animal1)

	carro1 := Carro{Marca: "Ford", Modelo: "Explorer"}
	carro2 := Carro{
		Marca:  "Toyota",
		Modelo: "Camry",
	}
	fmt.Println(carro1)
	fmt.Println(carro2.Marca)

	animal2 := Animal{
		"Gato",
		4,
	}
	fmt.Println(animal2)

	persona2 := Persona{
		Nombre: "Juan",
		Edad:   50,
		Emails: []string{"abc@abc.com", "def@def.com"},
	}
	fmt.Println(persona2)
}

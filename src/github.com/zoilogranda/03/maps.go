package main

import "fmt"

func main() {
	idiomas := make(map[string]string)
	idiomas["es"] = "español"
	idiomas["en"] = "ingles"
	idiomas["fr"] = "frances"
	fmt.Println(idiomas)

	nombres := map[string]string{"p": "Pedro", "c": "Carlos", "d": "Daniel"}
	fmt.Println(nombres["p"])

	carros := map[string]string{
		"Toyota":    "Corolla",
		"Ford":      "Explorer",
		"Chevrolet": "Malibu",
	}
	delete(carros, "Ford")
	fmt.Println(carros)
	fmt.Println(carros["Jeep"])

	if carro, ok := carros["Jeep"]; ok {
		fmt.Println("si existe", ok, carro)
	} else {
		fmt.Println("No existe", ok)
	}

	numeros := map[int]string{
		1:  "Es un uno",
		20: "Es un veinte",
		14: "Es un catorce",
		-5: "Es es negativo",
	}

	fmt.Println(numeros)
}

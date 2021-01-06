package main

import "fmt"

func main() {
	planeta1 := "Mercurio"
	planeta2 := "Tierra"

	var segundos float64 = 1000000000

	resultado1 := Edad(segundos, planeta1)
	resultado2 := Edad(segundos, planeta2)

	fmt.Printf("Su edad en el planeta %s es %f años", planeta1, resultado1)
	fmt.Printf("Su edad en el planeta %s es %f años", planeta2, resultado2)
}

const segundosTerrestres = 31557600

func Edad(edadEnSegundos float64, planet string) float64 {
	Planetas := map[string]float64{
		"Tierra":   segundosTerrestres,
		"Mercurio": segundosTerrestres * 0.2408467,
		"Venus":    segundosTerrestres * 0.61519726,
		"Marte":    segundosTerrestres * 1.8808158,
		"Jupiter":  segundosTerrestres * 11.862615,
		"Saturno":  segundosTerrestres * 29.447498,
		"Urano":    segundosTerrestres * 84.016846,
		"Neptuno":  segundosTerrestres * 164.79132,
	}
	return calcularEdad(Planetas[planet], edadEnSegundos)
}

func calcularEdad(edadPlanetaEnSegundos float64, edadEnSegundos float64) float64 {
	return edadEnSegundos / edadPlanetaEnSegundos
}

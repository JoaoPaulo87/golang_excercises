package main

import "fmt"

func main() {

	anio := 2001
	bisiestoResultado := EsBisiesto(anio)

	fmt.Printf("Es el año %d biciesto?\n", anio)

	if bisiestoResultado {

		fmt.Println("Si es biciesto")
		return
	}

	fmt.Println("No es biciesto")
}

func EsBisiesto(anio int) bool {
	switch {
	case anio%400 == 0:
		return true
	case anio%100 == 0:
		return false
	case anio%4 == 0:
		return true
	default:
		return false
	}
}

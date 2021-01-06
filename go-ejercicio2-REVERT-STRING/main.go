package main

import "fmt"

func main() {
	palabra := "Hola!"
	fmt.Println(palabra)

	palabraInvertida := ReverseString(palabra)
	fmt.Println(palabraInvertida)
}

func ReverseString(palabra string) string {
	//Convertimos la palabra en un arreglo de runes. Rune es un tipo de dato, lo que hacemos aca es
	//convertir el string en su equivalente en runes- Rune literals are just 32-bit integer values
	//(however they're untyped constants, so their type can change). They represent unicode codepoints.
	//For example, the rune literal 'a' is actually the number 97.
	runes := []rune(palabra)

	//dentro de un mismo for hacemos dos recorridos y basicamente se cambian de lugar las letras
	for iSubIndex, jSubIndex := 0, len(runes)-1; iSubIndex < jSubIndex; iSubIndex, jSubIndex = iSubIndex+1, jSubIndex-1 {
		runes[iSubIndex], runes[jSubIndex] = runes[jSubIndex], runes[iSubIndex]
	}
	return string(runes)
}

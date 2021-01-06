package main

import (
	"fmt"
	"os"
	"sort"
	"strings"
)

type Equipo struct {
	nombre               string
	partidosJugadosTotal int
	ganadosTotal         int
	empatadosTotal       int
	perdidosTotal        int
	puntosTotal          int
}

type DesenlacePartido struct {
	ganado  string
	perdido string
	empate  string
}

func main() {

	primerPartido := "Allegoric Alaskans;Blithering Badgers;win"
	segundoPartido := "Devastating Donkeys;Courageous Californians;draw"
	terceroPartido := "Devastating Donkeys;Allegoric Alaskans;win"
	cuartoPartido := "Courageous Californians;Blithering Badgers;loss"
	quintoPartido := "Blithering Badgers;Devastating Donkeys;loss"
	sextoPartido := "Allegoric Alaskans;Courageous Californians;win"

	// Guardo los resultados en un slice
	listadoResultados := []string{primerPartido,
		segundoPartido, terceroPartido, cuartoPartido, quintoPartido, sextoPartido,
	}

	// Uso un map para guardar los resultados una vez que se les sacó el ; como separación
	mapEquiposEnlimpio := make(map[string]Equipo)

	// Limpio los resultados y quedan guardados en un slice
	for index := 0; index < len(listadoResultados); index++ {

		resultadoEquiposEnLimpio := LimpiarResultadoPartido(listadoResultados[index])

		equipoA := mapEquiposEnlimpio[resultadoEquiposEnLimpio[0]]
		equipoB := mapEquiposEnlimpio[resultadoEquiposEnLimpio[1]]

		equipoA.nombre = resultadoEquiposEnLimpio[0]
		equipoB.nombre = resultadoEquiposEnLimpio[1]

		equipoA.partidosJugadosTotal++
		equipoB.partidosJugadosTotal++

		switch resultadoEquiposEnLimpio[2] {
		case "win":
			equipoA.ganadosTotal++
			equipoA.puntosTotal += 3
			equipoB.perdidosTotal++
		case "loss":
			equipoA.perdidosTotal++
			equipoB.ganadosTotal++
			equipoB.puntosTotal += 3
		case "draw":
			equipoA.empatadosTotal++
			equipoA.puntosTotal++
			equipoB.empatadosTotal++
			equipoB.puntosTotal++
		default:
			fmt.Println("Resultado erroneo.")
		}

		mapEquiposEnlimpio[resultadoEquiposEnLimpio[0]] = equipoA
		mapEquiposEnlimpio[resultadoEquiposEnLimpio[1]] = equipoB
	}

	// Una vez los resultados fueron limpiados, los guardamos en un slice de Equipo
	ordenarResultados := make([]Equipo, 0, len(mapEquiposEnlimpio))

	for _, equipo := range mapEquiposEnlimpio {
		ordenarResultados = append(ordenarResultados, equipo)
	}

	// En caso de tener el mismo resultado se ordena alfabeticamente y ademas ordenamos los equipos por puntos
	sort.Slice(ordenarResultados, func(i, j int) bool {
		if ordenarResultados[i].puntosTotal == ordenarResultados[j].puntosTotal {
			return ordenarResultados[i].nombre < ordenarResultados[j].nombre
		}
		return ordenarResultados[i].puntosTotal > ordenarResultados[j].puntosTotal
	})

	const imprimirResultados = "%-31v |  %2v |  %2v |  %2v |  %2v |  %2v\n"

	fmt.Fprintf(os.Stdout, imprimirResultados, "Team", "MP", "W", "D", "L", "P")

	// Imprimimos el resultado de la grilla
	for _, equipo := range ordenarResultados {
		fmt.Fprintf(os.Stdout, imprimirResultados, equipo.nombre, equipo.partidosJugadosTotal, equipo.ganadosTotal, equipo.empatadosTotal, equipo.perdidosTotal, equipo.puntosTotal)
	}

}

func LimpiarResultadoPartido(resultado string) []string {

	resultadoEnLimpio := strings.Split(resultado, ";")

	return resultadoEnLimpio
}

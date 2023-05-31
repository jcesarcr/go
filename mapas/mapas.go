package mapas

import "fmt"

func MostrarMapas() {
	paises := make(map[string]string)
	fmt.Println(paises)

	paises["Mexico"] = "DF"
	paises["Argentina"] = "Buenos Aires"
	fmt.Println(paises)
	fmt.Println(paises["Argentina"])

	campeonato := map[string]int{
		"Barcelona":   30,
		"Real Madrid": 38,
		"Chivas":      37,
		"Boca jr":     30,
	}
	fmt.Println(campeonato)
	for equipo, puntaje := range campeonato {
		fmt.Printf("equipo %s, puntahe de %d \n", equipo, puntaje)
	}

	delete(campeonato, "Real Madrid")
	fmt.Println(campeonato)
	puntaje, existe := campeonato["Chivass"]
	fmt.Printf("puntaje es: %d y el equipo existe = %t", puntaje, existe)
}

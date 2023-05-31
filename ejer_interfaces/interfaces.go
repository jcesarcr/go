package ejer_interfaces

import (
	"c/julio/go/interfaces"
	"fmt"
)

func HumanosRespirando(hu interfaces.Humano) {
	hu.Respirar()
	fmt.Printf("Soy un %s y estoy respirando\n", hu.Sexo())

}

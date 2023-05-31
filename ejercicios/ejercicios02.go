package ejercicios

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var numero1 int
var err error
var texto string

func CrearTabla() string {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("INgrese numero 1: ")
	if scanner.Scan() {
		numero1, err = strconv.Atoi(scanner.Text())
		if err != nil {
			panic("El dato ingresado es incorrecto " + err.Error())
		}
	}

	for i := 1; i <= 10; i++ {
		//fmt.Println(numero1 + " x " + i + " = " + numero1*i)
		//fmt.Printf("%v x %v = %v \n", numero1, i, numero1*i)
		texto += fmt.Sprintf("%d x %d = %d \n", numero1, i, numero1*i)
	}

	return texto
}

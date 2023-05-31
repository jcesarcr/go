package variables

import (
	"fmt"
	"strconv"
	"time"
)

var Nombre string
var Estado bool
var Sueldo float32
var Fecha time.Time

func RestoVariables() {
	Nombre = "Julio"
	Estado = true
	Sueldo = 3.14
	Fecha = time.Now().Local()
	fmt.Println(Nombre)
	fmt.Println(Estado)
	fmt.Println(Sueldo)
	fmt.Println(Fecha)
}

func ConviertoaTexto(numero int) (bool, string) {
	var texto string
	texto = strconv.Itoa(numero) //string(numero)
	return true, texto
}

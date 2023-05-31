package main

import (
	"c/julio/go/goroutines"
	"c/julio/go/middleware"
	"fmt"
)

func main() {

	fmt.Println("hello world")
	/*
		fmt.Println("abc")
		variables.MuestroEnteros()
		variables.RestoVariables()
		fmt.Println("======")
		estado, texto := variables.ConviertoaTexto(1111)
		fmt.Println(estado)
		fmt.Println(texto)
		fmt.Println("======")

		if os := runtime.GOOS; os == "Linux." {
			fmt.Println("no es windows, es: ", os)
		} else {
			fmt.Println("Windows: ", os)
		}

		switch os := runtime.GOOS; os {
		case "linux":
			fmt.Println("Esto es linux")
		case "darwin":
			fmt.Println("Esto es darwin")
		default:
			fmt.Printf("%s \n", os)
		}

		numero, texto := ejercicios.ConvNumerico("1345")
		fmt.Println(numero)
		fmt.Println(texto)

		teclado.IngresoNumeros()


		iteraciones.Iterar()

		tabla := ejercicios.CrearTabla()
		fmt.Println(tabla)
	*/
	//files.GrabarTabla()
	//files.SumaTabla()
	//files.LeoArchivo()
	//funciones.Calculos()
	//funciones.LlamarClosure()
	/*funciones.Exponencia(2)
	arreglos_slices.MuestrArreglos()
	arreglos_slices.MuestraSlice()
	arreglos_slices.Capacidad()
	*/
	//mapas.MostrarMapas()
	//users.AltaUsuario()

	/*Pedro := new(modelos.Hombre)
	e.HumanosRespirando(Pedro)

	Juana := new(modelos.Mujer)
	e.HumanosRespirando(Juana)
	*/
	//dp.VemosDefer()
	//dp.EjemploPanic()

	canal1 := make(chan bool)
	go goroutines.MiNombreLentooo("Julio", canal1)
	defer func() {
		<-canal1
	}()
	fmt.Println("Estoy aqui")

	//var x string
	//fmt.Scanln(&x)
	//webserver.MiWebServer()

	middleware.MiMiddleware()
}

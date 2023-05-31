package files

import (
	"bufio"
	"c/julio/go/ejercicios"
	"fmt"
	"os"
)

var filename string = "./files/txt/tabla.txt"

func GrabarTabla() {
	var texto string = ejercicios.CrearTabla()
	archivo, err := os.Create(filename)
	if err != nil {
		fmt.Println("Error al crear archivo ", err)
		return
	}
	fmt.Fprintln(archivo, texto)
	archivo.Close()
}

func SumaTabla() {
	var texto string = ejercicios.CrearTabla()
	if Append(filename, texto) == false {
		fmt.Println("error al concatenar contenido")
	}

}

func Append(filen string, texto string) bool {
	arch, err := os.OpenFile(filename, os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("Error durante el append", err.Error())
		return false
	}
	_, err = arch.WriteString(texto)
	if err != nil {
		fmt.Println("Erorr durante WriteString")
		return false
	}
	arch.Close()
	return true
}

func LeoArchivo() {
	/*archivo, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Error al  leer archivo ", err.Error())
		return
	}
	fmt.Println(string(archivo))*/
	archivo, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error leyendo archivo ", err.Error())
	}
	scanner := bufio.NewScanner(archivo)
	for scanner.Scan() {
		registro := scanner.Text()
		fmt.Println("> ", registro)
	}
	archivo.Close()
}

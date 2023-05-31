package arreglos_slices

import "fmt"

var tablaS []int = []int{2, 3, 4}
var arreglo [10]int = [10]int{6, 44, 33, 22, 11, 34, 55, 66, 77, 88}

func MuestraSlice() {
	fmt.Println(tablaS)
	porcion := arreglo[3:]
	fmt.Println(arreglo)
	fmt.Println(porcion)

	porcion1 := arreglo[:5]
	fmt.Println(porcion1)
	porcion2 := arreglo[6:7]
	fmt.Println(porcion2)

}

func Capacidad() {
	elementos := make([]int, 5, 20)
	fmt.Printf("largo %d, capacidad %d", len(elementos), cap(elementos))

	nums := make([]int, 0, 0)
	for i := 0; i < 100; i++ {
		nums = append(nums, i)
	}
	fmt.Println(nums)
	fmt.Printf("largo %d, capacidad %d", len(nums), cap(nums))
}

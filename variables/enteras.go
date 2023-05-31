package variables

import "fmt"

func MuestroEnteros() {
	var intcomun int
	intde32 := int32(10)
	intde64 := int64(100)

	fmt.Println("variables")
	fmt.Println("intcomun", intcomun)
	fmt.Println("int32 ", intde32)
	fmt.Println("int63 ", intde64)
}

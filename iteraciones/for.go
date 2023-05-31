package iteraciones

import (
	"fmt"
)

func Iterar() {
	i := 0
	for i < 10 {
		fmt.Println(i)
		i++
	}

	for i := 0; i < 100; i += 10 {
		if i == 50 {
			continue
		}
		fmt.Println(i)
	}
}

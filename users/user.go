package users

import (
	"c/julio/go/modelos"
	"fmt"
	"time"
)

func AltaUsuario() {
	u := new(modelos.User)
	u.AddUser(10, "Julio", time.Now(), true)
	fmt.Println(u)
}

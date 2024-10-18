package actions

import (
	"fmt"
)

func Create(tipo string, name string, value float64) {
	fmt.Println("create called", tipo, name, value)
}

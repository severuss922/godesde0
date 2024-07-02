package main

import (
	"fmt"

	"github.com/severuss922/godesde0/variables"
)

func main() {
	estado, texto := variables.ConviertoATexto(123)
	fmt.Println(estado)
	fmt.Println(texto)
}

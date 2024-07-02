package main

import (
	"fmt"
	"runtime"

	"github.com/severuss922/godesde0/ejercicios"
	"github.com/severuss922/godesde0/variables"
)

func main() {
	estado, texto := variables.ConviertoATexto(123)
	fmt.Println(estado)
	fmt.Println(texto)
	if os := runtime.GOOS; os == "linux" || os == "OS X." {
		fmt.Println("Esto no es windows")
	} else {
		fmt.Println("Esto es windows!!")
	}

	switch os := runtime.GOOS; os {
	case "linux":
		fmt.Println("Esto es linux")
	case "darwin":
		fmt.Println("Esto es darwin")
	default:
		fmt.Printf("%s \n", os)
	}

	fmt.Println("--------------------------------------------------------------------------")

	convertido, mensaje := ejercicios.FuncionEjercicio("99l")
	fmt.Println(convertido)
	fmt.Println(mensaje)
}

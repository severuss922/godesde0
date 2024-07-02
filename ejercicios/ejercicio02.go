package ejercicios

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func ValidarError() string {
	scaner := bufio.NewScanner(os.Stdin)
	var numero int
	var err error
	var texto string

	fmt.Println("Ingrese el numero")
	for scaner.Scan() {
		numero, err = strconv.Atoi(scaner.Text())
		if err != nil {
			fmt.Println("Ingrese el numero nuevamente")
		} else {
			break
		}
	}
	fmt.Println("")
	for i := 1; i <= 10; i++ {
		texto += fmt.Sprintf("%d x %d = %d \n", numero, i, i*numero)
	}

	return texto
}

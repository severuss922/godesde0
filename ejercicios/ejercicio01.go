package ejercicios

import (
	"strconv"
)

func FuncionEjercicio(palabra string) (int, string) {
	if convertido, error := strconv.Atoi(palabra); convertido > 100 {
		return convertido, "Es mayor a 100"
	} else if error != nil {
		return 0, error.Error()
	} else {
		return convertido, "Es menor a 100"
	}
}

func ConvNumerico(texto string) (int, string) {
	num, _ := strconv.Atoi(texto)
	if num > 100 {
		return num, "Es mayor a 100"
	} else {
		return num, "Es menor a 100"
	}
}

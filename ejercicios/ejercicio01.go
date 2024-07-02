package ejercicios

import (
	"strconv"
)

func FuncionEjercicio(palabra string) (int, string) {
	if convertido, error := strconv.Atoi(palabra); convertido > 100 {
		return convertido, "Es mayor a 100"
	} else if error != nil && error.Error() != "" {
		return 0, error.Error()
	} else {
		return convertido, "Es menor a 100"
	}
}

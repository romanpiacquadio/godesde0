package ejercicios

import (
	"strconv"
)

func StringANum (texto string) (int, string) {
	num, err := strconv.Atoi(texto)

	if err != nil {
		return 0, "Hubo un error"+err.Error()
	}
	if num >= 100 {
		return num, "Es mayor igual que 100"
	} else {
		return num, "Es menor que 100"
	}
}
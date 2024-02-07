package main

import (
	"fmt"
	// "runtime"
	// "github.com/romanpiacquadio/godesde0/variables"
	"github.com/romanpiacquadio/godesde0/ejercicios"
)

func main() {
	// estado, texto := variables.ConviertoATexto(1588)
	// fmt.Println(estado)
	// fmt.Println(texto)
	
	// if os := runtime.GOOS; os == "Linux." || os == "OS X."{
	// 	fmt.Println("Esto no es windows, es", os)
	// } else {
	// 	fmt.Println("Esto es windows")
	// }

	// switch os := runtime.GOOS; os {
	// case "linux":
	// 	fmt.Println("Esto es Linux")
	// case "darwin":
	// 	fmt.Println("Esto es Darwin")
	// default:
	// 	fmt.Printf("%s \n", os)
	// }
	numero, palabra := ejercicios.StringANum("ffff")
	fmt.Println(numero, palabra)
}

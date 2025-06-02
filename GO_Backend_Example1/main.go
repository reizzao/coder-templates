package main

import (
	"fmt"

	tt_create_example1 "github.com/reizzao/backendgo/entities/example1/usecases/create"
)

func main() {
	fmt.Println("Hello Main!")

	tt_create_example1.TesterCreate()

	nome := "João"
	// msg := "Mostre o valor de %v", nome
	msg := fmt.Sprintf("Mostre o valor de %v", nome)
	fmt.Println(msg)

}

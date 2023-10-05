package main

import (
	"fmt"
	"os"
	"strconv"
	// "reflect"
	// "strconv"
)

func main() {
	var argumentos = os.Args
	var unidadeDestino string
	var valorDestino float64

	if len(argumentos) < 3 {
		fmt.Println("Conversor: <valores> <unidades>")
		os.Exit(1)
	}
	unidadeOrigem := argumentos[len(argumentos)-1]
	valoresOrigem := argumentos[1 : len(argumentos)-1]

	if unidadeOrigem == "celsius" {
		unidadeDestino = "fahrenheit"
	} else if unidadeDestino == "quilometros" {
		unidadeDestino = "milhas"
	} else {
		fmt.Printf("'%s' Não é uma unidade valida\n", unidadeOrigem)
		os.Exit(1)
	}
	for index, value := range valoresOrigem {
		valorOrigem, err := strconv.ParseFloat(value, 64)
		if err != nil {
			fmt.Printf(
				"O valor %s na posição %d não é um numero valido", value, index)
			os.Exit(1)
		}
		if unidadeOrigem == "celsius" {
			valorDestino = valorOrigem*1.8 + 32
		}else{
			valorDestino = valorOrigem / 1.60934
		}
	fmt.Printf("%.2f %s = %.2f %s\n",valorOrigem,unidadeOrigem, valorDestino,unidadeDestino)
	}


}

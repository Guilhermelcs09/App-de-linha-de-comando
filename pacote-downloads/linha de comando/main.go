package main

import (
	"fmt"
	app "linha-de-comando/apl"
	"log"
	"os"
)

func main() {
	fmt.Println("inicio")

	aplicacao := app.Gerar()
	if erro := aplicacao.Run(os.Args); erro != nil {
		log.Fatal(erro)
	}
}

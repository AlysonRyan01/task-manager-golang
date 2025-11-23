package main

import (
	app "app/internal/usecase"
	"log"
	"os"
)

func main() {
	aplicacao := app.Gerar()
	if err := aplicacao.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

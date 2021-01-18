package main

import (
	"log"

	"github.com/Ivovuceticc/microblogging/bd"
	"github.com/Ivovuceticc/microblogging/handlers"
)

func main() {
	if bd.ChequeoConnection() == 0 {
		log.Fatal("Sin conexión a la BD")
		return
	}
	handlers.Manejadores()
}

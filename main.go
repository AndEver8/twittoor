package main

import (
	"log"

	"github.com/AndEver8/twittoor/bd"
	"github.com/AndEver8/twittoor/handlers"
)

func main() {
	if bd.ChequeoConection() == 0 {
		log.Fatal("Sin conexion a la BD")
		return
	}

	handlers.Manejadores()
}

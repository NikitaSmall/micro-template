package main

import (
	"log"

	"github.com/nikitasmall/master-service/router"
)

func main() {
	r := router.CreateRouter()

	log.Println("Master service was started.")
	r.Run(":3000")
}

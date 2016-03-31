package main

import (
	"log"
	"os"

	"github.com/nikitasmall/master-service/config"
	"github.com/nikitasmall/master-service/router"
)

func init() {
	config.InitEnv(".env")
}

func main() {
	r := router.CreateRouter()

	log.Println("Master service was started.")
	r.Run(os.Getenv("PORT"))
}

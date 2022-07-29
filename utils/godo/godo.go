package godo

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func GodoGet(key string) (value string) {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("error cargando el archivo .env")
	}
	value = os.Getenv(key)
	return
}
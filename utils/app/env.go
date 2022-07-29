package app
import (
	"fmt"
	env "github.com/joho/godotenv"
)

func LoadEnv() {
	fmt.Println("loading env")
	err := env.Load()
	if err != nil {
		panic(err)
	}
}

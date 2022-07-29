package main

import (
	"fmt"
	"os"
	"rabbit/interfaces"
	"rabbit/utils/app"
	"rabbit/utils/operations"
)

// ocurrencia de un solo parametro
func main() {

	app.LoadEnv()

	queueConf := interfaces.Queue{
		Name:      os.Getenv("QUEUE_NAME"),
		Durable:   false,
		Delete:    false,
		Exclusive: false,
		NoWait:    false,
	}

	fmt.Println("server on")
	operations.ConsumeQM(queueConf)

}

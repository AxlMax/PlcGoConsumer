package operations

import (
	"encoding/json"
	"fmt"
	"log"
	"rabbit/database/models"
	"rabbit/database/repositories"
	"rabbit/interfaces"
	"rabbit/utils/godo"
	"strings"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}

func ConsumeQM(queueConfig interfaces.Queue) {
	conn, err := amqp.Dial(godo.GodoGet("RABBIT_CONN"))
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		queueConfig.Name,      // name
		queueConfig.Durable,   // durable
		queueConfig.Delete,    // delete when unused
		queueConfig.Exclusive, // exclusive
		queueConfig.NoWait,    // no-wait
		nil,                   // arguments
	)
	failOnError(err, "Failed to declare a queue")

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	failOnError(err, "Failed to register a consumer")

	for d := range msgs {
		var dataIn interfaces.PLCinfo

		json.Unmarshal([]byte(d.Body), &dataIn)

		splitedId := strings.Split(dataIn.Values[0].ID,".")

		var PlcData = models.PlcModel{
			Disp:            fmt.Sprintf("%s.%s",splitedId[0],splitedId[1]),
			PresionCActual:  float32(dataIn.Values[0].V),
			PresionChActual: float32(dataIn.Values[1].V),
			TempActual:      float32(dataIn.Values[2].V),
			Temperatura:     []float32{float32(dataIn.Values[2].V)},
			PresionCh:       []float32{float32(dataIn.Values[1].V)},
			PresionC:        []float32{float32(dataIn.Values[0].V)},
		}

		err, state := repositories.BuscarPlc(PlcData.Disp)
		failOnError(err, "Failed at search plc")

		if state {

			err = repositories.CrearPLC(PlcData)
			failOnError(err, "Failed at created PLC")

			time.Sleep(10 * time.Second)

		} else {
			err = repositories.ActualizarPLC(PlcData)
			failOnError(err, "failed on update PLC")

			time.Sleep(5 * time.Second)
		}
	}
	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
}

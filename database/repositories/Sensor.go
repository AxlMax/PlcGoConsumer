package repositories

import (
	"context"
	"rabbit/database/models"
	"rabbit/utils/app"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const COLLECION_SENSOR = "sensor"

func CrearSensor(body models.Sensor) (sensor models.Sensor, err error) {
	db := app.GetConnection()

	result, err := db.Collection(COLLECION_SENSOR).InsertOne(context.TODO(), body)

	if err != nil {
		return
	}

	id, _ := result.InsertedID.(primitive.ObjectID)
	err = db.Collection(COLLECION_SENSOR).FindOne(context.TODO(), bson.M{"_id": id}).Decode(&sensor)

	return
}

func BuscarSensor(id string) (err error, state bool) {
	db := app.GetConnection()
	state = false

	results, err := db.Collection(COLLECION_SENSOR).Find(
		context.TODO(),
		bson.M{
			"disp": id,
		},
	)

	if err != nil {
		return
	}

	var sensor models.Sensor
	for results.Next(context.TODO()) {
		err = results.Decode(&sensor)
		if err != nil {
			return
		}
	}

	if sensor.ID.Hex() == "000000000000000000000000" {
		state = true
	}
	return
}

func AgregarDato(temperatura float32, corriente float32, nivel float32, presion float32, id string) (err error) {
	db := app.GetConnection()

	_, err = db.Collection(COLLECION_SENSOR).UpdateOne(
		context.TODO(),
		bson.M{"disp": id},
		bson.M{
			"$push": bson.M{
				"temperatura": temperatura,
				"presion":     presion,
				"corriente":   corriente,
				"nivel":       nivel,
			},
			"$set": bson.M{
				"tempActual":  temperatura,
				"presActual":  presion,
				"corrActual":  corriente,
				"nivelActual": nivel,
			},
		},
	)
	return
}

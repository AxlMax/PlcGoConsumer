package repositories

import (
	"context"
	"log"
	"rabbit/database/models"
	"rabbit/utils/app"

	"go.mongodb.org/mongo-driver/bson"
)

const COLLECION_PLC = "plc"

func failOnError(err error, msg string) {
	if err != nil {
		log.Panic(msg, err)
	}
}

func CrearPLC(body models.PlcModel) (err error) {
	db := app.GetConnection()

	_, err = db.Collection(COLLECION_PLC).InsertOne(context.TODO(), body)

	return
}

func BuscarPlc(id string) (err error, state bool) {
	db := app.GetConnection()
	state = false

	results, err := db.Collection(COLLECION_PLC).Find(
		context.TODO(),
		bson.M{
			"disp": id,
		},
	)

	if err != nil {
		return
	}

	var plc models.PlcModel
	for results.Next(context.TODO()) {
		err = results.Decode(&plc)
		if err != nil {
			return
		}
	}

	if plc.ID.Hex() == "000000000000000000000000" {
		state = true
	}
	return
}

func ActualizarPLC(body models.PlcModel) (err error) {
	db := app.GetConnection()

	_, err = db.Collection(COLLECION_PLC).UpdateOne(
		context.TODO(),
		bson.M{"disp": body.Disp},
		bson.M{
			"$push": bson.M{
				"temperatura": body.TempActual,
				"presionC"   : body.PresionCActual,
				"presionCh"  : body.PresionChActual,
			},
			"$set": bson.M{
				"tempActual"  :  body.TempActual,
				"presCActual" :  body.PresionCActual,
				"presChActual":  body.PresionChActual,
			},
		},
	)
	return
}

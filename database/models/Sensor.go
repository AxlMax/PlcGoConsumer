package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Sensor struct {
	ID              primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Disp            string             `json:disp bson:"disp"`
	Temperatura     []float32          `json:"temperatura" bson:"temperatura"`
	TempActual      float32            `json:"tempActual" bson:"tempActual"`
	Presion         []float32          `json:"presion" bson:"presion"`
	PresionActual   float32            `json:"presActual" bson:"presActual"`
	Corriente       []float32          `json:"corriente" bson:"corriente"`
	CorrienteActual float32            `json:"corrActual" bson:"corrActual"`
	Nivel           []float32          `json:"nivel" bson:"nivel"`
	NivelActual     float32            `json:"nivelActual" bson:"nivelActual"`
}

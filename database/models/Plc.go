package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type PlcModel struct {
	ID              primitive.ObjectID `json:"id"           bson:"_id,omitempty"`
	Disp            string             `json:disp           bson:"disp"`
	Temperatura     []float32          `json:"temperatura"  bson:"temperatura"`
	TempActual      float32            `json:"tempActual"   bson:"tempActual"`
	PresionC        []float32          `json:"presionC"     bson:"presionC"`
	PresionCActual  float32            `json:"presCActual"  bson:"presCActual"`
	PresionCh       []float32          `json:"presionCh"    bson:"presionCh"`
	PresionChActual float32            `json:"presChActual" bson:"presChActual"`
}

package models

import "go.mongodb.org/mongo-driver/bson/primitive"

//Create Struct
type Planet struct {
	ID      primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name    string             `json:"name,omitempty" bson:"name,omitempty"`
	Climate string             `json:"climate" bson:"climate,omitempty"`
	Terrain string             `json:"terrain" bson:"terrain,omitempty"`
	Movies  string             `json:"movies" bson:"films,omitempty"`
}

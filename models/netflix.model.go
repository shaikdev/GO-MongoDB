package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Movie struct {
	Id      primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Movie   string             `json:"movie,omitempty"`
	Watched bool               `json:"watched"`
}

func (body *Movie) BodyCheckForMovie() bool {
	if body.Movie == "" {
		return true
	} else {
		return false
	}
}

package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Project struct {
	ID          primitive.ObjectID   `bson:"_id,omitempty"`
	Name        string               `bson:"name"`
	Description string               `bson:"description"`
	Owner       primitive.ObjectID   `bson:"owner"`
	Team        primitive.ObjectID   `bson:"team"`
	Phases      []primitive.ObjectID `bson:"phases"`
	CreatedAt   primitive.DateTime   `bson:"createdAt"`
	UpdatedAt   primitive.DateTime   `bson:"updatedAt"`
}

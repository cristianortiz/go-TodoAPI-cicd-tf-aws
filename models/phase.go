package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Phase struct {
	ID           primitive.ObjectID   `bson:"_id,omitempty"`
	ProjectID    primitive.ObjectID   `bson:"projectId"`
	Name         string               `bson:"name"`
	Description  string               `bson:"description"`
	Tasks        []primitive.ObjectID `bson:"tasks"`
	Dependencies []primitive.ObjectID `bson:"dependencies"`
	Status       string               `bson:"status"`
	CreatedAt    primitive.DateTime   `bson:"createdAt"`
	UpdatedAt    primitive.DateTime   `bson:"updatedAt"`
}

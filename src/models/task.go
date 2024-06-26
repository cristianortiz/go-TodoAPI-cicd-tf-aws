package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Task struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	PhaseID     primitive.ObjectID `bson:"phaseId"`
	Name        string             `bson:"name"`
	Description string             `bson:"description"`
	AssigneedTo primitive.ObjectID `bson:"assignedTo"`
	Status      string             `bson:"status"`
	CreatedAt   primitive.DateTime `bson:"createdAt"`
	UpdatedAt   primitive.DateTime `bson:"updatedAt"`
}

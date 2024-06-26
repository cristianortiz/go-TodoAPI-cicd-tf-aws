package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID              primitive.ObjectID   `bson:"_id,omitempty"`
	Name            string               `bson:"name"`
	Email           string               `bson:"email"`
	PasswordHash    string               `bson:"passwordHash"`
	CreatedProjects []primitive.ObjectID `bsonn:"createdProjects"`
	AssignedTasks   []primitive.ObjectID `bson:"assignedTasks"`
	Teams           []primitive.ObjectID `bson:"teams"`
}

package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Team struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	ProjectID primitive.ObjectID `bson:"projectId"`
	Name      string             `bson:"name"`
	Members   []Member           `bson:"members"`
	CreatedAt primitive.DateTime `bson:"createdAt"`
	UpdatedAt primitive.DateTime `bson:"updatedAt"`
}

type Member struct {
	UserID primitive.ObjectID `bson:"userId"`
	Role   string             `bson:"role"`
}

package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Developer struct {
	ID    primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name  string `json:"name,omitempty" bson:"name,omitempty"`
	Lead  bool   `json:"lead,omitempty" bson:"lead,omitempty"`
	Email string `json:"email,omitempty" bson:"email,omitempty"`
}

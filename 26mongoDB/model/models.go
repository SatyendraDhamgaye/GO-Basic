package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"sort.com/goserver/go/pkg/mod/go.mongodb.org/mongo-driver@v1.13.1/bson/primitive"
)

type Netflix struct {
	ID      primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Movie   string             `json:"movie,omitempty"`
	Watched bool               `json:"watched,omitempty"`
}

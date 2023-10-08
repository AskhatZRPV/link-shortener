package entity

import "go.mongodb.org/mongo-driver/bson/primitive"

type Link struct {
	ID     primitive.ObjectID `bson:"_id,omitempty"`
	URL    string             `bson:"url"`
	Hash   string             `bson:"hash"`
	Domain string             `bson:"domain"`
}

package entity

import "go.mongodb.org/mongo-driver/bson/primitive"

type Link struct {
	ID     primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	URL    string             `bson:"url" json:"url,omitempty"`
	Hash   string             `bson:"hash" json:"hash,omitempty"`
	Domain string             `bson:"domain" json:"domain,omitempty"`
}

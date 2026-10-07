package entity

import "go.mongodb.org/mongo-driver/v2/bson"

type Session struct {
	TokenId *bson.ObjectID `bson:"_id"`
}

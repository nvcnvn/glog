package dbctx

import (
	"labix.org/v2/mgo/bson"
)

type Photo struct {
	Id            bson.ObjectId `bson:"_id"`
	PHash         int64
	Part0         int8
	Part1         int8
	Part2         int8
	Part3         int8
	Part4         int8
	Part5         int8
	Part6         int8
	Part7         int8
	CheckSum      string
	Location      string
	Description   string
	SavedLocation string
	SavedId       bson.ObjectId `bson:"_id"`
}

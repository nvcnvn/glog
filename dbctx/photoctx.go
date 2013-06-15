package dbctx

import (
	"github.com/bufio/toys/model"
)

type Photo struct {
	Id            model.Identifier `bson:"_id"`
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
	SavedId       model.Identifier `bson:"_id"`
	
}

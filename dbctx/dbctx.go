package dbctx

import (
	"github.com/bufio/mtoy"
	"github.com/bufio/toys/model"
	"github.com/nvcnvn/feeds"
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
	"time"
)

type Category struct {
	Id         bson.ObjectId `bson:"_id"`
	Name       string
	Count      uint
	LastThread bson.ObjectId `bson:",omitempty"`
	LastUpdate time.Time
}

func (c *Category) GetId() model.Identifier {
	return mtoy.ID{c.Id}
}

type Thread struct {
	Id      bson.ObjectId `bson:"_id"`
	CatId   bson.ObjectId
	Content string
	Tags    []string
	feeds.Item
}

func (t *Thread) GetId() model.Identifier {
	return mtoy.ID{t.Id}
}

type DBContext struct {
	catColl *mgo.Collection
	thrColl *mgo.Collection
}

func NewDBContext(database *mgo.Database) *DBContext {
	db := &DBContext{}
	db.catColl = database.C("test_category")
	db.thrColl = database.C("test_thread")
	return db
}

func (db *DBContext) SaveCategory(cat *Category) error {
	if !cat.GetId().Valid() {
		cat.Id = bson.NewObjectId()
	}
	return db.catColl.Insert(cat)
}

func (db *DBContext) GetAllCategory() ([]Category, error) {
	var results []Category
	err := db.catColl.Find(nil).All(&results)
	if err == nil {
		return results, nil
	}
	return nil, err
}

func (db *DBContext) SaveThread(thr *Thread) error {
	catUpdateQuery := make(bson.M)
	now := time.Now()
	if !thr.Id.Valid() {
		//insert new thread
		thr.Id = bson.NewObjectId()
		thr.Created = now
		catUpdateQuery["$inc"] = bson.M{"count": 1}
	}

	thr.Updated = now
	thr.Link = &feeds.Link{Href: "/thread?id=" + thr.GetId().Encode()}
	thr.Author = &feeds.Author{"demotration", "nguyen@example.com"}
	catUpdateQuery["$set"] = bson.M{"lastthread": thr.Id}
	err := db.catColl.UpdateId(thr.CatId, catUpdateQuery)
	if err != nil {
		return err
	}

	return db.thrColl.Insert(thr)
}

func (db *DBContext) NewestThreads(n int) ([]Thread, error) {
	results := make([]Thread, 0, n)
	err := db.thrColl.Find(nil).Sort("-_id").Limit(n).All(&results)
	if err == nil {
		return results, nil
	}
	return nil, err
}

func (db *DBContext) GetThread(id bson.ObjectId) (*Thread, error) {
	result := &Thread{}
	err := db.thrColl.FindId(id).One(result)
	if err == nil {
		return result, nil
	}
	return nil, err
}

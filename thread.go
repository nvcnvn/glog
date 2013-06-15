package main

import (
	"github.com/bufio/mtoy"
	//"github.com/bufio/toys/util/forms"
	"github.com/nvcnvn/glog/dbctx"
	"labix.org/v2/mgo/bson"
)

func NewThread(c *controller) {
	data := c.ViewData("New Blog")
	CatLst, _ := c.db.GetAllCategory()
	data["CatLst"] = CatLst
	c.View("newthread_form.tmpl", data)
}

func NewThread2(c *controller) {
	// cat := dbctx.Category{}
	// cat.Name = "test cat"
	// c.Print(c.db.SaveCategory(&cat))
	thr := dbctx.Thread{}
	if idStr := c.Post("CatId", false); bson.IsObjectIdHex(idStr) {
		thr.CatId = &mtoy.ID{bson.ObjectIdHex(idStr)}
	} else {
		c.View("newthread_error.tmpl", "")
		return
	}
	thr.Content = c.Post("Content", true)
	thr.Tags = []string{"test", "demo"}
	thr.Description = c.Post("Item.Description", true)
	thr.Title = c.Post("Item.Title", true)
	if err := c.db.SaveThread(&thr); err != nil {
		c.View("newthread_error.tmpl", err.Error())
	}
	c.Redirect("/thread?id="+thr.Id.Encode(), 303)
}

func ViewThread(c *controller) {
	idStr := c.Get("id", false)
	if !bson.IsObjectIdHex(idStr) {
		c.Print("Not found")
		return
	}
	id := bson.ObjectIdHex(idStr)
	thr, err := c.db.GetThread(id)
	if err != nil {
		c.Print("Not found")
	}
	data := c.ViewData("View thread")
	data["Thread"] = thr
	c.View("viewthread.tmpl", data)
}

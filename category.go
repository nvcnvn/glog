package main

import (
	"github.com/nvcnvn/gotest/dbctx"
)

func NewCat(c *controller) {
	c.View("newcat_form.tmpl", c.ViewData("New Category"))
}

func NewCat2(c *controller) {
	cat := dbctx.Category{}
	cat.Name = c.Post("Name", true)
	c.db.SaveCategory(&cat)
}

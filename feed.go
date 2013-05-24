package main

import (
	"github.com/gorilla/feeds"
	"time"
)

func Feed(c *controller) {
	feed := &feeds.Feed{
		Title:       "Demo RSS/Atom Golang MongoDB OpenSihft",
		Link:        &feeds.Link{Href: "http://gotest-openvn.rhcloud.com/"},
		Description: "Khong biet noi gi",
		Author:      &feeds.Author{"Cao Nguyen", "nguyen@open-vn.org"},
		Created:     time.Now(),
		Copyright:   "copyright © no one",
	}

	thrs, err := c.db.NewestThreads(10)
	if err != nil {
		return
	}
	items := make([]*feeds.Item, len(thrs), len(thrs))
	for i := 0; i < len(thrs); i++ {
		items[i] = &thrs[i].Item
	}
	feed.Items = items

	var xml string
	if c.Get("format", false) == "rss" {
		xml, _ = feed.ToRss()
	} else {
		xml, _ = feed.ToAtom()
	}
	c.Response.Header().Set("Content-type", "application/xml")
	c.Print(xml)
}

package main

import (
	"github.com/gorilla/feeds"
	"time"
)

func Feed(c *controller) {
	feed := &feeds.Feed{
		Title:       "jmoiron.net blog",
		Link:        &feeds.Link{Href: "http://jmoiron.net/blog"},
		Description: "discussion about tech, footie, photos",
		Author:      &feeds.Author{"Jason Moiron", "jmoiron@jmoiron.net"},
		Created:     time.Now(),
		Copyright:   "This work is copyright © Benjamin Button",
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

	rss, _ := feed.ToRss()
	c.Response.Header().Set("Content-type", "application/xml")
	c.Print(rss)
}

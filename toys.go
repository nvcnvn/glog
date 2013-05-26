package main

import (
	"github.com/bufio/mongotoys/mgoauth"
	"github.com/bufio/mongotoys/mgosessions"
	"github.com/bufio/toys"
	"github.com/bufio/toys/secure/membership"
	"github.com/bufio/toys/secure/membership/sessions"
	"github.com/bufio/toys/view"
	"github.com/nvcnvn/glog/dbctx"
	"labix.org/v2/mgo"
	"net/http"
	"path"
	"time"
)

const (
	dbname string = "test"
)

type route struct {
	pattern string
	fn      func(*controller)
}

type controller struct {
	toys.Controller
	sess sessions.Provider
	auth membership.Authenticater
	tmpl *view.View
	db   *dbctx.DBContext
}

func (c *controller) View(page string, data interface{}) error {
	return c.tmpl.Load(c, page, data)
}

type handler struct {
	path           string
	dbsess         *mgo.Session
	tmpl           *view.View
	_subRoutes     []route
	_defaultHandle func(*controller)
}

func (h *handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	c := controller{}
	c.Init(w, r)
	c.SetPath(h.path)

	dbsess := h.dbsess.Clone()
	defer dbsess.Close()

	//database collection (table)
	database := dbsess.DB(dbname)
	sessColl := database.C("toysSession")
	userColl := database.C("toysUser")
	rememberColl := database.C("toysUserRemember")

	//web session
	c.sess = mgosessions.NewMgoProvider(w, r, sessColl)

	//web authenthicator
	c.auth = mgoauth.NewAuthDBCtx(w, r, c.sess, userColl, rememberColl)

	//view template
	c.tmpl = h.tmpl

	//db context
	c.db = dbctx.NewDBContext(database)

	//process
	for _, rt := range h._subRoutes {
		if match(path.Join(h.path, rt.pattern), r.URL.Path) {
			rt.fn(&c)
			return
		}
	}
	h._defaultHandle(&c)
}

// Handler returns a http.Handler
func Handler(path string, dbsess *mgo.Session, tmpl *view.View) *handler {
	h := &handler{}
	h.dbsess = dbsess
	h.tmpl = tmpl
	h.path = path
	h.initSubRoutes()

	dbsess.DB(dbname).C("toysSession").EnsureIndex(mgo.Index{
		Key:         []string{"lastactivity"},
		ExpireAfter: 7200 * time.Second,
	})

	dbsess.DB(dbname).C("toysUser").EnsureIndex(mgo.Index{
		Key:    []string{"email"},
		Unique: true,
	})

	return h
}

// match is a wrapper function for path.Math
func match(pattern, name string) bool {
	ok, err := path.Match(pattern, name)
	if err != nil {
		return false
	}
	return ok
}

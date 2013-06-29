package main

import (
	"github.com/kidstuff/mtoy/mgoauth"
	"github.com/kidstuff/mtoy/mgosessions"
	"github.com/kidstuff/toys"
	"github.com/kidstuff/toys/secure/membership"
	"github.com/kidstuff/toys/secure/membership/sessions"
	"github.com/kidstuff/toys/view"
	"github.com/nvcnvn/glog/dbctx"
	"labix.org/v2/mgo"
	"net/http"
	"path"
	"time"
)

type route struct {
	pattern string
	fn      func(*controller)
}

type controller struct {
	dbsess *mgo.Session
	toys.Controller
	sess sessions.Provider
	auth membership.Authenticater
	tmpl *view.View
	db   *dbctx.DBContext
}

func (c *controller) View(page string, data interface{}) error {
	return c.tmpl.Load(c, page, data)
}

func (c *controller) Close() {
	c.dbsess.Close()
}

type handler struct {
	path           string
	dbsess         *mgo.Session
	tmpl           *view.View
	_subRoutes     []route
	_defaultHandle func(*controller)
	notifer        membership.Notificater
}

func (h *handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if h.path == r.URL.Path {
		c := h.newcontroller(w, r)
		h._defaultHandle(&c)
		c.Close()
		return
	}
	for _, rt := range h._subRoutes {
		if match(path.Join(h.path, rt.pattern), r.URL.Path) {
			c := h.newcontroller(w, r)
			rt.fn(&c)
			c.Close()
			return
		}
	}
}

func (h *handler) newcontroller(w http.ResponseWriter, r *http.Request) controller {
	c := controller{}
	c.Init(w, r)
	c.SetPath(h.path)

	dbsess := h.dbsess.Clone()
	c.dbsess = dbsess
	//database collection (table)
	database := dbsess.DB(DBNAME)
	sessColl := database.C("toysSession")
	userColl := database.C("toysUser")
	rememberColl := database.C("toysUserRemember")

	//web session
	c.sess = mgosessions.NewMgoProvider(w, r, sessColl)
	//web authenthicator
	c.auth = mgoauth.NewAuthDBCtx(w, r, c.sess, userColl, rememberColl)
	c.auth.SetNotificater(h.notifer)
	//view template
	c.tmpl = h.tmpl
	//db context
	c.db = dbctx.NewDBContext(database)
	return c
}

// Handler returns a http.Handler
func Handler(path string, dbsess *mgo.Session, tmpl *view.View) *handler {
	h := &handler{}
	h.dbsess = dbsess
	h.tmpl = tmpl
	h.path = path
	h.initSubRoutes()

	h.notifer = membership.NewSMTPNotificater(
		CONFIG.Get("smtp_email").(string),
		CONFIG.Get("smtp_pass").(string),
		CONFIG.Get("smtp_addr").(string),
		int(CONFIG.Get("smtp_port").(float64)),
	)

	dbsess.DB(DBNAME).C("toysSession").EnsureIndex(mgo.Index{
		Key:         []string{"lastactivity"},
		ExpireAfter: 7200 * time.Second,
	})

	dbsess.DB(DBNAME).C("toysUser").EnsureIndex(mgo.Index{
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

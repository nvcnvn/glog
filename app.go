package main

import (
	"fmt"
	"github.com/kidstuff/toys/confg"
	_ "github.com/kidstuff/toys/confg/jsonconfg"
	"github.com/kidstuff/toys/locale"
	"github.com/kidstuff/toys/view"
	"labix.org/v2/mgo"
	"net/http"
	_ "net/http/pprof"
	"os"
)

var (
	CONFIG      confg.Configurator
	DBNAME      string = "test"
	MODELDRIVER string = "mtoy"
)

func main() {
	var err error
	// Configuration variable
	CONFIG, err = confg.Open("jsonconfg", "glog-confg.json")
	if err != nil {
		panic("mgoauth: cannot load mgoauth-confg.json " + err.Error())
	}

	var (
		// host
		host = os.Getenv("OPENSHIFT_INTERNAL_IP") + ":" + os.Getenv("OPENSHIFT_INTERNAL_PORT")
		// cnnStr the connection string to MongoDB
		cnnStr = os.Getenv("OPENSHIFT_MONGODB_DB_URL")
		// langRoot the path to language folder in file system
		langRoot = os.Getenv("OPENSHIFT_REPO_DIR") + "language"
		// langDefaultSet the default language set
		langDefaultSet = "en"
		// tmplDefaultSet the path to template folder in files system
		tmplRoot = os.Getenv("OPENSHIFT_REPO_DIR") + "template"
		// tmplDefaultSet the default template set
		tmplDefaultSet = "default"
		// rsrcRoot the path to static folder in file system
		rsrcRoot = os.Getenv("OPENSHIFT_REPO_DIR") + "statics"
		// rsrcPrefix the URL path for static file server
		rsrcPrefix = "/statics/"
		//toysignPath the URL path for toysign
		toysignPath = "/"
	)

	//database session
	dbsess, err := mgo.Dial(cnnStr)
	if err != nil {
		panic(err)
	}
	defer dbsess.Close()

	//multi language support
	lang := locale.NewLang(langRoot)
	if err := lang.Parse(langDefaultSet); err != nil {
		fmt.Println(err.Error())
	}

	//template for cms
	tmpl := view.NewView(tmplRoot)
	tmpl.SetLang(lang)
	tmpl.HandleResource(rsrcPrefix, rsrcRoot)
	tmpl.Watch = true
	if err := tmpl.Parse(tmplDefaultSet); err != nil {
		fmt.Println(err.Error())
	}

	http.Handle(toysignPath, Handler(toysignPath, dbsess, tmpl))
	http.ListenAndServe(host, nil)
}

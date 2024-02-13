package main

import (
	"fmt"
	"github.com/alexedwards/scs/v2"
	"insta/pkg/config"
	"insta/pkg/handler"
	"insta/pkg/render"
	"log"
	"net/http"
)

// change port here
const portNumber = ":80"

var app config.AppConfig
var session *scs.SessionManager

func main() {

	//change when going live
	app.InProduction = false

	//
	//session
	//session = scs.New()
	//session.Lifetime = 24 * time.Hour
	//session.Cookie.Persist = true
	//session.Cookie.SameSite = http.SameSiteLaxMode
	//session.Cookie.Secure = app.InProduction
	//
	//app.Session = session

	//cache
	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache")
	}

	app.TemplateCache = tc
	//activate caching
	app.UseCache = false

	repo := handler.NewRepo(&app)
	handler.NewHandlers(repo)

	render.NewTemplates(&app)

	fmt.Println(fmt.Sprintf("Starting application on port %s", portNumber))
	srv := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}
	err = srv.ListenAndServe()
	if err != nil {
		log.Fatalln(err)
	}

}

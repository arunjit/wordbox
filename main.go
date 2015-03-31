package main

import (
	"net/http"

	"appengine"

	"github.com/go-martini/martini"
	"github.com/martini-contrib/binding"
	"github.com/martini-contrib/render"
)

// As defined in app.yaml
const APIBaseURL = "/api"

// Context is the request context.
type Context struct {
	Request        *http.Request
	ResponseWriter http.ResponseWriter
	MartiniCtx     martini.Context
	AppEngineCtx   appengine.Context
}

func init() {
	m := martini.New()
	m.Use(martini.Logger())
	m.Use(martini.Recovery())
	m.Use(render.Renderer(render.Options{
		IndentJSON: martini.Env == martini.Dev,
	}))
	m.Use(func(c martini.Context, r *http.Request, rw http.ResponseWriter) {
		ctx := &Context{r, rw, c, appengine.NewContext(r)}
		c.Map(ctx)
	})

	r := martini.NewRouter()
	r.Group(APIBaseURL, func(api martini.Router) {
		api.Get("/words", binding.Form(GetParams{}), Get)
		api.Post("/words", binding.Json(Word{}), Add)
	})

	m.MapTo(r, (*martini.Routes)(nil))
	m.Action(r.Handle)

	http.Handle("/", m)
}

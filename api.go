package main

import (
	"net/http"

	"github.com/martini-contrib/render"
)

// Defaults
const (
	DefaultResultLimit = 1
	MinimumResultLimit = 1
	MaximumResultLimit = 10
)

// GetParams are the parameters for a query
type GetParams struct {
	Limit int `form:"s"`
}

// Get gets a word from the master wordlist.
func Get(p GetParams, r render.Render) {
	if p.Limit < MinimumResultLimit || p.Limit > MaximumResultLimit {
		p.Limit = DefaultResultLimit
	}
	result := make([]*Word, p.Limit)
	for i := 0; i < p.Limit; i++ {
		result[i] = NewWord("bazinga")
	}
	r.JSON(http.StatusOK, result)
}

// Add adds a new word to the master wordlist.
func Add(c *Context, word Word) {
	_, err := word.Save(c.AppEngineCtx)
	if err != nil {
		http.Error(c.ResponseWriter, err.Error(), http.StatusInternalServerError)
		return
	}
	// c.ResponseWriter.Header().Set("Location", "/api/words/"+id)
	c.ResponseWriter.WriteHeader(http.StatusCreated)
}

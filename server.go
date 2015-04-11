package wordbox

import (
	"log"

	"github.com/GoogleCloudPlatform/go-endpoints/endpoints"
)

func init() {
	wordService := &WordService{}
	api, err := endpoints.RegisterService(wordService,
		"wordbox", "v1", "A box of words", true)
	if err != nil {
		log.Fatalf("Register service: %v", err)
	}

	register := func(orig, name, method, path, desc string) {
		m := api.MethodByName(orig)
		if m == nil {
			log.Fatalf("Missing method %s", orig)
		}
		i := m.Info()
		i.Name, i.HTTPMethod, i.Path, i.Desc = name, method, path, desc
	}

	register("Get", "get", "GET", "words/one", "Get a word.")
	register("Add", "add", "PUT", "words/new", "Add new words.")
	register("Count", "count", "GET", "words/count", "Count all words.")
	endpoints.HandleHTTP()
}

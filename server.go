package wordbox

import (
	"log"

	"github.com/GoogleCloudPlatform/go-endpoints/endpoints"
)

func init() {
	wordService := &WordService{}
	api, err := endpoints.RegisterService(wordService, "words", "v1", "Words API", true)
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

	register("Get", "words.list", "GET", "words", "List words.")
	register("Add", "words.add", "PUT", "words", "Add a word.")
	register("Count", "words.count", "GET", "words/count", "Count all words.")
	endpoints.HandleHTTP()
}

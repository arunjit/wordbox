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

	register("GetPublic", "getpublic", "GET", "words/public", "Get a word from the master wordlist.")
	register("AddPublic", "addpublic", "PUT", "words/public", "Add new words to the master wordlist.")
	register("CountPublic", "countpublic", "GET", "words/public/count",
		"Count all words in the master wordlist.")
	endpoints.HandleHTTP()
}

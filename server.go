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

	register("Add", "add", "PUT", "words/{wordlist}", "Add new words to a named wordlist.")
	register("AddPublic", "addpublic", "PUT", "words/public", "Add new words to the master wordlist.")
	register("Count", "count", "GET", "words/{wordlist}/count",
		"Count all words in a named wordlist.")
	register("CountPublic", "countpublic", "GET", "words/public/count",
		"Count all words in the master wordlist.")
	register("Get", "get", "GET", "words/{wordlist}", "Get a word from a named wordlist.")
	register("GetPublic", "getpublic", "GET", "words/public", "Get a word from the master wordlist.")
	endpoints.HandleHTTP()
}

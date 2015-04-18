package wordbox

import (
	"log"

	"github.com/GoogleCloudPlatform/go-endpoints/endpoints"
)

func init() {
	config := GetConfig("secrets.json")
	auth := &EndpointsAuth{config}
	wordboxAPI := &WordboxAPI{auth}
	api, err := endpoints.RegisterService(wordboxAPI,
		"wordbox", "v1", "A box of words", true)
	if err != nil {
		log.Fatalf("Register service: %v", err)
	}

	register := func(orig, name, method, path, desc string, restrict bool) {
		m := api.MethodByName(orig)
		if m == nil {
			log.Fatalf("Missing method %s", orig)
		}
		i := m.Info()
		i.Name, i.HTTPMethod, i.Path, i.Desc = name, method, path, desc
		if restrict {
			i.Scopes, i.ClientIds, i.Audiences = config.Scopes, config.ClientIds, config.Audiences
		}
	}

	register("Add", "add", "PUT", "words/{wordlist}",
		"Add new words to a named wordlist.", true)
	register("AddPublic", "addpublic", "PUT", "words/public",
		"Add new words to the master wordlist.", true)
	register("Count", "count", "GET", "words/{wordlist}/count",
		"Count all words in a named wordlist.", true)
	register("CountPublic", "countpublic", "GET", "words/public/count",
		"Count all words in the master wordlist.", false)
	register("Get", "get", "GET", "words/{wordlist}",
		"Get a word from a named wordlist.", true)
	register("GetPublic", "getpublic", "GET", "words/public",
		"Get a word from the master wordlist.", false)
	register("SetUsed", "setused", "PUT", "words/setused",
		"Mark a word as used", true)
	endpoints.HandleHTTP()
}

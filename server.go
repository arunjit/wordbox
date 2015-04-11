package wordbox

import (
	"log"

	"github.com/GoogleCloudPlatform/go-endpoints/endpoints"
)

// const clientId = "YOUR-CLIENT-ID"

var (
	scopes    = []string{endpoints.EmailScope}
	clientIds = []string{ /*clientId, */ endpoints.APIExplorerClientID}
	audiences = []string{ /*clientId*/ }
)

type endpointsAuth struct{}

func (a *endpointsAuth) CheckAuth(c endpoints.Context) error {
	u, err := endpoints.CurrentUser(c, scopes, audiences, clientIds)
	if err != nil {
		return endpoints.UnauthorizedError
	}
	if u == nil {
		return endpoints.UnauthorizedError
	}
	c.Debugf("Current user: %#v", u)
	if u.Email != "AN EMAIL" { // TODO(arunjit): The obvious
		return endpoints.ForbiddenError
	}
	return nil
}

func init() {
	wordService := &WordService{&endpointsAuth{}}
	api, err := endpoints.RegisterService(wordService,
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
			i.Scopes, i.ClientIds, i.Audiences = scopes, clientIds, audiences
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
	endpoints.HandleHTTP()
}

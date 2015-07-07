package wordbox

import (
	"encoding/json"
	"io/ioutil"
	"log"

	"github.com/GoogleCloudPlatform/go-endpoints/endpoints"
)

// Config is the auth config.
type Config struct {
	Scopes        []string `json:"scopes"`
	ClientIds     []string `json:"client_ids"`
	Audiences     []string `json:"audiences"`
	AllowedEmails []string `json:"allowed_emails"`
}

// EndpointsAuth is an [Authenticator] for endpoints.
type EndpointsAuth struct {
	Config Config
}

// CheckAuth is the implementation of the [Authenticator] interface for endpoints.
func (a *EndpointsAuth) CheckAuth(c endpoints.Context) error {
	user, err := endpoints.CurrentUser(c, a.Config.Scopes, a.Config.Audiences, a.Config.ClientIds)
	if err != nil || user == nil {
		return endpoints.UnauthorizedError
	}
	if !contains(a.Config.AllowedEmails, user.Email) {
		return endpoints.ForbiddenError
	}
	return nil
}

// GetConfig loads a configuration file.
func GetConfig(path string) Config {
	contents, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatalf("Read config: %v", err)
	}
	var config Config
	if err := json.Unmarshal(contents, &config); err != nil {
		log.Fatalf("Parse config: %v", err)
	}
	config.Scopes = append(config.Scopes, endpoints.EmailScope)
	config.ClientIds = append(config.ClientIds, endpoints.APIExplorerClientID)
	return config
}

func contains(list []string, value string) bool {
	for _, item := range list {
		if item == value {
			return true
		}
	}
	return false
}

package wordbox

import "github.com/GoogleCloudPlatform/go-endpoints/endpoints"

// Authenticator is an interface to validate an endpoints request.
// TODO(arunjit): Make generic.
type Authenticator interface {
	CheckAuth(c endpoints.Context) error
}

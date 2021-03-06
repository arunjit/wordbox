package wordbox

import "github.com/GoogleCloudPlatform/go-endpoints/endpoints"

var (
	// ErrNotImplemented error when something isn't implemented.
	ErrNotImplemented = endpoints.NewAPIError("NotImplemented", "Not implemented", 501)

	// ErrMissingValue error when a required value is missing.
	ErrMissingValue = endpoints.NewBadRequestError("Missing value")
)

// WordboxAPI is the endpoints service.
type WordboxAPI struct {
	auth Authenticator
}

// GetReq is the request struct to fetch a word from a named wordlist.
type GetReq struct {
	WordList string `json:"wordlist" endpoints:"req"`
}

// AddReq is the request struct to add new words to a named wordlist.
type AddReq struct {
	WordList string   `json:"wordlist" endpoints:"req"`
	Words    []string `json:"words" endpoints:"req"`
}

// AddPublicReq is the request struct to add new words to the master wordlist.
type AddPublicReq struct {
	Words []string `json:"words" endpoints:"req"`
}

// SetUsedReq is the request struct to mark a word as used.
type SetUsedReq struct {
	ID string `json:"id"`
}

// Get fetches a word from a named wordlist.
func (s *WordboxAPI) Get(c endpoints.Context, r *GetReq) (*Word, error) {
	if err := s.auth.CheckAuth(c); err != nil {
		return nil, err
	}
	return nil, ErrNotImplemented
}

// GetPublic fetches a word from the master wordlist.
func (s *WordboxAPI) GetPublic(c endpoints.Context) (*Word, error) {
	return PublicWord(c)
}

// Add adds new words to a named wordlist.
func (s *WordboxAPI) Add(c endpoints.Context, r *AddReq) error {
	if err := s.auth.CheckAuth(c); err != nil {
		return err
	}
	return ErrNotImplemented
}

// AddPublic adds new words to the master wordlist.
func (s *WordboxAPI) AddPublic(c endpoints.Context, r *AddPublicReq) error {
	if err := s.auth.CheckAuth(c); err != nil {
		return err
	}
	words := make([]*Word, len(r.Words))
	for i := 0; i < len(r.Words); i++ {
		if r.Words[i] == "" {
			return ErrMissingValue
		}
		words[i] = &Word{Word: r.Words[i], Public: true}
	}
	return AddAllWords(c, words)
}

// Count is a count of words in a wordlist.
type Count struct {
	N int `json:"count"`
}

// Count counts the words in a named wordlist.
func (s *WordboxAPI) Count(c endpoints.Context) (*Count, error) {
	if err := s.auth.CheckAuth(c); err != nil {
		return nil, err
	}
	return nil, ErrNotImplemented
}

// CountPublic counts the words in the master wordlist.
func (s *WordboxAPI) CountPublic(c endpoints.Context) (*Count, error) {
	n, err := PublicWordCount(c)
	if err != nil {
		return nil, err
	}
	return &Count{n}, nil
}

// SetUsed sets a word as used.
func (s *WordboxAPI) SetUsed(c endpoints.Context, req *SetUsedReq) error {
	if err := s.auth.CheckAuth(c); err != nil {
		return err
	}
	return ErrNotImplemented
}

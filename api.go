package wordbox

import "github.com/GoogleCloudPlatform/go-endpoints/endpoints"

var (
	// ErrNotImplemented error when something isn't implemented.
	ErrNotImplemented = endpoints.NewAPIError("NotImplemented", "Not implemented", 501)

	// ErrMissingValue error when a required value is missin.
	ErrMissingValue = endpoints.NewBadRequestError("Missing value")
)

// WordService is the endpoints service.
type WordService struct {
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

// Get fetches a word from a named wordlist.
func (s *WordService) Get(c endpoints.Context, r *GetReq) (*Word, error) {
	err := s.auth.CheckAuth(c)
	if err != nil {
		return nil, err
	}
	return nil, ErrNotImplemented
}

// GetPublic fetches a word from the master wordlist.
func (s *WordService) GetPublic(c endpoints.Context) (*Word, error) {
	return PublicWord(c)
}

// Add adds new words to a named wordlist.
func (s *WordService) Add(c endpoints.Context, r *AddReq) error {
	err := s.auth.CheckAuth(c)
	if err != nil {
		return err
	}
	return ErrNotImplemented
}

// AddPublic adds new words to the master wordlist.
func (s *WordService) AddPublic(c endpoints.Context, r *AddPublicReq) error {
	err := s.auth.CheckAuth(c)
	if err != nil {
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
func (s *WordService) Count(c endpoints.Context) (*Count, error) {
	err := s.auth.CheckAuth(c)
	if err != nil {
		return nil, err
	}
	return nil, ErrNotImplemented
}

// CountPublic counts the words in the master wordlist.
func (s *WordService) CountPublic(c endpoints.Context) (*Count, error) {
	n, err := PublicWordCount(c)
	if err != nil {
		return nil, err
	}
	return &Count{n}, nil
}

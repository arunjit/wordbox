package wordbox

import (
	"errors"

	"github.com/GoogleCloudPlatform/go-endpoints/endpoints"
)

var (
	// ErrNotImplemented error when something isn't implemented
	ErrNotImplemented = errors.New("Not implemented")
)

// WordService is the endpoints service.
type WordService struct{}

// GetReq is the request struct to fetch a word from a named wordlist.
type GetReq struct {
	WordList string `json:"wordlist"`
}

// AddReq is the request struct to add new words.
type AddReq struct {
	Words []string `json:"words"`
}

// Get fetches a word from a named wordlist.
func (s *WordService) Get(c endpoints.Context) (*Word, error) {
	return nil, ErrNotImplemented
}

// GetPublic fetches a word from the master wordlist.
func (s *WordService) GetPublic(c endpoints.Context) (*Word, error) {
	return PublicWord(c)
}

// Add adds new words to a named wordlist.
func (s *WordService) Add(c endpoints.Context, r *AddReq) error {
	return ErrNotImplemented
}

// AddPublic adds new words to the master wordlist.
func (s *WordService) AddPublic(c endpoints.Context, r *AddReq) error {
	words := make([]*Word, len(r.Words))
	for i := 0; i < len(r.Words); i++ {
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

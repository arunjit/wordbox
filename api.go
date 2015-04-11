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

// GetPublic fetches a word from the master wordlist.
func (s *WordService) GetPublic(c endpoints.Context) (*Word, error) {
	return PublicWord(c)
}

// AddPublicReq is the request struct to add new words.
type AddPublicReq struct {
	Words []string `json:"words"`
}

// AddPublic adds new words to the master wordlist.
func (s *WordService) AddPublic(c endpoints.Context, r *AddPublicReq) error {
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

// CountPublic counts the words in a named wordlist/the master wordlist.
func (s *WordService) CountPublic(c endpoints.Context) (*Count, error) {
	n, err := PublicWordCount(c)
	if err != nil {
		return nil, err
	}
	return &Count{n}, nil
}

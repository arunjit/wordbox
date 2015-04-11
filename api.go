package wordbox

import (
	"errors"

	"github.com/GoogleCloudPlatform/go-endpoints/endpoints"
)

var (
	// ErrNotImplemented error when something isn't implemented
	ErrNotImplemented = errors.New("Not implemented")

	// ErrMustBePublicOrWordList error when trying to get/add a word that is neither public nor
	// in a named wordlist.
	ErrMustBePublicOrWordList = errors.New("Word must be either public or belong to a named list")
)

// WordService is the endpoints service.
type WordService struct{}

// GetReq is the request struct for fetching a wordlist.
type GetReq struct {
	// If true, fetches a word from the master wordlist.
	Public bool `json:"public"`

	// If set, fetches a word from the given [WordList].
	// If [Public] is true and a [WordList] is set, fetches a word from
	// [WordList] that is also public, or nil if not found.
	// Not implemented
	WordList string `json:"wordlist"`
}

// Get fetches words from the master wordlist.
// TODO(arunjit): namespaced wordlists.
func (s *WordService) Get(c endpoints.Context, r *GetReq) (*Word, error) {
	if !r.Public {
		if r.WordList == "" {
			return nil, ErrMustBePublicOrWordList
		}
		return nil, ErrNotImplemented
	}
	return PublicWord(c)
}

// AddReq is the request struct to add a new word.
type AddReq struct {
	// The word to add
	Word string `json:"word"`

	// If true, the word gets added to the master wordlist.
	Public bool `json:"public"`

	// If set, the word gets added to a named wordlist.
	// If [Public] is true and a [WordList] is set, both actions are performed.
	// Not implemented
	WordList string `json:"wordlist"`
}

// Add adds a new word to the master wordlist.
func (s *WordService) Add(c endpoints.Context, r *AddReq) error {
	if !r.Public {
		if r.WordList == "" {
			return ErrMustBePublicOrWordList
		}
		return ErrNotImplemented
	}
	word := &Word{Word: r.Word, Public: r.Public}
	_, err := word.Save(c)
	return err
}

// Count is a count of words in the master wordlist.
type Count struct {
	N int `json:"count"`
}

// Count counts the words in the master wordlist.
func (s *WordService) Count(c endpoints.Context) (*Count, error) {
	n, err := GetWordCount(c)
	if err != nil {
		return nil, err
	}
	return &Count{n}, nil
}

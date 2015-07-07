package wordbox

import (
	"errors"
	"math/rand"

	"appengine"
	"appengine/datastore"
)

const (
	// wordKind is the Datastore kind for the word
	wordKind = "W"

	// wordMaxUses is the maximum times a word can be used.
	wordMaxUses = 10
)

var (
	// ErrNoMoreWords ...
	ErrNoMoreWords = errors.New("No words left in the list")
)

// Word is a word entity.
type Word struct {
	Key *datastore.Key `datastore:"-" json:"id"`

	// The word
	Word string `datastore:"w" json:"word"`

	// The number of times this word has been retrieved.
	Uses int `datastore:"u" json:"uses"`

	// Whether the word should show up in the master wordlist.
	Public bool `datastore:"p" json:"public"`
}

// PublicWord fetches a word from the master wordlist.
func PublicWord(c appengine.Context) (*Word, error) {
	for i := 0; i < wordMaxUses; i++ {
		// TODO(arunjit): Improve perf by trading memory for RPC calls.
		// Create a projection query for all entities where u<wordMaxUses and
		// ascending order of uses, on key and "u". Map `Uses` to an array of
		// keys, then iterate over the keys of the map, starting at 0.
		q := datastore.NewQuery(wordKind).
			Filter("p =", true).
			Filter("u =", i).
			KeysOnly()
		keys, err := q.GetAll(c, nil)
		if err != nil {
			return nil, err
		}
		if len(keys) == 0 {
			continue
		}

		key := keys[rand.Intn(len(keys))]

		var word Word
		err = datastore.Get(c, key, &word)
		if err != nil {
			return nil, err
		}
		word.Key = key
		return &word, nil
	}
	return nil, ErrNoMoreWords
}

// PublicWordCount counts all words in the Datastore.
func PublicWordCount(c appengine.Context) (int, error) {
	return datastore.NewQuery(wordKind).Filter("p =", true).Count(c)
}

// AddAllWords adds words in a single batch
func AddAllWords(c appengine.Context, words []*Word) error {
	keys := make([]*datastore.Key, len(words))
	for i := 0; i < len(words); i++ {
		keys[i] = datastore.NewIncompleteKey(c, wordKind, nil)
	}
	_, err := datastore.PutMulti(c, keys, words)
	return err
}

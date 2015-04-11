package wordbox

import (
	"appengine"
	"appengine/datastore"
)

const (
	// wordKind is the Datastore kind for the word
	wordKind = "W"
)

// Word is a word entity.
type Word struct {
	ID string `datastore:"-" json:"id"`

	// The word
	Word string `datastore:"w" json:"word"`

	// The number of times this word has been retrieved.
	Uses int `datastore:"u" json:"uses"`

	// Whether the word should show up in the master wordlist.
	Public bool `datastore:"p" json:"public"`
}

// PublicWord fetches a word from the master wordlist.
func PublicWord(c appengine.Context) (*Word, error) {
	q := datastore.NewQuery(wordKind).
		Filter("p =", true).
		Order("u")
	t := q.Run(c)
	var word Word
	key, err := t.Next(&word)
	if err != nil {
		return nil, err
	}
	word.ID = key.StringID()
	return &word, nil
}

// WordByID gets a word entity by ID.
func WordByID(c appengine.Context, id string) (*Word, error) {
	w := &Word{ID: id}
	return w.Get(c)
}

// Get gets a word given it's ID.
func (w *Word) Get(c appengine.Context) (*Word, error) {
	key, err := datastore.DecodeKey(w.ID)
	if err != nil {
		return nil, err
	}
	if err := datastore.Get(c, key, w); err != nil {
		return nil, err
	}
	return w, nil
}

// Save saves a new word in the Datastore.
func (w *Word) Save(c appengine.Context) (string, error) {
	key, err := datastore.Put(c, datastore.NewIncompleteKey(c, wordKind, nil), w)
	if err != nil {
		return "", err
	}
	return key.Encode(), nil
}

// GetWordCount counts all words in the Datastore.
func GetWordCount(c appengine.Context) (int, error) {
	return datastore.NewQuery(wordKind).Count(c)
}

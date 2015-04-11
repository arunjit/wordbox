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

// PublicWordCount counts all words in the Datastore.
func PublicWordCount(c appengine.Context) (int, error) {
	return datastore.NewQuery(wordKind).Filter("p =", true).Count(c)
}

// Save saves a new word in the Datastore.
func (w *Word) Save(c appengine.Context) (string, error) {
	key, err := datastore.Put(c, datastore.NewIncompleteKey(c, wordKind, nil), w)
	if err != nil {
		return "", err
	}
	return key.Encode(), nil
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

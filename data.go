package main

import (
	"appengine"
	"appengine/datastore"
)

const (
	// WordKind is the Datastore kind for the word
	WordKind = "W"
)

// Word is a word entity.
type Word struct {
	ID   string `datastore:"-" json:"id"`
	Word string `datastore:"w" json:"word"`
}

// NewWord creates a new word entity.
func NewWord(w string) *Word {
	return &Word{"", w}
}

// WordByID gets a word entity by ID.
func WordByID(c appengine.Context, id string) (*Word, error) {
	w := &Word{id, ""}
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
	key, err := datastore.Put(c, datastore.NewIncompleteKey(c, WordKind, nil), w)
	if err != nil {
		return "", err
	}
	return key.Encode(), nil
}

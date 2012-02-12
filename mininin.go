package mininin

import (
	"github.com/fsouza/gokabinet/kc"
)

type Generator interface {
	Generate() string
}

type Shortener struct {
	g Generator
	urlsDb *kc.DB
	visitsDb *kc.DB
}

func NewShortener(g Generator, urlsDbName, visitsDbName string) (*Shortener, error) {
	urlsDb, err := kc.Open(urlsDbName, kc.WRITE)
	if err != nil {
		return nil, err
	}

	visitsDb, err := kc.Open(visitsDbName, kc.WRITE)
	if err != nil {
		return nil, err
	}

	s := &Shortener{
		g: g,
		urlsDb: urlsDb,
		visitsDb: visitsDb,
	}

	return s, nil
}

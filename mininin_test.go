package mininin

import (
	"github.com/fsouza/gokabinet/kc"
	"reflect"
	"testing"
)

func TestShortenerShouldReceiveTheGeneratorInstanceAndTheNameOfTheDatabases(t *testing.T) {
	g := GeneratorMock("random")
	shortener, _ := NewShortener(&g, "/tmp/urls.kch", "/tmp/visits.kch")

	if shortener.g.Generate() != g.Generate() {
		t.Errorf("Should store the generator with shortener")
	}
}

func TestShortenerShouldStoreInternalInstancesOfTheUrlDatabaseInWriteMode(t *testing.T) {
	tmpDb, _ := kc.Open("/tmp/removeme.kch", kc.WRITE)
	defer tmpDb.Close()

	g := GeneratorMock("random")
	shortener, _ := NewShortener(&g, "/tmp/urls.kch", "/tmp/visits.kch")

	typeUrlDb := reflect.TypeOf(shortener.urlsDb)
	if typeUrlDb != reflect.TypeOf(tmpDb) || shortener.urlsDb == nil {
		t.Errorf("Should store the urls database instance with the given name")
	}
}

func TestShortenerShouldStoreInternalInstanceOfTheVisitsDatabaseInWriteMode(t *testing.T) {
	tmpDb, _ := kc.Open("/tmp/removeme.kch", kc.WRITE)
	defer tmpDb.Close()

	g := GeneratorMock("random")
	shortener, _ := NewShortener(&g, "/tmp/urls.kch", "/tmp/visits.kch")

	typeVisitsDb := reflect.TypeOf(shortener.visitsDb)
	if typeVisitsDb != reflect.TypeOf(tmpDb) || shortener.visitsDb == nil {
		t.Errorf("Should store the visits database instance with the given name")
	}
}

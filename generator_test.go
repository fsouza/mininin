package mininin

import (
	"testing"
)

type GeneratorMock string

func (g GeneratorMock) Generate() string {
	return string(g)
}

func TestGeneratorMockShouldGenerateItself(t *testing.T) {
	g := GeneratorMock("$XKD")
	got := g.Generate()
	if got != "$XKD" {
		t.Errorf("expected: $XKD, got: %s", got)
	}
}

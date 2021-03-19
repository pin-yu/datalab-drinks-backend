package persistence

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIces(t *testing.T) {
	ices, _ := NewIcesRepository().ReadIces()

	assert.Equal(t, uint(1), ices.Ices[0].ID)
	assert.Equal(t, "熱", ices.Ices[0].Tag)
}

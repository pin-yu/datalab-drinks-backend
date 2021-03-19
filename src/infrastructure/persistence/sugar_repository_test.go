package persistence

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSugar(t *testing.T) {
	sugars, _ := NewSugarsRepository().ReadSugars()

	assert.Equal(t, uint(1), sugars.Sugars[0].ID)
	assert.Equal(t, "無糖", sugars.Sugars[0].Tag)
}

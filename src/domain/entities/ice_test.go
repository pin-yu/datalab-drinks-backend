package entities

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsHot(t *testing.T) {
	ice1 := Ice{
		ID:  1,
		Tag: "熱",
	}

	ice2 := Ice{
		ID:  2,
		Tag: "去冰",
	}

	assert.True(t, ice1.IsHot())
	assert.False(t, ice2.IsHot())
}

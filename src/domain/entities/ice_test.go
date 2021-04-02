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

func TestIsValidIce(t *testing.T) {
	ice1 := Ice{
		ID:  1,
		Tag: "熱",
	}

	ice2 := Ice{
		ID:  2,
		Tag: "去冰",
	}

	ice3 := Ice{
		ID:  3,
		Tag: "少冰",
	}

	ice4 := Ice{
		ID:  4,
		Tag: "正常冰",
	}

	assert.True(t, ice1.IsValidIce(true, true))
	assert.True(t, ice1.IsValidIce(true, false))
	assert.False(t, ice1.IsValidIce(false, true))

	assert.True(t, ice2.IsValidIce(true, true))
	assert.False(t, ice2.IsValidIce(true, false))
	assert.True(t, ice2.IsValidIce(false, true))

	assert.True(t, ice3.IsValidIce(true, true))
	assert.False(t, ice3.IsValidIce(true, false))
	assert.True(t, ice3.IsValidIce(false, true))

	assert.True(t, ice4.IsValidIce(true, true))
	assert.False(t, ice4.IsValidIce(true, false))
	assert.True(t, ice4.IsValidIce(false, true))
}

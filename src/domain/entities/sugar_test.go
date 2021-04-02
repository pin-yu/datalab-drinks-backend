package entities

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsNormalSugar(t *testing.T) {
	sugar1 := Sugar{
		ID:  1,
		Tag: "無糖",
	}

	sugar2 := Sugar{
		ID:  2,
		Tag: "微糖",
	}

	sugar3 := Sugar{
		ID:  3,
		Tag: "半糖",
	}

	sugar4 := Sugar{
		ID:  4,
		Tag: "全糖",
	}

	assert.False(t, sugar1.IsNormalSugar())
	assert.False(t, sugar2.IsNormalSugar())
	assert.False(t, sugar3.IsNormalSugar())
	assert.True(t, sugar4.IsNormalSugar())
}

func TestIsValidSugar(t *testing.T) {
	sugar1 := Sugar{
		ID:  1,
		Tag: "無糖",
	}

	sugar2 := Sugar{
		ID:  2,
		Tag: "微糖",
	}

	sugar3 := Sugar{
		ID:  3,
		Tag: "半糖",
	}

	sugar4 := Sugar{
		ID:  4,
		Tag: "正常糖",
	}

	assert.True(t, sugar1.IsValidSugar(true))
	assert.False(t, sugar1.IsValidSugar(false))

	assert.True(t, sugar2.IsValidSugar(true))
	assert.False(t, sugar2.IsValidSugar(false))

	assert.True(t, sugar3.IsValidSugar(true))
	assert.False(t, sugar3.IsValidSugar(false))

	assert.True(t, sugar4.IsValidSugar(true))
	assert.True(t, sugar4.IsValidSugar(false))
}

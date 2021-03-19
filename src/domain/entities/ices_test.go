package entities

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAdjustIceList(t *testing.T) {
	ices := Ices{
		Ices: []Ice{
			{ID: 1, Tag: "熱"},
			{ID: 2, Tag: "去冰"},
			{ID: 3, Tag: "少冰"},
			{ID: 4, Tag: "正常冰"},
		},
	}

	assert.Equal(t, uint(1), ices.AdjustIceList(true, true)[0].ID)
	assert.Equal(t, uint(4), ices.AdjustIceList(true, true)[3].ID)

	assert.Equal(t, uint(1), ices.AdjustIceList(true, false)[0].ID)
	shouldPanic(t, &ices)

	assert.Equal(t, uint(2), ices.AdjustIceList(false, true)[0].ID)
}

func shouldPanic(t *testing.T, ices *Ices) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()

	_ = ices.AdjustIceList(true, false)[1]
}

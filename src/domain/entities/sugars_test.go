package entities

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAdjustSugarList(t *testing.T) {
	sugars := Sugars{
		Sugars: []Sugar{
			{ID: 1, Tag: "無糖"},
			{ID: 2, Tag: "微糖"},
			{ID: 3, Tag: "半糖"},
			{ID: 4, Tag: "正常糖"},
		},
	}

	assert.Equal(t, uint(1), sugars.AdjustSugarList(true)[0].ID)
	assert.Equal(t, uint(2), sugars.AdjustSugarList(true)[1].ID)
	assert.Equal(t, uint(4), sugars.AdjustSugarList(false)[0].ID)
}

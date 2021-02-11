package entities

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPrice(t *testing.T) {
	item := Item{
		MediumPrice: 50,
		LargePrice:  100,
	}

	mediumOrder := Order{
		Size: "medium",
		Item: item,
	}

	largeOrder := Order{
		Size: "large",
		Item: item,
	}

	assert.Equal(t, uint(50), mediumOrder.Price())
	assert.Equal(t, uint(100), largeOrder.Price())
}

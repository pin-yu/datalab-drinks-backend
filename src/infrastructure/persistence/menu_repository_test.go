package persistence

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReadMenu(t *testing.T) {
	menu, _ := NewMenuRepository().ReadMenu()

	assert.Equal(t, "2021-spring", menu.MenuVersion)
	assert.Equal(t, "咖啡系列", menu.Menu[0].Series)
	assert.Equal(t, uint(1), menu.Menu[0].Items[0].ID)
	assert.Equal(t, "cama 經典黑咖啡", menu.Menu[0].Items[0].Item)
	assert.Equal(t, uint(60), menu.Menu[0].Items[0].LargePrice)
	assert.Equal(t, uint(45), menu.Menu[0].Items[0].MediumPrice)
	assert.True(t, menu.Menu[0].Items[0].Sugar)
	assert.True(t, menu.Menu[0].Items[0].Cold)
	assert.True(t, menu.Menu[0].Items[0].Hot)
}

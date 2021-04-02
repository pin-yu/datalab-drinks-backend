package responses

import (
	"testing"

	"github.com/pin-yu/datalab-drinks-backend/src/infrastructure/persistence"
	"github.com/stretchr/testify/assert"
)

func TestConvertMenuEntityToResponse(t *testing.T) {
	menu, _ := persistence.NewMenuRepository().ReadMenu()
	sugars, _ := persistence.NewSugarsRepository().ReadSugars()
	ices, _ := persistence.NewIcesRepository().ReadIces()

	menuResponse := ConvertMenuEntityToResponse(menu, sugars, ices)

	assert.Equal(t, uint(5), menuResponse.Menu[0].Items[4].ID)
	assert.Equal(t, "香草拿鐵", menuResponse.Menu[0].Items[4].Item)

	assert.Equal(t, uint(75), menuResponse.Menu[0].Items[4].MediumPrice)
	assert.Equal(t, uint(95), menuResponse.Menu[0].Items[4].LargePrice)

	assert.Equal(t, uint(1), menuResponse.Menu[0].Items[4].Sugars[0].ID)
	assert.Equal(t, "無糖", menuResponse.Menu[0].Items[4].Sugars[0].Tag)
	assert.False(t, menuResponse.Menu[0].Items[4].Sugars[0].Enable)

	assert.Equal(t, uint(4), menuResponse.Menu[0].Items[4].Sugars[3].ID)
	assert.Equal(t, "正常糖", menuResponse.Menu[0].Items[4].Sugars[3].Tag)
	assert.True(t, menuResponse.Menu[0].Items[4].Sugars[3].Enable)

	assert.Equal(t, "熱", menuResponse.Menu[0].Items[4].Ices[0].Tag)
}

package responses

import (
	"github.com/pin-yu/datalab-drinks-backend/src/domain/entities"
)

// MenuResponse represents the domain entity Menu
type MenuResponse struct {
	MenuVersion string           `json:"menu_version"`
	Menu        []SeriesResponse `json:"menu"`
}

// SeriesResponse represents the domain entity Series
type SeriesResponse struct {
	Series string         `json:"series"`
	Items  []ItemResponse `json:"items"`
}

// MenuResponse represents the domain entity Item
type ItemResponse struct {
	ID          uint            `json:"item_id"`
	Item        string          `json:"item" gorm:"not null"`
	MediumPrice uint            `json:"medium_price"`
	LargePrice  uint            `json:"large_price"`
	Sugars      []SugarResponse `json:"sugars"`
	Ices        []IceResponse   `json:"ices"`
}

type SugarResponse struct {
	ID     uint   `json:"sugar_id"`
	Tag    string `json:"sugar_tag"`
	Enable bool   `json:"enable"`
}

type IceResponse struct {
	ID     uint   `json:"ice_id"`
	Tag    string `json:"ice_tag"`
	Enable bool   `json:"enable"`
}

func ConvertMenuEntityToResponse(menu *entities.Menu, sugars *entities.Sugars, ices *entities.Ices) *MenuResponse {
	return &MenuResponse{
		MenuVersion: menu.MenuVersion,
		Menu:        convertSeriesToReponse(menu.Menu, sugars, ices),
	}
}

func convertSeriesToReponse(menu []entities.Series, sugars *entities.Sugars, ices *entities.Ices) []SeriesResponse {
	seriesRes := []SeriesResponse{}

	for _, m := range menu {
		seriesRes = append(seriesRes, SeriesResponse{
			Series: m.Series,
			Items:  convertItemsToResponse(m.Items, sugars, ices),
		})
	}

	return seriesRes
}

func convertItemsToResponse(items []entities.Item, sugars *entities.Sugars, ices *entities.Ices) []ItemResponse {
	itemsRes := []ItemResponse{}

	for _, item := range items {
		itemsRes = append(itemsRes, ItemResponse{
			ID:          item.ID,
			Item:        item.Item,
			MediumPrice: item.MediumPrice,
			LargePrice:  item.LargePrice,
			Sugars:      convertSugarToResponse(sugars, item.Sugar),
			Ices:        convertIceToResponse(ices, item.Hot, item.Cold),
		})
	}

	return itemsRes
}

func convertSugarToResponse(sugars *entities.Sugars, sugarAdjustable bool) []SugarResponse {
	sugarsRes := []SugarResponse{}

	for _, s := range sugars.Sugars {
		sugarsRes = append(sugarsRes, SugarResponse{
			ID:     s.ID,
			Tag:    s.Tag,
			Enable: s.IsValidSugar(sugarAdjustable),
		})
	}

	return sugarsRes
}

func convertIceToResponse(ices *entities.Ices, hotAdjustable bool, iceAdjustable bool) []IceResponse {
	icesRes := []IceResponse{}

	for _, i := range ices.Ices {
		icesRes = append(icesRes, IceResponse{
			ID:     i.ID,
			Tag:    i.Tag,
			Enable: i.IsValidIce(hotAdjustable, iceAdjustable),
		})
	}

	return icesRes
}

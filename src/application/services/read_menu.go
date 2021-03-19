package services

import (
	"net/http"

	"github.com/pin-yu/datalab-drinks-backend/src/infrastructure/persistence"
	"github.com/pin-yu/datalab-drinks-backend/src/interface/responses"
)

// ReadCamaMenu returns http status code and menu in json fomat
func ReadCamaMenu() *responses.Response {
	menu, err := persistence.NewMenuRepository().ReadMenu()
	if err != nil {
		return responses.NewResponse(http.StatusInternalServerError, "error occurs at reading menu", nil)
	}

	sugars, err := persistence.NewSugarsRepository().ReadSugars()
	if err != nil {
		return responses.NewResponse(http.StatusInternalServerError, "error occurs at reading menu", nil)
	}

	ices, err := persistence.NewIcesRepository().ReadIces()
	if err != nil {
		return responses.NewResponse(http.StatusInternalServerError, "error occurs at reading menu", nil)
	}

	// build response of menu
	menuResponse := responses.ConvertMenuEntityToResponse(menu, sugars, ices)

	return responses.NewResponse(http.StatusOK, "ok", menuResponse)
}

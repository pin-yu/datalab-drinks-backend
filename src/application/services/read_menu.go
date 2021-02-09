package services

import (
	"net/http"

	"github.com/pinyu/datalab-drinks-backend/src/infrastructure/persistence"
	"github.com/pinyu/datalab-drinks-backend/src/interface/responses"
)

// ReadCamaMenu returns http status code and menu in json fomat
func ReadCamaMenu() *responses.Response {
	menuRepo := persistence.NewMenuRepository()
	menu, err := menuRepo.ReadMenu()
	if err != nil {
		return responses.NewResponse(http.StatusInternalServerError, "error occurs at reading menu", nil)
	}

	return responses.NewResponse(http.StatusOK, "ok", menu)
}

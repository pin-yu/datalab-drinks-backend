package services

import (
	"github.com/pinyu/datalab-drinks-backend/src/domain/menus"
)

// GetCamaMenus return a map object of Cama drinks
func GetCamaMenus() *menus.CamaMenu {
	return menus.GetMenus()
}

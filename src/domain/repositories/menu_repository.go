package repositories

import (
	"github.com/pin-yu/datalab-drinks-backend/src/domain/entities"
)

// MenuRepository is Menu CRUD interface
type MenuRepository interface {
	ReadMenu() (*entities.Menu, error)
}

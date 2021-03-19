package repositories

import (
	"github.com/pin-yu/datalab-drinks-backend/src/domain/entities"
)

// IcesRepository is Ices interface
type IcesRepository interface {
	ReadIces() (*entities.Ices, error)
}

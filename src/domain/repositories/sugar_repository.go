package repositories

import (
	"github.com/pin-yu/datalab-drinks-backend/src/domain/entities"
)

// SugarsRepository is Sugars interface
type SugarsRepository interface {
	ReadSugars() (*entities.Sugars, error)
}

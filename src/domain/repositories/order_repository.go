package repositories

import (
	"github.com/pin-yu/datalab-drinks-backend/src/domain/entities"
	"github.com/pin-yu/datalab-drinks-backend/src/interface/requests"
)

// OrderRepository is Order CRUD interface
type OrderRepository interface {
	HasOrdered(string) bool

	ValidateItemID(uint) (*entities.Item, error)
	ValidateSugarID(uint) (*entities.Sugar, error)
	ValidateIceID(uint) (*entities.Ice, error)

	CreateOrder(*requests.OrderRequestBody) error
	QueryOrders() (*entities.Orders, error)
	QueryPrice(itemID uint, size string) (uint, error)

	// one time operation
	MigrateTable()
	DropTable()
}

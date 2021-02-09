package repositories

import (
	"github.com/pinyu/datalab-drinks-backend/src/domain/entities"
	"github.com/pinyu/datalab-drinks-backend/src/interface/requests"
)

// OrderRepository is Order CRUD interface
type OrderRepository interface {
	Exist(string) bool

	ValidateItemID(uint) (*entities.Item, error)
	ValidateSugarID(uint) (*entities.Sugar, error)
	ValidateIceID(uint) (*entities.Ice, error)

	CreateOrder(*requests.OrderRequestBody) error
	QueryWeekOrders() ([]entities.Order, error)

	// one time operation
	MigrateTable()
	DropTable()
}

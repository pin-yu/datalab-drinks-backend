package persistence

import (
	"fmt"
	"log"

	"github.com/pinyu/datalab-drinks-backend/src/domain/entities"
	"github.com/pinyu/datalab-drinks-backend/src/domain/repositories"
	"github.com/pinyu/datalab-drinks-backend/src/interface/requests"
	"github.com/pinyu/datalab-drinks-backend/src/utils"
	"gorm.io/gorm"
)

// orderRepository implements repository.OrderRepository
type orderRepository struct {
}

// NewOrderRepository returns initialized OrderRepositoryImpl
func NewOrderRepository() repositories.OrderRepository {
	return &orderRepository{}
}

// Exist returns true if the order exists
func (o *orderRepository) Exist(orderBy string) bool {
	db := newDBDriver()

	result := db.Where("order_by = ? AND created_at > ?", orderBy, utils.OrderIntervalStartTime()).Find(&entities.Order{})

	// if someone has ordered in this week, return true
	if result.RowsAffected > 0 {
		return true
	}

	return false
}

func (o *orderRepository) CreateOrder(orderRequest *requests.OrderRequestBody) error {
	db := newDBDriver()

	if err := db.Create(&entities.Order{
		OrderBy: orderRequest.OrderBy,
		Size:    orderRequest.Size,
		ItemID:  orderRequest.ItemID,
		SugarID: orderRequest.SugarID,
		IceID:   orderRequest.IceID,
	}).Error; err != nil {
		return fmt.Errorf("error occurs in creating an order record: %w", err)
	}
	return nil
}

func (o *orderRepository) ValidateItemID(id uint) (*entities.Item, error) {
	db := newDBDriver()

	item := &entities.Item{}

	result := db.Find(item, id)
	if result.RowsAffected > 0 {
		return item, nil
	}

	return nil, fmt.Errorf("invalid item_id")
}

func (o *orderRepository) ValidateSugarID(id uint) (*entities.Sugar, error) {
	db := newDBDriver()

	sugar := &entities.Sugar{}

	result := db.Find(sugar, id)
	if result.RowsAffected > 0 {
		return sugar, nil
	}

	return nil, fmt.Errorf("invalid sugar_id")
}

// ValidateIceID is more complicated because we need check whether the drinks could be made as hot or cold
func (o *orderRepository) ValidateIceID(id uint) (*entities.Ice, error) {
	db := newDBDriver()

	ice := &entities.Ice{}

	result := db.Find(ice, id)
	if result.RowsAffected > 0 {
		return ice, nil
	}

	return nil, fmt.Errorf("invalid ice_id")
}

func (o *orderRepository) QueryWeekOrders() ([]entities.Order, error) {
	db := newDBDriver()

	orders := []entities.Order{}
	err := db.Select("id, order_by, item_id, size, sugar_id, ice_id, Max(created_at)").Where("created_at > ?", utils.OrderIntervalStartTime()).Group("order_by").Order("Max(created_at)").Preload("Item").Preload("Sugar").Preload("Ice").Find(&orders).Error
	if err != nil {
		return nil, fmt.Errorf("error occurs in QueryWeekOrders: %w", err)
	}

	return orders, nil
}

// MigrateTable will create tables and insert some mandatory rows for this application
func (o *orderRepository) MigrateTable() {
	db := newDBDriver()

	// create tables
	autoMigrate(db)

	// insert mandatory rows into tables
	insertMandatoryRows(db)
}

func autoMigrate(db *gorm.DB) {
	err := db.AutoMigrate(&entities.Order{})
	if err != nil {
		log.Fatalf("error occurs in MigrateTable: %v", err)
	}
}

// insertMandatoryRows will inserts mandatory rows into Item, Sugar, Ice table.
func insertMandatoryRows(db *gorm.DB) {
	menuRepo := NewMenuRepository()
	menu, err := menuRepo.ReadMenu()
	if err != nil {
		log.Fatalf("error occurs at insertMandatoryRows: %v", err)
	}

	insertItemRows(db, menu)
	insertSugarRows(db, menu)
	insertIceRows(db, menu)
}

func insertItemRows(db *gorm.DB, menu *entities.Menu) {
	for _, series := range menu.Menu {
		db.Create(series.Items)
	}
}

func insertSugarRows(db *gorm.DB, menu *entities.Menu) {
	db.Create(menu.Sugar)
}

func insertIceRows(db *gorm.DB, menu *entities.Menu) {
	db.Create(menu.Ice)
}

func (o *orderRepository) DropTable() {
	conn := newDBDriver()

	// tables to drop
	tables := []interface{}{&entities.Order{}, &entities.Item{}, &entities.Sugar{}, &entities.Ice{}}

	err := conn.Migrator().DropTable(tables...)
	if err != nil {
		log.Fatalf("error occurs in DropTable: %v", err)
	}
}

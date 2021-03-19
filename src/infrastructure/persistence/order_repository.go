package persistence

import (
	"fmt"
	"log"
	"time"

	"github.com/pin-yu/datalab-drinks-backend/src/domain/entities"
	"github.com/pin-yu/datalab-drinks-backend/src/domain/repositories"
	"github.com/pin-yu/datalab-drinks-backend/src/interface/requests"
	"github.com/pin-yu/datalab-drinks-backend/src/utils"
	"gorm.io/gorm"
)

// orderRepository implements repository.OrderRepository
type orderRepository struct {
}

// NewOrderRepository returns initialized OrderRepositoryImpl
func NewOrderRepository() repositories.OrderRepository {
	return &orderRepository{}
}

// HasOrdered returns true if the order exists
func (o *orderRepository) HasOrdered(orderBy string) bool {
	db := newDBDriver()

	result := db.Where("order_by = ? AND order_time >= ?", orderBy, utils.OrderIntervalStartTime().UnixNano()).Find(&[]entities.Order{})

	// if someone has ordered in this week, return true
	return result.RowsAffected > 0
}

func (o *orderRepository) CreateOrder(orderRequest *requests.OrderRequestBody) error {
	db := newDBDriver()

	if err := db.Create(&entities.Order{
		OrderBy:   orderRequest.OrderBy,
		Size:      orderRequest.Size,
		ItemID:    orderRequest.ItemID,
		SugarID:   orderRequest.SugarID,
		IceID:     orderRequest.IceID,
		OrderTime: time.Now().UnixNano(),
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

func (o *orderRepository) QueryOrders() (*entities.Orders, error) {
	db := newDBDriver()

	orders := entities.NewOrders()

	if err := queryDetailOrders(o, db, orders); err != nil {
		return nil, fmt.Errorf("error occurs in queryDetailOrders: %w", err)
	}

	if err := queryAggregateOrders(o, db, orders); err != nil {
		return nil, fmt.Errorf("error occurs in queryAggregateOrders: %w", err)
	}

	return orders, nil
}

func queryDetailOrders(o *orderRepository, db *gorm.DB, orders *entities.Orders) error {
	err := db.Select("id, order_by, item_id, size, sugar_id, ice_id, Max(order_time) as order_time").Where("order_time >= ?", utils.OrderIntervalStartTime().UnixNano()).Group("order_by").Order("Max(order_time)").Preload("Item").Preload("Sugar").Preload("Ice").Find(&orders.DetailOrders).Error
	if err != nil {
		return err
	}

	return nil
}

func queryAggregateOrders(o *orderRepository, db *gorm.DB, orders *entities.Orders) error {
	subQuery := db.Model(entities.Order{}).Select("order_by, item_id, size, sugar_id, ice_id").Where("order_time >= ?", utils.OrderIntervalStartTime().UnixNano()).Group("order_by").Order("Max(order_time)")

	err := db.Table("(?) as u", subQuery).Select("count(*) as number, item_id, item, size, sugars.tag as sugar_tag, ices.tag as ice_tag").Joins("JOIN items on item_id=items.id").Joins("JOIN sugars on sugar_id=sugars.id").Joins("JOIN ices on ice_id=ices.id").Group("item_id, size, sugar_id, ice_id").Find(&orders.AggregateOrders).Error
	if err != nil {
		return err
	}

	for i, aggOrder := range orders.AggregateOrders {
		price, err := o.QueryPrice(aggOrder.ItemID, aggOrder.Size)
		if err != nil {
			return err
		}
		orders.AggregateOrders[i].SubTotalPrice = price * aggOrder.Number
	}

	// count the total price according to the aggreate orders
	orders.CountTotalPrice()

	return nil
}

// QueryPrice returns price given item_id and size
func (o *orderRepository) QueryPrice(itemID uint, size string) (uint, error) {
	var price struct {
		Price uint
	}
	var err error

	db := newDBDriver()

	if size == "medium" {
		err = db.Model(&entities.Item{}).Select("medium_price as price").Find(&price, itemID).Error
	} else {
		err = db.Model(&entities.Item{}).Select("large_price as price").Find(&price, itemID).Error
	}

	if err != nil {
		return 0, fmt.Errorf("error occurs at QueryPrice: %w", err)
	}

	return price.Price, nil
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
	insertItemRows(db)
	insertSugarRows(db)
	insertIceRows(db)
}

func insertItemRows(db *gorm.DB) {
	menuRepo := NewMenuRepository()
	menu, err := menuRepo.ReadMenu()
	if err != nil {
		log.Fatalf("error occurs at insertItemRows: %v", err)
	}

	for _, series := range menu.Menu {
		db.Create(series.Items)
	}
}

func insertSugarRows(db *gorm.DB) {
	sugarsRepo := NewSugarsRepository()
	sugars, err := sugarsRepo.ReadSugars()
	if err != nil {
		log.Fatalf("error occurs at insertSugarRows: %v", err)
	}

	db.Create(sugars.Sugars)
}

func insertIceRows(db *gorm.DB) {
	icesRepo := NewIcesRepository()
	ices, err := icesRepo.ReadIces()
	if err != nil {
		log.Fatalf("error occurs at insertIcesRows: %v", err)
	}

	db.Create(ices.Ices)
}

func (o *orderRepository) DropTable() {
	db := newDBDriver()

	// tables to drop
	tables := []interface{}{&entities.Order{}, &entities.Item{}, &entities.Sugar{}, &entities.Ice{}}

	err := db.Migrator().DropTable(tables...)
	if err != nil {
		log.Fatalf("error occurs in DropTable: %v", err)
	}
}

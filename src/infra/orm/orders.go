package orm

import (
	"log"
	"time"

	"gorm.io/gorm"
)

// Order will be pluralized to Orders by gorm
type Order struct {
	ID        uint
	OrderBy   string
	Item      string
	Sugar     uint8
	Ice       uint8
	CreatedAt time.Time
	UpdatedAt time.Time
}

func migrateOrder(db *gorm.DB) {
	db.AutoMigrate(&Order{})
	log.Println("migrate order table")
}

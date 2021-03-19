package entities

// Item represents Cama drinks items
type Item struct {
	ID          uint   `yaml:"item_id" gorm:"primaryKey"`
	Item        string `yaml:"item" gorm:"not null"`
	MediumPrice uint   `yaml:"medium_price"`
	LargePrice  uint   `yaml:"large_price"`
	Sugar       bool   `yaml:"sugar_adjustable" gorm:"not null"`
	Cold        bool   `yaml:"cold_adjustable"  gorm:"not null"`
	Hot         bool   `yaml:"hot_adjustable" gorm:"not null"`
}

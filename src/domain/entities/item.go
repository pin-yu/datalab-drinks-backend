package entities

// Item represents Cama drinks items
type Item struct {
	ID          uint   `yaml:"item_id" json:"item_id" gorm:"primaryKey"`
	Item        string `yaml:"item" json:"item" gorm:"not null"`
	MediumPrice uint   `yaml:"medium_price" json:"medium_price"`
	LargePrice  uint   `yaml:"large_price" json:"large_price"`
	Cold        bool   `yaml:"cold" json:"cold" gorm:"not null"`
	Hot         bool   `yaml:"hot" json:"hot" gorm:"not null"`
}

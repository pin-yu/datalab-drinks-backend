package entities

// Sugar represents drinks sugar
type Sugar struct {
	ID  uint   `yaml:"sugar_id" json:"sugar_id" gorm:"primaryKey"`
	Tag string `yaml:"sugar_tag" json:"sugar_tag" gorm:"not null"`
}

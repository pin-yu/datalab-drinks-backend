package entities

// Ice represents drinks ice
type Ice struct {
	ID  uint   `yaml:"ice_id" gorm:"primaryKey"`
	Tag string `yaml:"ice_tag"  gorm:"not null"`
}

// IsHot returns true if this ID stands for hot drinks
func (ice *Ice) IsHot() bool {
	hotID := uint(1)
	return ice.ID == hotID
}

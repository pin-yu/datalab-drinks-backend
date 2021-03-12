package entities

// Ice represents drinks ice
type Ice struct {
	ID  uint   `yaml:"ice_id" json:"ice_id" gorm:"primaryKey"`
	Tag string `yaml:"ice_tag" json:"ice_tag" gorm:"not null"`
}

// IsHotID return true if this ID means hot
func (ice *Ice) IsHotID() bool {
	return ice.ID == 1
}

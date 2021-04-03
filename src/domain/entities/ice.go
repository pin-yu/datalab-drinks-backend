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

func (ice *Ice) IsValidIce(hotAdjustable bool, iceAdjustable bool) bool {
	if hotAdjustable && iceAdjustable {
		return true
	}

	if hotAdjustable && !iceAdjustable {
		return ice.IsHot()
	}

	if iceAdjustable && !hotAdjustable {
		return !ice.IsHot()
	}

	return false
}

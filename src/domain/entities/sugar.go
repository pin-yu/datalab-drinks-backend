package entities

// Sugar represents drinks sugar
type Sugar struct {
	ID  uint   `yaml:"sugar_id" gorm:"primaryKey"`
	Tag string `yaml:"sugar_tag" gorm:"not null"`
}

func (s *Sugar) IsNormalSugar() bool {
	normalSugarID := uint(4)
	return s.ID == normalSugarID
}

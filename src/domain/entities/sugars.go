package entities

// Sugars represents sugar list
type Sugars struct {
	Sugars []Sugar `yaml:"sugars"`
}

func (s *Sugars) AdjustSugarList(sugarAdjustable bool) []Sugar {
	if sugarAdjustable {
		return s.Sugars[:]
	} else {
		return s.Sugars[3:]
	}
}

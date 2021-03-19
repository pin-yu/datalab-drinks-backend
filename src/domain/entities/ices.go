package entities

import "log"

// Ices represents ice list
type Ices struct {
	Ices []Ice `yaml:"ices"`
}

func (i *Ices) AdjustIceList(hotAdjustable bool, iceAdjustable bool) []Ice {
	if hotAdjustable && iceAdjustable {
		return i.Ices[:]
	} else if hotAdjustable {
		return i.Ices[0:1]
	} else if iceAdjustable {
		return i.Ices[1:]
	} else {
		log.Fatal("Invalid arguments in AdjustIceList")
	}

	return nil
}

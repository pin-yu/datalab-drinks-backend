package entities

// Menu represents Cama's menu
type Menu struct {
	MenuVersion string   `yaml:"menu_version" json:"menu_version"`
	Menu        []Series `yaml:"menu" json:"menu"`
	Sugar       []Sugar  `yaml:"sugar" json:"sugar"`
	Ice         []Ice    `yaml:"ice" json:"ice"`
}

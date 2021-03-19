package entities

// Menu represents Cama's menu
// Menu is a domain entity for internal usage instead of presentation
type Menu struct {
	MenuVersion string   `yaml:"menu_version"`
	Menu        []Series `yaml:"menu"`
}

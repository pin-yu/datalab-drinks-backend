package menus

import (
	"log"
	"path/filepath"

	"github.com/pinyu/datalab-drinks-backend/src/utils"
	"gopkg.in/yaml.v2"
)

var (
	basePath = utils.GetBasePath()
	camaYaml = filepath.Join(basePath, "../../assets/cama_menu.yaml")
)

// CamaMenu contains two fields: MenuVersion, Menu
type CamaMenu struct {
	MenuVersion string   `yaml:"menu_version" json:"menu_version"`
	Menu        []Series `yaml:"menu" json:"menu"`
	Sugar       []Sugar  `yaml:"sugar" json:"sugar"`
	Ice         []Ice    `yaml:"ice" json:"ice"`
}

// Sugar contains two fields: ID, Tag
type Sugar struct {
	ID  string `yaml:"id" json:"id"`
	Tag string `yaml:"tag" json:"tag"`
}

// Ice contains two fields: ID, Tag
type Ice struct {
	ID  string `yaml:"id" json:"id"`
	Tag string `yaml:"tag" json:"tag"`
}

// Series contains two fields: Series, Items
type Series struct {
	Series string  `yaml:"series" json:"series"`
	Items  []Items `yaml:"items" json:"items"`
}

// Items contains four fields: Item, Price, Cold, Hot
type Items struct {
	ID     int      `yaml:"id" json:"id"`
	Item   string   `yaml:"item" json:"item"`
	Flavor []string `yaml:"flavor" json:"flavor"`
	Prices Prices   `yaml:"prices" json:"prices"`
	Cold   bool     `yaml:"cold" json:"cold"`
	Hot    bool     `yaml:"hot" json:"hot"`
}

// Prices contains two fields: Large, Medium
type Prices struct {
	Large  int `yaml:"large" json:"large"`
	Medium int `yaml:"medium" json:"medium"`
}

// GetMenus will return a pointer of type CamaDrinks
func GetMenus() *CamaMenu {
	yamlContent := utils.ReadFile(camaYaml)

	menus := &CamaMenu{}
	err := yaml.Unmarshal([]byte(yamlContent), menus)
	if err != nil {
		log.Fatalf("error occurs in GetCamaDrinks: %v", err)
	}

	return menus
}

package persistence

import (
	"github.com/pinyu/datalab-drinks-backend/src/config"
	"github.com/pinyu/datalab-drinks-backend/src/domain/entities"
	"github.com/pinyu/datalab-drinks-backend/src/domain/repositories"

	"github.com/pinyu/datalab-drinks-backend/src/utils"
	"gopkg.in/yaml.v2"
)

var (
	camaYaml = config.GetCamaYamlPath()
)

// menuRepository implements repository.MenuRepository
type menuRepository struct {
}

// NewMenuRepository returns initialized MenuRepositoryImpl
func NewMenuRepository() repositories.MenuRepository {
	return &menuRepository{}
}

// ReadMenu returns cama menu
func (m *menuRepository) ReadMenu() (*entities.Menu, error) {
	yamlContent := utils.ReadFile(camaYaml)

	menu := &entities.Menu{}
	err := yaml.Unmarshal([]byte(yamlContent), menu)
	if err != nil {
		return nil, err
	}

	return menu, nil
}

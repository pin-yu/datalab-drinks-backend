package persistence

import (
	"github.com/pin-yu/datalab-drinks-backend/src/config"
	"github.com/pin-yu/datalab-drinks-backend/src/domain/entities"
	"github.com/pin-yu/datalab-drinks-backend/src/domain/repositories"

	"github.com/pin-yu/datalab-drinks-backend/src/utils"
	"gopkg.in/yaml.v2"
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
	yamlContent := utils.ReadFile(config.GetCamaYamlPath())

	menu := &entities.Menu{}

	err := yaml.Unmarshal([]byte(yamlContent), menu)
	if err != nil {
		return nil, err
	}

	return menu, nil
}

package persistence

import (
	"github.com/pin-yu/datalab-drinks-backend/src/config"
	"github.com/pin-yu/datalab-drinks-backend/src/domain/entities"
	"github.com/pin-yu/datalab-drinks-backend/src/domain/repositories"
	"github.com/pin-yu/datalab-drinks-backend/src/utils"
	"gopkg.in/yaml.v2"
)

type sugarsRepository struct {
}

// NewSugarRepository returns initialized SugarRepositoryImpl
func NewSugarsRepository() repositories.SugarsRepository {
	return &sugarsRepository{}
}

func (s *sugarsRepository) ReadSugars() (*entities.Sugars, error) {
	yamlContent := utils.ReadFile(config.GetSugarYamlPath())

	sugars := &entities.Sugars{}

	err := yaml.Unmarshal([]byte(yamlContent), sugars)
	if err != nil {
		return nil, err
	}

	return sugars, nil
}

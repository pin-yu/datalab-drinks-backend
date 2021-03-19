package persistence

import (
	"github.com/pin-yu/datalab-drinks-backend/src/config"
	"github.com/pin-yu/datalab-drinks-backend/src/domain/entities"
	"github.com/pin-yu/datalab-drinks-backend/src/domain/repositories"
	"github.com/pin-yu/datalab-drinks-backend/src/utils"
	"gopkg.in/yaml.v2"
)

type icesRepository struct {
}

// NewIcesRepository returns initialized IcesRepositoryImpl
func NewIcesRepository() repositories.IcesRepository {
	return &icesRepository{}
}

func (s *icesRepository) ReadIces() (*entities.Ices, error) {
	yamlContent := utils.ReadFile(config.GetIceYamlPath())

	ices := &entities.Ices{}

	err := yaml.Unmarshal([]byte(yamlContent), ices)
	if err != nil {
		return nil, err
	}

	return ices, nil
}

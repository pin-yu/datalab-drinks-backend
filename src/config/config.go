package config

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/pin-yu/datalab-drinks-backend/src/utils"
)

var (
	basePath = utils.GetBasePath()
)

// GetDBPath return db path corresponding to the GIN_MODE
func GetDBPath() string {
	var dbMode string

	env := os.Getenv("GIN_MODE")
	switch env {
	case "release":
		dbMode = "release.db"
	case "test":
		dbMode = "test.db"
	default:
		dbMode = "dev.db"
	}

	return filepath.Join(basePath,
		fmt.Sprintf("../infrastructure/local/%v", dbMode))
}

// GetCamaYamlPath return cama yaml path
func GetCamaYamlPath() string {
	return getPathFromBase("../../assets/cama_menu.yaml")
}

func GetSugarYamlPath() string {
	return getPathFromBase("../../assets/sugar.yaml")
}

func GetIceYamlPath() string {
	return getPathFromBase("../../assets/ice.yaml")
}

func getPathFromBase(path string) string {
	return filepath.Join(basePath, path)
}

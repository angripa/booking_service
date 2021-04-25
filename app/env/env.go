package env

import (
	"log"
	"os"

	"github.com/angripa/booking_service"
)

var (
	Conf *models.Config
)

func initConfig() {
	var err error
	configFolder := "config"
	env := os.Getenv("ENV")
	if env == "" {
		env = "dev"
	}
	Conf, err = getConfiguration(configFolder, env)
	if err != nil {
		log.Fatal(err)
	}
}

func getConfiguration(configFolder string, envType string) (*models.Config, error) {
	if envType == "" {
		envType = "dev"
	}
	var configuration structs.Config
	file, err := os.Open(fmt.Sprintf("%s/%s.json", configFolder, envType))
	if err != nil {
		return nil, err
	}
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&configuration)
	if err != nil {
		return nil, err
	}
	return &configuration, err
}

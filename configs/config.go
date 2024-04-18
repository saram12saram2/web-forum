package configs

import (
	"encoding/json"
	"os"
)

type Config struct {
	Port    string
	Migrate string
	DB      struct {
		Dsn    string
		Driver string
	}
}

func LoadConfig() (Config, error) {
	configFile, err := os.Open("./configs/config.json")
	if err != nil {
		return Config{}, err
	}
	defer configFile.Close()

	var config Config
	decoder := json.NewDecoder(configFile)
	err = decoder.Decode(&config)
	if err != nil {
		return Config{}, err
	}
	return config, nil
}

package openframe

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Config struct {
	Network struct {
		ApiBaseUrl string `json:"api_base"`
		AppBaseUrl string `json:"app_base"`
		ApiUrl string `json:"api_url"`
	}
	Autoboot bool `json:"autoboot"`
	PubSubUrl string `json:"pubsub_url"`
}

func (config *Config) Save(configFile string) error {
	bytes, err := json.Marshal(config)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	return ioutil.WriteFile(configFile, bytes, 0644)
}

func LoadConfig(filePath string) (Config, error) {
	file, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Println(err.Error())
		return Config{}, err
	}

	var config Config
	json.Unmarshal(file, &config)

	return config, nil
}

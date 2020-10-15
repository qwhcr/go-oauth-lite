package config

import (
	"io/ioutil"
	"log"

	configConstants "go-oauth-lite/constants/config"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Tenancy string `yaml:"tenancy"`
	Port    string `yaml:"port"`
    FirebaseSecretURL string `yaml:"firebaseSecretURL"`
    LoginAssetsURL string `yaml:"loginAssetsURL"`
}

var config Config = Config{}

func ReadConfig() {
	configBytes, err := ioutil.ReadFile(configConstants.ConfigPath)
	if err != nil {
		log.Fatalln(err)
	}

	err = yaml.Unmarshal(configBytes, &config)
	if err != nil {
		log.Fatalln(err)
	}
}

func GetConfig() *Config {
	return &config
}

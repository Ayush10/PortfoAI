package configs

import (
    "gopkg.in/yaml.v3"
    "io/ioutil"
    "log"
)

type AppConfig struct {
    Database struct {
        Host     string `yaml:"host"`
        Port     string `yaml:"port"`
        User     string `yaml:"user"`
        Password string `yaml:"password"`
        Dbname   string `yaml:"dbname"`
    } `yaml:"database"`
}

var AppConfiguration *AppConfig

func LoadConfig() {
    file, err := ioutil.ReadFile("configs/config.yaml")
    if err != nil {
        log.Fatalf("Unable to read configuration file: %v", err)
    }

    config := &AppConfig{}
    err = yaml.Unmarshal(file, config)
    if err != nil {
        log.Fatalf("Failed to parse configuration file: %v", err)
    }

    AppConfiguration = config
    log.Println("Configuration file loaded successfully")
}

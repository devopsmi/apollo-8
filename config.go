package main

import (
  "log"
  "os"
	
	"github.com/monologid/apollo-8/model"
	"github.com/monologid/apollo-8/helper"
  "github.com/BurntSushi/toml"  
)

var APP_HOST string
var APP_PORT string

// ReadBaseConfig is a method to read default configuration file
// in toml format
func ReadBaseConfig() model.ConfigSchema {
  file := helper.ReadFile(`config.toml`)
  conf := model.ConfigSchema{}
  
  if _, err := toml.Decode(file, &conf); err != nil {
    log.Fatal(err)
  }
  
  return conf
}

// InitializeMainConfig is a method to initalize main config
func InitializeMainConfig() {
  conf := ReadBaseConfig()
  APP_HOST = conf.Apollo8.Host
  APP_PORT = conf.Apollo8.Port
}

// InitializeServicesConfig is a method to initialize
// main services configuration file
func InitializeServicesConfig() {
  serviceConfigFile, err := os.OpenFile(`apoconf.json`, os.O_RDWR|os.O_CREATE, 0755) 
  if err != nil {
    log.Fatal(err)
  }
  
  if err := serviceConfigFile.Close(); err != nil {
    log.Fatal(err)
  }
}
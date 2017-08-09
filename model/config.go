package model

// ConfigSchema is a schema for default config file 
// in .toml format
type ConfigSchema struct {
  Apollo8 struct { 
    Host string
    Port string 
    Username string
    Password string
    Websrc string
  }
  
  Nginx struct { 
    Dir string 
  }
  
  Services struct { 
    Dir string 
  }
}